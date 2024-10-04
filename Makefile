
generate-go:
	wit-bindgen-go generate --out ./components/tinygo -w imports -p wasmcv.org ./wit/

generate-rust:
	wit-bindgen rust --out-dir ./components/rust/wasmcv/src -w imports ./wit/

generate-c:
	wit-bindgen c --out-dir ./components/c/wasmcv -w imports ./wit/

generate: generate-go generate-rust generate-c

# build basic example guest module
basic:
	cd examples/basic/modules/processor; tinygo build -o ../processor.wasm -target=wasm-unknown .

# build multi example guest modules
multi:
	cd examples/multi/modules/processc; ./build.sh
	cd examples/multi/modules/processor; tinygo build -o ../processor.wasm -target=wasm-unknown .
	cd examples/multi/modules/processrs; cargo build --target wasm32-unknown-unknown --release; \
		cp ./target/wasm32-unknown-unknown/release/processrs.wasm ../

# build wasmcam example guest module
wasmcam:
	cd examples/wasmcam/modules/processor; tinygo build -o ../processor.wasm -target=wasm-unknown .

examples: basic multi wasmcam
