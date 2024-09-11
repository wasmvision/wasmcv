# WASMcam

Example that reads frames from a connected webcam, and processes them using a WASM guest module.

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
Performed GaussianBlur on image
Frame complete
```
