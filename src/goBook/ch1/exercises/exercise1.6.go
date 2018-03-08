package main

import (
	"os"
	"io"
	"image/color"
	"math/rand"
	"image/gif"
	"image"
	"math"
)

var paletteTwo = []color.Color{
	color.Black, color.RGBA{0x55, 0xff, 0x34, 0xff},
	color.RGBA{0xff, 0x00, 0x12, 0xff}, color.RGBA{0x12, 0x12, 0xff, 0xff},
	color.RGBA{0xAA, 0x59, 0x99, 0xff}, color.RGBA{0x60, 0x70, 0x50, 0xff}}

func main() {
	lissajousTwo(os.Stdout)
	lissajousTwo(os.Stdout)
	lissajousTwo(os.Stdout)
}

func lissajousTwo(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletteTwo)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Int31n(5)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
