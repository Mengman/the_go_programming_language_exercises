package main

// Exercis e 3.12: Write a function that reports whether two str ings are anagrams of each other,
// that is, the y contain the same letters in a different order.

import (
	"flag"
	"fmt"
)

func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	rune1 := []rune(s1)
	rune2 := []rune(s2)

	for i, r1 := range rune1 {
		r2 := rune2[len(rune2)-1-i]

		if r1 != r2 {
			return false
		}
	}

	return true
}

func main() {
	s1 := flag.String("s1", "", "s1 string")
	s2 := flag.String("s2", "", "s2 string")
	flag.Parse()
	if isAnagrams(*s1, *s2) {
		fmt.Println(fmt.Sprintf("'%s' and '%s' are anagrams", *s1, *s2))
	} else {
		fmt.Println(fmt.Sprintf("'%s' and '%s' are not anagrams", *s1, *s2))
	}
}
