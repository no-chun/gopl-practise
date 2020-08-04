package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var Green = color.RGBA{
	R: 0,
	G: 128,
	B: 0,
	A: 0xFF,
}

var palette = []color.Color{color.White, color.Black, Green}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	file, err := os.OpenFile("./ch1/1-5/out.gif", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
	}
	lissajous(file)
}

func lissajous(out io.Writer) {
	const (
		cycle  = 5
		res    = 0.001
		size   = 100
		frames = 64
		delay  = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: frames}
	phase := 0.0
	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}
