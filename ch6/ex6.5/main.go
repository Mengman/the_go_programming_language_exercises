package main

import (
	"bytes"
	"fmt"
)

const bitSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var num int
	for _, word := range s.words {
		num += popCount(word)
	}
	return num
}

func popCount(x uint) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return int(count)
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/bitSize, uint(x%bitSize)
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

func (s *IntSet) Copy() *IntSet {
	words := make([]uint, len(s.words))
	copy(words, s.words)
	return &IntSet{
		words: words,
	}
}

func main() {
	fmt.Printf("This is a %d-bit platform\n", bitSize)
	s1 := IntSet{}
	v := 3
	fmt.Printf("s has value %d %v\n", v, s1.Has(v))
	fmt.Printf("add %d to set\n", v)
	s1.Add(v)
	fmt.Printf("s has value %d %v s: %v\n ", v, s1.Has(v), &s1)

	s2 := IntSet{}
	s2.Add(0)
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	fmt.Printf("s1: %v, s2: %v\n", &s1, &s2)
	s1.UnionWith(&s2)
	fmt.Printf("s1 union with s2: %v, len: %d\n", &s1, s1.Len())

	s1Copy := *s1.Copy()
	s1.Remove(0)
	s1.Remove(1)
	s1.Remove(100)
	fmt.Printf("s1: %v; s1Copy: %v\n", &s1, &s1Copy)
	s1Copy.Clear()
	fmt.Printf("Clear s1Copy: %v\n", &s1Copy)
}
