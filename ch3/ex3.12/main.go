package main

import (
	"flag"
	"fmt"
)

func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]bool)
	for _, l := range s1 {
		m[l] = true
	}

	for _, l := range s2 {
		if _, ok := m[l]; !ok {
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
