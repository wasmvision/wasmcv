//go:build tinygo

package main

import (
	"unsafe"

	"github.com/hybridgroup/mechanoid/convert"
	"wasmcv.org/wasm/cv/mat"
)

//go:wasmimport hosted println
func println(ptr, size uint32)

//export process
func process(image mat.Mat) mat.Mat {
	println(convert.StringToWasmPtr("Cols: " +
		convert.IntToString(int(image.Cols())) +
		" Rows: " +
		convert.IntToString(int(image.Rows())) +
		" Type: " +
		convert.IntToString(int(image.Mattype())) +
		" Size: " +
		convert.IntToString(int(image.Size().Len()))))

	return image
}

//export malloc
func malloc(size uint32) uint32 {
	data := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(unsafe.SliceData(data)))

	return uint32(ptr)
}

//export free
func free(ptr uint32) {
}
