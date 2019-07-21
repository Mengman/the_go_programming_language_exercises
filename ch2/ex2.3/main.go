package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Mengman/the_go_programming_language_exercises/ch2/ex2.3/popcount"
)

func main() {
	arg := os.Args[1]
	t, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	p := popcount.PopCount(t)
	fmt.Printf("%d %d\n", t, p)

}
