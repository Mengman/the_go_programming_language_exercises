package main

// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.

import (
	"fmt"
	"os"
	"unicode"
)

func countSymbol(s string) map[string]int {
	counter := map[string]int{
		"letter": 0,
		"digit":  0,
		"other":  0,
	}

	for _, r := range s {
		if unicode.IsLetter(r) {
			counter["letter"]++
		} else if unicode.IsDigit(r) {
			counter["digit"]++
		} else {
			counter["other"]++
		}
	}

	return counter
}

func main() {
	s := os.Args[1]
	fmt.Printf("%v \n", countSymbol(s))
}
