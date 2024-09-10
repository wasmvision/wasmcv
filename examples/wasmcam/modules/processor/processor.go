//go:build tinygo

package main

import "github.com/hybridgroup/wasmcv/components/tinygo/wasm/cv/mat"

//go:wasmimport hosted complete
func complete()

//go:wasmimport hosted showInfo
func showInfo(cols, rows uint32)

//go:export process
func process(image mat.Mat) mat.Mat {
	showInfo(image.Cols(), image.Rows())

	complete()
	return image
}

func main() {}
