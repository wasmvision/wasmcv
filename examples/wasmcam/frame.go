package main

import (
	"github.com/orsinium-labs/wypes"
	"gocv.io/x/gocv"
)

var (
	frame gocv.Mat
)

func matColsFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Cols())
}

func matRowsFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Rows())
}
