package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/orsinium-labs/wypes"
	"github.com/tetratelabs/wazero"
	"gocv.io/x/gocv"
)

//go:embed modules/processor.wasm
var processorFrameWasm []byte

//go:embed modules/processrs.wasm
var processrsFrameWasm []byte

//go:embed modules/processc.wasm
var processcFrameWasm []byte

var (
	processor = flag.String("processor", "tinygo", "which wasmCV processor to use (tinygo|rust|c)")

	frame gocv.Mat
)

func main() {
	flag.Parse()

	var module []byte
	switch *processor {
	case "tinygo":
		module = processorFrameWasm
	case "rust":
		module = processrsFrameWasm
	case "c":
		module = processcFrameWasm
	default:
		log.Panicf("unsupported processor: %s", *processor)
	}

	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	println("Defining host functions...")
	modules := wypes.Modules{
		"hosted": wypes.Module{
			"println": wypes.H1(hostPrintln),
		},
		"wasm:cv/mat": wypes.Module{
			"[method]mat.cols": wypes.H1(matColsFunc),
			"[method]mat.rows": wypes.H1(matRowsFunc),
			"[method]mat.type": wypes.H1(matTypeFunc),
		},
	}

	err := modules.DefineWazero(r, nil)
	if err != nil {
		fmt.Printf("error define host functions: %v", err)
		return
	}

	fmt.Printf("Loading %s wasmCV guest module...\n", *processor)
	mod, err := r.Instantiate(ctx, module)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}
	process := mod.ExportedFunction("process")

	// Open the webcam.
	deviceID := "/dev/video0"
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
		fmt.Print("\033[H\033[2J")

		fmt.Printf("Running %s wasmCV module\n", *processor)

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
