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
	for py := 0; py < height; py = py + 2 {
		y1 := float64(py)/float64(height)*(ymax-ymin) + ymin
		y2 := float64(py+1)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px = px + 2 {
			x1 := float64(px)/float64(width)*(xmax-xmin) + xmin
			x2 := float64(px+1)/float64(width)*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			img.Set(px/2, py/2, color.Gray{Y: (mandelbrot(z1).Y + mandelbrot(z2).Y + mandelbrot(z3).Y + mandelbrot(z4).Y) / 4.0})
		}
	}
	file, _ := os.OpenFile("./out.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	_ = png.Encode(file, img)
}

func mandelbrot(z complex128) color.Gray {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{uint8(255 - contrast*n)}
		}
	}
	return color.Gray{}
}
