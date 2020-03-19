package main

// Exercise 3.10: Write a non-rec ursive version of comma, using bytes.Buffer instead of string concatenation.

import (
	"bytes"
	"flag"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	if len(s) <= 3 {
		return s
	}

	n := 0
	for ; n+3 <= len(s); n += 3 {
		if n != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[n : n+3])
	}

	if n < len(s)-1 {
		buf.WriteString(",")
		buf.WriteString(s[n:])
	}

	return buf.String()
}

func main() {
	s := flag.String("num", "1234567", "input number string")
	flag.Parse()
	fmt.Println(comma(*s))
}
