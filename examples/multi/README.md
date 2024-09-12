# Multi

Example that reads frames from a connected webcam, and processes them using a WASM guest module written using TinyGo, Rust, or C.

## Compile the TinyGo guest module

```shell
$ cd modules/processor
$ tinygo build -o ../processor.wasm -target=wasm-unknown .
$ cd ..
```

## Compile the Rust guest module

```shell
$ cd modules/processrs
$ cargo build --target wasm32-unknown-unknown --release
$ cd ../..
$ cp ./modules/processrs/target/wasm32-unknown-unknown/release/processrs.wasm ./modules/
```

## Run the host application using the TinyGo wasmCV guest module

```shell                           
$ go run . -processor=tinygo
Defining host functions...
Loading tinygo wasmCV guest module...
Start reading device: /dev/video0
```

It will then capture each frame, display some stats, and them perform a gaussian blur on the source image.

```shell
Running tinygo wasmCV module
Read frame 34
Cols: 640 Rows: 480
```

## Run the host application using the Rust wasmCV guest module

```shell                           
$ go run . -processor=rust
Defining host functions...
Loading rust wasmCV guest module...
Start reading device: /dev/video0
```

```shell
Running rust wasmCV module
Read frame 28
Cols: 640 Rows: 480
```
