package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func commaFloat(s string) (string, error) {
	splits := strings.Split(s, ".")
	if len(splits) > 2 {
		return "", fmt.Errorf("not a float number")
	} else if len(splits) < 1 {
		return s, nil
	}

	intPart := comma(splits[0])
	decimalPart := ""
	if len(splits) > 1 {
		decimalPart = "." + comma(splits[1])
	}
	return intPart + decimalPart, nil
}

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
	output, err := commaFloat(*s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(output)
	}
}
