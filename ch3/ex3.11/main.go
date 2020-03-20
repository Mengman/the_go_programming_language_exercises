package main

// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func commaFloat(s string) (string, error) {
	splits := strings.Split(s, ".")
	if len(splits) > 2 {
		return "", fmt.Errorf("not a float number")
	} else if len(splits) < 1 {
		return s, nil
	}

	var sign string
	intPart := splits[0]

	decimalPart := ""
	if len(splits) > 1 {
		decimalPart = "." + reverseString(comma(reverseString(splits[1])))
	}

	if strings.HasPrefix(intPart, "+") || strings.HasPrefix(intPart, "-") {
		sign = intPart[:1]
		intPart = intPart[1:]
	}

	intPart = sign + comma(intPart)

	return intPart + decimalPart, nil
}

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
	output, err := commaFloat(*s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(output)
	}
}
