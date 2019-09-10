package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	strs := flag.Args()
	var numbers []int
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "input '%s' is not a number\n", s)
			os.Exit(1)
		}
		numbers = append(numbers, n)
	}
	maxN := max(numbers...)
	minN := min(numbers...)
	fmt.Fprintf(os.Stdout, "Max number %d\nMin number %d\n", maxN, minN)
}

func max(vals ...int) int {
	if len(vals) < 2 {
		return vals[0]
	}

	m := vals[0]
	for _, v := range vals[1:] {
		if m < v {
			m = v
		}
	}

	return m
}

func min(vals ...int) int {
	if len(vals) < 2 {
		return vals[0]
	}

	m := vals[0]
	for _, v := range vals[1:] {
		if m > v {
			m = v
		}
	}
	return m
}
