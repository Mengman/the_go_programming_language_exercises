package main

import (
	"sort"
	"fmt"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

	s := []string{"a", "b", "c", "c", "b", "a"}
	ss := sort.StringSlice(s)
	fmt.Printf("IsPalindrome(%v): %v\n", s, IsPalindrome(ss))

	s = []string{"a", "b", "c", "d", "e", "f"}
	ss = sort.StringSlice(s)
	fmt.Printf("IsPalindrome(%v): %v\n", s, IsPalindrome(ss))

}
