package main

// Exercise 3.10: Write a non-rec ursive version of comma, using bytes.Buffer instead of string concatenation.

import (
	"bytes"
	"flag"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	var digitals []string
	if len(s) <= 3 {
		return s
	}

	n := len(s)
	for ; n-3 >= 0; n -= 3 {
		digitals = append(digitals, s[n-3:n])
		if n-3 != 0 {
			digitals = append(digitals, ",")
		}
	}

	if n > 0 {
		digitals = append(digitals, s[0:n])
	}

	for i := len(digitals) - 1; i >= 0; i-- {
		buf.WriteString(digitals[i])
	}

	return buf.String()
}

func main() {
	s := flag.String("num", "1234567", "input number string")
	flag.Parse()
	fmt.Println(comma(*s))
}
