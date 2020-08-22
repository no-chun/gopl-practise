package main

import (
	"gopl-practise/ch8/8-5/mandelbrot"
	"image/png"
	"os"
	"runtime"
)

func main() {
	workers := runtime.GOMAXPROCS(-1)
	img := mandelbrot.ConcurrentRender(workers)
	png.Encode(os.Stdout, img)
}
