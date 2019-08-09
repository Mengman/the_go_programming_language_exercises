package main

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
