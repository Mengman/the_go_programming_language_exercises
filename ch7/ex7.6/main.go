package main

// Exercis e 7.6: Add support for Kelvin temperatures to tempflag.

import (
	"fmt"
	"flag"
	"github.com/Mengman/the_go_programming_language_exercises/ch7/ex7.6/tempconv"
)

var celsius = tempconv.CelsiusFlag("celsius", 20.0, "the celsius temperature")
var kelvin = tempconv.KelvinFlag("kelvin", 20.0, "the kelvin temperature")

func main() {
	flag.Parse()
	
	fmt.Println(*celsius)
	fmt.Println(*kelvin)
}