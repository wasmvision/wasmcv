module github.com/wasmvision/wasmcv/examples/basic/modules/processor

go 1.23.0

toolchain go1.24.1

replace wasmcv.org/wasm/cv => ../../../../components/tinygo/wasm/cv

require (
	github.com/hybridgroup/mechanoid v0.2.0
	wasmcv.org/wasm/cv v0.6.0
)

require go.bytecodealliance.org/cm v0.3.0
