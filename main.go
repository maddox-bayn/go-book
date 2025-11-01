package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// Define our color palette - the only colors we can use in the GIF
var palette = []color.Color{
	color.White, // Index 0 - background
	color.Black, // Index 1 - the curve
}

const (
	whiteIndex = 0 // Background color
	blackIndex = 1 // Drawing color
)

func main() {
	// Seed random number generator so we get different patterns each time
	rand.Seed(time.Now().UTC().UnixNano())

	// Create a file to save the GIF
	f, err := os.Create("lissajous.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Generate the animation and write to file
	lissajous(f)
}

// lissajous generates an animated GIF of a Lissajous curve
// It can write to any io.Writer (file, stdout, network connection, etc.)
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of complete oscillations
		res     = 0.001 // Angular resolution (smaller = smoother curve)
		size    = 100   // Canvas goes from -size to +size
		nframes = 64    // Number of frames in animation
		delay   = 8     // Delay between frames in 10ms units (80ms total)
	)

	// Random frequency creates variety in curve patterns
	// Multiplying by 3.0 gives frequencies between 0 and 3
	freq := rand.Float64() * 3.0

	// Create GIF structure to hold all frames
	anim := gif.GIF{LoopCount: nframes}

	// Phase controls the animation - it shifts the Y wave
	phase := 0.0

	// Generate each frame of the animation
	for i := 0; i < nframes; i++ {
		// Create a new blank image for this frame
		// Rectangle from (0,0) to (201,201)
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Draw the Lissajous curve for this frame
		// t represents the angle parameter (goes through 5 complete circles)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// X oscillates with sine of t
			x := math.Sin(t)

			// Y oscillates with sine of t*freq, shifted by phase
			// freq changes the frequency ratio between X and Y
			// phase creates the animation by shifting the pattern
			y := math.Sin(t*freq + phase)

			// Convert from mathematical coordinates (-1 to 1) to pixels
			// 1. Multiply by size to get -100 to 100
			// 2. Add 0.5 for rounding
			// 3. Convert to int
			// 4. Add size to shift to 0 to 200 range
			xPixel := size + int(x*float64(size)+0.5)
			yPixel := size + int(y*float64(size)+0.5)

			// Draw a black pixel at this position
			img.SetColorIndex(xPixel, yPixel, blackIndex)
		}

		// Increment phase for next frame (creates the rotation effect)
		phase += 0.1

		// Add this frame to the animation
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Encode all frames as an animated GIF and write to output
	gif.EncodeAll(out, &anim)
}
