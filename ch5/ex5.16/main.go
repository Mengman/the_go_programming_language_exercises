package main

// Exercise 5.16: Write a variadic version of strings.Join.

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	strs := flag.Args()
	s := join(strs...)
	fmt.Fprintln(os.Stdout, s)
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(s)
	}
	return sb.String()
}
