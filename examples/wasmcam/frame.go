package main

import (
	"image"

	"github.com/orsinium-labs/wypes"
	"gocv.io/x/gocv"
)

var (
	frame     gocv.Mat
	processed gocv.Mat
)

func matColsFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Cols())
}

func matRowsFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Rows())
}

func matTypeFunc(matref wypes.UInt32) wypes.UInt32 {
	return wypes.UInt32(frame.Type())
}

func matGaussianBlurFunc(s *wypes.Store, matref wypes.UInt32, size0 wypes.UInt32, size1 wypes.UInt32, sigmaX0 wypes.Float32, sigmaY0 wypes.Float32, border0 wypes.UInt32, result wypes.Result[wypes.UInt32, wypes.UInt32, wypes.UInt32]) wypes.Void {
	err := gocv.GaussianBlur(frame, &processed, image.Pt(int(size0), int(size1)), float64(sigmaX0), float64(sigmaY0), gocv.BorderType(border0))
	if err != nil {
		println("Error processing image")
		result.IsError = true
		result.Error = wypes.UInt32(1)
		result.Lower(s)
		return wypes.Void{}
	}

	result.IsError = false
	result.OK = wypes.UInt32(1)
	result.Lower(s)

	return wypes.Void{}
}
