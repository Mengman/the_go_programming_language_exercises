package main

import (
	"flag"
	"fmt"
)

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	out := ""
	n := 0
	for ; n+3 <= len(s); n += 3 {
		out += s[n:n+3] + ","
	}
	out += s[n:]
	return out
}

func main() {
	s := flag.String("num", "1234567", "input number string")
	flag.Parse()
	fmt.Println(comma(*s))
}
