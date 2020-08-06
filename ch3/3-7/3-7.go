package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	file, _ := os.OpenFile("./out.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	_ = png.Encode(file, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: uint8(255 - contrast*n), A: 255}
		}
	}
	return color.Black
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{Y: 255 - contrast*i}
		}
	}
	return color.Black
}
