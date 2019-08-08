package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func countSHA256Diff(a, b string) int {
	c1 := sha256.Sum256([]byte(a))
	c2 := sha256.Sum256([]byte(b))
	diffNum := 0
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			diffNum++
		}
	}
	return diffNum
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	diffNum := countSHA256Diff(a, b)
	fmt.Println(fmt.Sprintf("SHA256 of '%v' and '%v' has %d different bits", a, b, diffNum))
}
