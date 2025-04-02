//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
	"go.bytecodealliance.org/cm"
	"wasmcv.org/wasm/cv/mat"
)

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

	return image
}
