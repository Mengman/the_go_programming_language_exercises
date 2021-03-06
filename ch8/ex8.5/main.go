package main

// Exercise 8.5: Take an existing CPU-bound sequential program, such as the Mandelbrot
// program of Section 3.3 or the 3-D surface computation of Section 3.2, and execute its main
// loop in parallel using channels for communication. How much faster does it run on a
// multiprocessor machine? What is the optimal number of goroutines to use?
// compare with ex3.5
import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	start := time.Now()
	wg := sync.WaitGroup{}
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)

	worker_num := runtime.GOMAXPROCS(-1)
	for i := 0; i < worker_num; i++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	png.Encode(f, img)
	log.Println("rendered in:", time.Since(start))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const s = 0x8E2DE2
	const e = 0x4A00E0

	sr, sg, sb := gradientColorStep(s, e, iterations)
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := uint8(sr * n)
			g := uint8(sg * n)
			b := uint8(sb * n)
			return color.RGBA{r, g, b, 0xff}
		}
	}
	r, g, b := hex2rgb(e)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}

func gradientColorStep(start int, end int, steps int) (int, int, int) {
	sr, sg, sb := hex2rgb(start)
	er, eg, eb := hex2rgb(end)
	stepR := sr + (er-sr)/steps
	stepG := sg + (eg-sg)/steps
	stepB := sb + (eb-sb)/steps
	return stepR, stepG, stepB
}

func hex2rgb(color int) (int, int, int) {
	r := (color >> 16) & 0xff
	g := (color >> 8) & 0xff
	b := color & 0xff
	return r, g, b
}
