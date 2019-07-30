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
	colorSteps    = 50
	red           = 0xff0000
	blue          = 0x0000ff
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

	maxZ := 0.0
	minZ := 0.0
	cellList := make([][][]float64, cells)
	for i := 0; i < cells; i++ {
		cellList[i] = make([][]float64, cells)
		for j := 0; j < cells; j++ {
			ax, ay, az, err := corner(i+1, j, graphName)
			bx, by, bz, err := corner(i, j, graphName)
			cx, cy, cz, err := corner(i, j+1, graphName)
			dx, dy, dz, err := corner(i+1, j+1, graphName)
			if err != nil {
				continue
			}
			z := (az + bz + cz + dz) / 4
			if z > maxZ {
				maxZ = z
			} else if z < minZ {
				minZ = z
			}

			cellList[i][j] = append(cellList[i][j], []float64{ax, ay, bx, by, cx, cy, dx, dy, z}...)
		}
	}

	stepLen := (maxZ - minZ) / colorSteps
	stepR, stepG, stepB := gradientColorStep(blue, red, colorSteps)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := cellList[i][j][8]
			n := int((z - minZ) / stepLen)
			r := stepR * n
			g := stepG * n
			b := stepB * n
			color := (r << 16) + (g << 8) + b
			f.WriteString(fmt.Sprintf("<polygon fill='#%06X' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				color, cellList[i][j][0], cellList[i][j][1], cellList[i][j][2], cellList[i][j][3], cellList[i][j][4],
				cellList[i][j][5], cellList[i][j][6], cellList[i][j][7]))

		}
	}

	f.WriteString(fmt.Sprintf("</svg>"))
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

func corner(i, j int, graphName string) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	var f zFunc
	switch graphName {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	default:
		fmt.Fprintln(os.Stderr, "unknown graph name")
		os.Exit(1)
	}
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, 0, errors.New("invalid number")
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}
