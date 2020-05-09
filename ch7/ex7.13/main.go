package main

import (
	"github.com/Mengman/the_go_programming_language_exercises/ch7/ex7.13/eval"
	"log"
)

func main() {
	var inputs = []string{
		"log(10)",
		"sqrt(A / pi)",
		"5 / 9 * pow(x, 3)",
	}

	for _, input := range inputs {
		expr, err := eval.Parse(input)
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("input: %s, eval.String(): %v", input, expr)
	}
}
