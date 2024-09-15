//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
	"github.com/wasmvision/wasmcv/components/tinygo/wasm/cv/cv"
	"github.com/wasmvision/wasmcv/components/tinygo/wasm/cv/mat"
	"github.com/wasmvision/wasmcv/components/tinygo/wasm/cv/types"
)

//go:wasmimport hosted complete
func complete()

//go:wasmimport hosted println
func println(ptr, size uint32)

//export process
func process(image mat.Mat) mat.Mat {
	println(convert.StringToWasmPtr("Cols: " +
		convert.IntToString(int(image.Cols())) +
		" Rows: " +
		convert.IntToString(int(image.Rows())) +
		" Type: " +
		convert.IntToString(int(image.Type()))))

	imageOut := cv.GaussianBlur(image, types.Size{5, 5}, 1.5, 1.5, types.BorderTypeBorderReflect101)
	println(convert.StringToWasmPtr("Performed GaussianBlur on image"))

	complete()
	return imageOut
}

func main() {}
