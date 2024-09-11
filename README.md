# wasmCV

wasmCV provides WebAssembly interface bindings for computer vision applications based on [OpenCV](https://github.com/opencv/opencv).

It includes [WIT](https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md) files defining the interface to be used between a WebAssembly host application and a WASM guest module intended to process  OpenCV `Mat` image frames.

These are then used to generate WASM bindings for TinyGo, Rust, and C. Those bindings can then be used in a WASM guest module to obtain information or perform operations on OpenCV image frames.

## Example wasmCV module

This TinyGo module exports a `process()` function to the WASM host application, which passes in the wasmCV image `Mat` to be processed. It then calls functions on that `Mat` which are handled by the host application by calling OpenCV to actually perform the operations.

```go
package main

import (
	"github.com/hybridgroup/mechanoid/convert"
	"github.com/hybridgroup/wasmcv/components/tinygo/wasm/cv/mat"
)

//go:wasmimport hosted println
func println(ptr, size uint32)

//go:export process
func process(image mat.Mat) mat.Mat {
	println(convert.StringToWasmPtr("Cols: " +
		convert.IntToString(int(image.Cols())) +
		" Rows: " +
		convert.IntToString(int(image.Rows())) +
		" Type: " +
		convert.IntToString(int(image.Type()))))

	return image
}

func main() {}
```

You can then compile this module using the TinyGo compiler.

```shell
tinygo build -o processor.wasm -target=wasm-unknown processor.go
```

Note that the `wasm-unknown` target can be used with wasmCV to produce very lightweight guest modules. The example above compiles to around 31k, including debug information.

```shell
-rwxrwxr-x 1 ron ron 31248 sep 11 11:00 processor.wasm
```

See the [basic example application here](./examples/basic) to give it a try.

## WASM Component Generation

WASM Guest bindings are generated using `wit-bindgen` v0.31 or above

https://github.com/bytecodealliance/wit-bindgen

Go bindings are generated by `wit-bindgen-go`

https://github.com/bytecodealliance/wasm-tools-go

### TinyGo

```shell
wit-bindgen-go generate --out ./components/tinygo ./wit/
```

### Rust

```shell
wit-bindgen rust --out-dir ./components/rust/ ./wit/
```

### C

```shell
wit-bindgen c --out-dir ./components/c/ ./wit/
```

