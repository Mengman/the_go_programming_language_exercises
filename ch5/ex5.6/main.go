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

type zFunc func(x, y float64) float64

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "please input svg file name and graph name [eggbox, saddle]\n")
		os.Exit(1)
	}

	fname := os.Args[1]
	graphName := os.Args[2]
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

	var zf zFunc
	switch graphName {
	case "saddle":
		zf = saddle
	case "eggbox":
		zf = eggbox
	default:
		fmt.Fprintln(os.Stderr, "unknown graph name")
		os.Exit(1)
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j, zf)
			bx, by, err := corner(i, j, zf)
			cx, cy, err := corner(i, j+1, zf)
			dx, dy, err := corner(i+1, j+1, zf)
			if err != nil {
				continue
			}
			f.WriteString(fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	f.WriteString(fmt.Sprintf("</svg>"))
}

func corner(i, j int, f zFunc) (sx, sy float64, err error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, errors.New("invalid number")
	}

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return y*y/a2 - x*x/b2
}
