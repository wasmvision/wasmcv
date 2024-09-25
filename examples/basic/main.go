package main

import (
	"context"
	_ "embed"
	"encoding/binary"
	"flag"
	"fmt"
	"log"

	"github.com/orsinium-labs/wypes"
	"github.com/tetratelabs/wazero"
	"gocv.io/x/gocv"
)

// processFrameWasm was generated by the following:
//
//	cd modules/processor; tinygo build -o ../processor.wasm -target=wasm-unknown .
//
//go:embed modules/processor.wasm
var processFrameWasm []byte

var frame gocv.Mat

var guestDataPtr uint32

func main() {
	flag.Parse()

	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	println("Defining host function...")
	modules := wypes.Modules{
		"hosted": wypes.Module{
			"println": wypes.H1(hostPrintln),
		},
		"wasm:cv/mat": wypes.Module{
			"[method]mat.cols":    wypes.H1(matColsFunc),
			"[method]mat.rows":    wypes.H1(matRowsFunc),
			"[method]mat.mattype": wypes.H1(matTypeFunc),
			"[method]mat.size":    wypes.H3(matSizeFunc),
		},
	}

	err := modules.DefineWazero(r, nil)
	if err != nil {
		fmt.Printf("error define host functions: %v", err)
		return
	}

	mod, err := r.Instantiate(ctx, processFrameWasm)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}
	process := mod.ExportedFunction("process")
	malloc := mod.ExportedFunction("malloc")
	res, err := malloc.Call(ctx, 256)
	if err != nil {
		log.Panicf("failed to call malloc: %v", err)
	}

	guestDataPtr = uint32(res[0])

	// Open the webcam.
	deviceID := "0"
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// streaming, capture from webcam
	frame = gocv.NewMat()
	defer frame.Close()

	fmt.Printf("Start reading device: %v\n", deviceID)
	i := 0
	for {
		if ok := webcam.Read(&frame); !ok {
			fmt.Printf("frame error %v\n", deviceID)
			continue
		}
		if frame.Empty() {
			continue
		}

		// clear screen
		//fmt.Print("\033[H\033[2J")

		i++
		fmt.Printf("Read frame %d\n", i+1)
		_, err := process.Call(ctx, 1)
		if err != nil {
			fmt.Printf("Error calling process: %v\n", err)
		}
	}
}

func hostPrintln(msg wypes.String) wypes.Void {
	println(msg.Unwrap())
	return wypes.Void{}
}

func matColsFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Cols())
}

func matRowsFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Rows())
}

func matTypeFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Type())
}

func matSizeFunc(s wypes.Store, matref wypes.UInt32, ptr wypes.UInt32) wypes.Void {
	dims := frame.Size()

	result := make([]uint32, len(dims))
	data := make([]byte, 4)
	for i, dim := range dims {
		result[i] = uint32(dim)
		binary.LittleEndian.PutUint32(data, uint32(dim))
		s.Memory.Write(guestDataPtr+uint32(i*4), data)
	}

	l := make([]byte, 8)
	binary.LittleEndian.PutUint32(l[0:], uint32(guestDataPtr))
	binary.LittleEndian.PutUint32(l[4:], uint32(len(dims)))

	s.Memory.Write(ptr.Unwrap(), l)

	return wypes.Void{}
}
