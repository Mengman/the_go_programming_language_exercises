package main
// Exercise 1.3: Experiment to measure the difference in running time between our 
// potentially inefficient versions and the one that uses strings.Join.

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(time.Now().Sub(start))

}