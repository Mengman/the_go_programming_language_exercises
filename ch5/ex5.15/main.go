package main

// Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these
// functions do when called with no arguments? Write variants that require at least one argument.

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	strs := flag.Args()

	if len(strs) < 1 {
		fmt.Fprintln(os.Stderr, "input value number must greater or equal than 1")
		os.Exit(1)
	}
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
	if len(vals) == 0 {
		panic("max(vals ...int) vals length must be greater or equal than 1")
	}

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
	if len(vals) == 0 {
		panic("min(vals ...int) vals length must be greater or equal than 1")
	}

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
