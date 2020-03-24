package main

// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256
// hashes. (See PopCount from Section 2.6.2.)

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}

	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i >= len(a) || i >= len(b) {
			count += 8
		} else {
			count += popCount(a[i] ^ b[i])
		}
	}

	return count
}

func countSHA256Diff(a, b string) int {
	c1 := sha256.Sum256([]byte(a))
	c2 := sha256.Sum256([]byte(b))

	return bitDiff(c1[:], c2[:])
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	diffNum := countSHA256Diff(a, b)
	fmt.Println(fmt.Sprintf("SHA256 of '%v' and '%v' has %d different bits", a, b, diffNum))
}
