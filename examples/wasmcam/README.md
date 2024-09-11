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
Read frame 2
Cols: 640 Rows: 480 Type: 16
Frame complete

Read frame 3
Cols: 640 Rows: 480 Type: 16
Frame complete

Read frame 4
Cols: 640 Rows: 480 Type: 16
Frame complete

Read frame 5
Cols: 640 Rows: 480 Type: 16
Frame complete

Read frame 6
Cols: 640 Rows: 480 Type: 16
Frame complete

Read frame 7
Cols: 640 Rows: 480 Type: 16
Frame complete
```
