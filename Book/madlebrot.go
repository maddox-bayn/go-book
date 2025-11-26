package Book

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func Wain() {
	// Create a 1024×1024 image
	const width, height = 1024, 1024
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Loop over every pixel
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*4 - 2 // map pixel -> y axis

		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*4 - 2 // map pixel -> x axis

			// Complex number c = x + yi
			c := complex(x, y)

			// Compute pixel color
			color := mandelbrot(c)

			// Set the color to the image
			img.Set(px, py, color)
		}
	}

	// Save image to file
	file, _ := os.Create("mandelbrot.png")
	defer file.Close()
	png.Encode(file, img)
}

// mandelbrot function determines the color of a point
func mandelbrot(c complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var z complex128

	for n := 0; n < iterations; n++ {
		z = z*z + c // Mandelbrot formula
		if cmplx.Abs(z) > 2 {
			// Escaped → color based on speed of escape
			return color.Gray{uint8(255 - contrast*n)}
		}
	}
	// Never escaped → black
	return color.Black
}
