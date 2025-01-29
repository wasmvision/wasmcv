module github.com/wasmvision/wasmcv/examples/basic/modules/processor

go 1.22.0

replace wasmcv.org/wasm/cv => ../../../../components/tinygo/wasm/cv

require (
	github.com/hybridgroup/mechanoid v0.2.0
	wasmcv.org/wasm/cv v0.6.0
)

require go.bytecodealliance.org/cm v0.1.1-0.20250127224617-c53528d4b8b4 // indirect
