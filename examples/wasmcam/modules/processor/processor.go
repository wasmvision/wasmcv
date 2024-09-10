//go:build tinygo

package main

import "github.com/hybridgroup/wasmcv/components/tinygo/cv/wasm/cv/mat"

//go:wasmimport hosted complete
func complete()

//go:export process
func process(image mat.Mat) mat.Mat {
	complete()
	return image
}

func main() {}
