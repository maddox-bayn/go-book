package Book

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var palette1 = []color.Color{color.White, color.Black}

const (
	whiteIndex1 = 0
	BlackIndex1 = 1
)

func Master() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	LLissajous(w)
}

func LLissajous(out io.Writer) {

	const (
		cycles  = 5
		res     = 0.001
		nframes = 64
		size    = 100
		delay   = 8
	)

	freq := rand.Float64() * 3.0

	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			xPixel := size + int(x*float64(size)+0.5)
			yPixel := size + int(y*float64(size)+0.5)

			img.SetColorIndex(xPixel, yPixel, BlackIndex1)
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
