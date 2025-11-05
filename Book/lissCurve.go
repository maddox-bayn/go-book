package Book

import (
	"image"
	"image/color"
	"image/png" // Use PNG for simplicity first
	"math"
	"os"
)

func Lcurve() {
	// Create canvas
	img := image.NewRGBA(image.Rect(0, 0, 201, 201))

	// Fill white background
	for y := 0; y < 201; y++ {
		for x := 0; x < 201; x++ {
			img.Set(x, y, color.White)
		}
	}

	// Draw one Lissajous curve
	freq := 2.0
	for t := 0.0; t < 5*2*math.Pi; t += 0.01 {
		x := math.Sin(t)
		y := math.Sin(t * freq)

		xPixel := 100 + int(x*90)
		yPixel := 100 + int(y*90)

		img.Set(xPixel, yPixel, color.Black)
	}

	// Save as PNG
	f, _ := os.Create("curve.png")
	png.Encode(f, img)
	f.Close()
}
