package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "parse form err: %v", err)
	}
	x := getValue(r.Form["x"])
	y := getValue(r.Form["y"])
	zoom := getValue(r.Form["zoom"])
	render(w, x, y, zoom)
}

func render(out io.Writer, x float64, y float64, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	exp2 := math.Exp2(1 - zoom)
	xmin, xmax := x-exp2, x+exp2
	ymin, ymax := y-exp2, y+exp2

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	_ = png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: uint8(255 - contrast*n)}
		}
	}
	return color.Black
}

func getValue(form []string) float64 {
	if len(form) == 0 {
		return 0
	}
	v, err := strconv.ParseFloat(form[0], 64)
	if err != nil {
		return 0
	}
	return v
}
