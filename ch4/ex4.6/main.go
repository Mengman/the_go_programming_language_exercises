package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpace(bytes []byte) []byte {
	lastSpace := false
	sHead := 0
	for sTail := 0; sTail < len(bytes); {
		r, size := utf8.DecodeRune(bytes[sTail:])
		sTail += size
		if !unicode.IsSpace(r) {
			utf8.EncodeRune(bytes[sHead:], r)
			sHead += size
			lastSpace = false
		} else {
			// if last rune is not a space, sHead move to next rune
			if !lastSpace {
				s := utf8.EncodeRune(bytes[sHead:], r)
				sHead += s
			}
			lastSpace = true
		}
	}
	return bytes[:sHead]
}

func main() {
	s := "我   会 写   Go 语言"
	fmt.Printf("%v\n", []byte(s))
	fmt.Printf("%v\n", []byte(squashSpace([]byte(s))))
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", squashSpace([]byte(s)))
}
