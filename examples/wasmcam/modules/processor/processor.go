//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
	"go.bytecodealliance.org/cm"
	"wasmcv.org/wasm/cv/cv"
	"wasmcv.org/wasm/cv/mat"
	"wasmcv.org/wasm/cv/types"
)

//go:wasmimport hosted complete
func complete()

//go:wasmimport hosted println
func println(ptr *byte, size uint32)

//export process
func process(image mat.Mat) mat.Mat {
	println(cm.LowerString("Cols: " +
		convert.IntToString(int(image.Cols())) +
		" Rows: " +
		convert.IntToString(int(image.Rows())) +
		" Type: " +
		convert.IntToString(int(image.Mattype()))))

	imageOut, _, isErr := cv.GaussianBlur(image, types.Size{X: 5, Y: 5}, 1.5, 1.5, types.BorderTypeBorderReflect101).Result()
	if isErr {
		println(cm.LowerString("Error processing image"))
		return image
	}

	println(cm.LowerString("Performed GaussianBlur on image"))

	complete()
	return imageOut
}
