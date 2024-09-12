# Basic

Basic example that reads frames from a connected webcam, and processes them using a wasmCV guest module compiled using [TinyGo](https://tinygo.org).

The host application is a Go application written using [GoCV Go language wrappers for OpenCV](https://github.com/hybridgroup/gocv) and the [Wazero WASM runtime](https://github.com/tetratelabs/wazero).

## Compile the guest module

```shell
$ cd modules/processor
$ tinygo build -o ../processor.wasm -target=wasm-unknown .
$ cd ..
```

## Run the host application

```shell                           
$ go run .
Defining host function...
Start reading device: /dev/video0
```

It will then capture each frame, display some stats, and them perform a gaussian blur on the source image.

```shell
Read frame 1
Cols: 640 Rows: 480 Type: 16
Frame complete
```
