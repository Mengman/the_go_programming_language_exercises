package main
// Exercise 1.4: Modify dup2 to print the names of all file3s in which each duplicated line occurs.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: $v\n", err)
				continue
			}
			countLines(f, counts, arg, lineFiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, lineFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, lineFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		lineFiles[input.Text()] = append(lineFiles[input.Text()], filename)
	}
}
