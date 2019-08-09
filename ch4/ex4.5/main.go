package main

import "fmt"

func elimDup(strs []string) []string {
	if len(strs) < 2 {
		return strs
	}
	p := 0
	for i := 1; i < len(strs); i++ {
		if strs[p] != strs[i] {
			p++
			if p != i {
				strs[p] = strs[i]
			}
		}
	}
	return strs[:p+1]
}

func main() {
	strs := []string{"a", "b", "c", "c", "c", "d", "d"}
	fmt.Printf("%v\n", elimDup(strs))
}
