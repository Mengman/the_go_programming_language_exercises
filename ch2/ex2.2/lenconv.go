package main

// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there are no arguments,
// and converts each number into units like temperature in Celsius and Fahren heit,
// length in feet and meters, weight in pounds and kilograms, and the like.

import (
	"fmt"
	"os"
	"strconv"
)

type Inch float64
type Centimetre float64

func (i Inch) String() string {
	return fmt.Sprintf("%.3finch", i)
}

func (c Centimetre) String() string {
	return fmt.Sprintf("%.3fcm", c)
}

func Inch2Centimetre(i Inch) Centimetre {
	return Centimetre(i * 2.54)
}

func Centimetre2Inch(c Centimetre) Inch {
	return Inch(c / 2.54)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "lenconv: %v\n", err)
			os.Exit(1)
		}

		c := Centimetre(t)
		i := Centimetre2Inch(c)

		fmt.Printf("%s = %s\n", c, i)
	}
}
