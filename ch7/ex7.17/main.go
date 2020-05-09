package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	className := os.Args[1]

	dec := xml.NewDecoder(os.Stdin)
	var stack []string

	selected := false

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)

			for _, attr := range tok.Attr {
				if attr.Name.Local == "class" && attr.Value == className {
					selected = true
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if selected {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				selected = false
			}
		}
	}
}
