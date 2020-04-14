package main

// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return count, err
	}
	*wc += WordCounter(count)
	return count, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return count, err
	}

	*lc += LineCounter(count)
	return count, nil
}

const text = `Do not go gentle into that good night,
Old age should burn and rave at close of day;
Rage, rage against the dying of the light.

Though wise men at their end know dark is right,
Because their words had forked no lightning they
Do not go gentle into that good night.

Good men, the last wave by, crying how bright
Their frail deeds might have danced in a green bay,
Rage, rage against the dying of the light.

Wild men who caught and sang the sun in flight,
And learn, too late, they grieved it on its way,
Do not go gentle into that good night.

Grave men, near death, who see with blinding sight
Blind eyes could blaze like meteors and be gay,
Rage, rage against the dying of the light.

And you, my father, there on that sad height,
Curse, bless, me now with your fierce tears, I pray.
Do not go gentle into that good night.
Rage, rage against the dying of the light.`

func main() {
	var wc WordCounter
	var lc LineCounter
	_, err := wc.Write([]byte(text))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = lc.Write([]byte(text))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(text)
	fmt.Printf("line count: %d word count: %d\n", lc, wc)
}
