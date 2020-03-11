package main

// Exercise 1.12: Modify the Lissajous server to read parameter values from the URL. For example,
// you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
// number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the
// string parameter into an integer. You can see its document ation with go doc strconv.Atoi.

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Printf("%s %s\n", k, v)
		}
		genLissajous(w, r.Form)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func genLissajous(out io.Writer, params url.Values) {
	cycles := 5
	res := 0.001
	size := 100
	nframes := 64
	delay := 8

	if val, ok := params["cycles"]; ok && len(val) > 0 {
		s, err := strconv.ParseInt(val[0], 10, 32)
		if err != nil {
			cycles = int(s)
		}
	}

	if val, ok := params["res"]; ok && len(val) > 0 {
		s, err := strconv.ParseFloat(val[0], 32)
		if err != nil {
			res = s
		}
	}

	if val, ok := params["size"]; ok && len(val) > 0 {
		s, err := strconv.ParseInt(val[0], 10, 32)
		if err != nil {
			size = int(s)
		}
	}

	if val, ok := params["nframes"]; ok && len(val) > 0 {
		s, err := strconv.ParseInt(val[0], 10, 32)
		if err != nil {
			nframes = int(s)
		}
	}

	if val, ok := params["delay"]; ok && len(val) > 0 {
		s, err := strconv.ParseInt(val[0], 10, 32)
		if err != nil {
			delay = int(s)
		}
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
