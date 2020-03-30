package main

// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq(s string, counter map[rune]int) {
	for _, r := range s {
		counter[r]++
	}
}

func main() {
	fname := os.Args[1]
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	counter := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		counter[word]++
	}
	fmt.Printf("%v\n", counter)
}
