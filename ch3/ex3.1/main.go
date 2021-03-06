package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "please input svg file name\n")
		os.Exit(1)
	}

	fname := os.Args[1]
	if !strings.HasSuffix(fname, ".svg") {
		fname += ".svg"
	}
	f, err := os.Create(fname)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			bx, by, err := corner(i, j)
			cx, cy, err := corner(i, j+1)
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			f.WriteString(fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	f.WriteString(fmt.Sprintf("</svg>"))
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, errors.New("invalid number")
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
