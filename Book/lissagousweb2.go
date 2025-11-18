package Book

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palettes = map[string]color.Color{
	"green": color.RGBA{0x00, 0xff, 0x00, 0xff},
	"red":   color.RGBA{0xff, 0x00, 0x00, 0xff},
	"blue":  color.RGBA{0x00, 0x00, 0xff, 0xff},
}

func Bain() {
	http.HandleFunc("/lissajous", Handler)
	http.ListenAndServe(":8000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Read URL parameters
	cycles := getIntParam(r, "cycles", 5)
	size := getIntParam(r, "size", 100)
	delay := getIntParam(r, "delay", 8)
	colorName := r.URL.Query().Get("color")

	// Build palette dynamically
	palette := []color.Color{color.White}
	col, ok := palettes[colorName]
	if !ok {
		col = palettes["green"] // default
	}
	palette = append(palette, col)

	w.Header().Set("Content-Type", "image/gif")
	lissajous(w, cycles, size, delay, palette)
}

func getIntParam(r *http.Request, name string, def int) int {
	v := r.URL.Query().Get(name)
	if v == "" {
		return def
	}
	val, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return val
}

func lissajous(out http.ResponseWriter, cycles, size, delay int, palette []color.Color) {
	const (
		res     = 0.001
		nframes = 64
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)), size+int(y*float64(size)), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
