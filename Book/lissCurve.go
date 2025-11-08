package Book

import (
	"fmt"
	"image"
	"image/color"
	"image/png" // Use PNG for simplicity first
	"math"
	"os"
)

func Lcurve() {
	phase := 0.0

	for frame := 0; frame < 5; frame++ {
		img := image.NewRGBA(image.Rect(0, 0, 201, 201))

		// White background
		for y := 0; y < 201; y++ {
			for x := 0; x < 201; x++ {
				img.Set(x, y, color.White)
			}
		}

		// Draw curve WITH PHASE
		freq := 2.0
		for t := 0.0; t < 5*2*math.Pi; t += 0.01 {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase) // <- Added phase!

			xPixel := 100 + int(x*90)
			yPixel := 100 + int(y*90)

			img.Set(xPixel, yPixel, color.Black)
		}

		phase += 76 // Increment for next frame

		f, _ := os.Create(fmt.Sprintf("frame%d.png", frame))
		png.Encode(f, img)
		f.Close()
	}
}
