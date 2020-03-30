package main

// Exercise 5.9: Write a function expand(s string, f func(string) string) string that
// replaces each substring ‘‘$foo’’ within s by the text returned by f("foo").

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s, func(ph string) string {
		return f(ph[1:])
	})
}

func main() {
	subs := make(map[string]string, 0)
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			kv := strings.Split(arg, "=")
			if len(kv) != 2 {
				fmt.Fprintf(os.Stderr, "reqired param format K=VAL got %s\n", arg)
				os.Exit(1)
			}

			subs[kv[0]] = kv[1]
		}
	}

	f := func(name string) string {
		v, ok := subs[name]
		if ok {
			return v
		} else {
			return "$" + name
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		fmt.Print(expand(input, f))
	}
}
