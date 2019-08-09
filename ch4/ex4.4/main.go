package main

import (
	"fmt"
	"os"
	"strconv"
)

func rotate(s []int, n int) []int {
	n = n % len(s)
	return append(s[n:], s[:n]...)
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v rotate %d %v\n", s, n, rotate(s, int(n)))
}
