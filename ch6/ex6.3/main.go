package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) AddAll(elems ...int) {
	for _, e := range elems {
		s.Add(e)
	}
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	swordNum := len(s.words)
	twordNum := len(t.words)

	for i := swordNum - 1; swordNum-1-i < (swordNum - twordNum); i-- {
		s.words[i] = 0
	}

	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) Difference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			diff := s.words[i] ^ tword
			s.words[i] &= diff
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	s, t := &IntSet{}, &IntSet{}
	s.AddAll(1, 2, 3)
	t.AddAll(2, 3, 4, 5)
	fmt.Printf("s: %v t: %v\n", s, t)
	s.IntersectWith(t)
	fmt.Printf("s intersect with t: %v\n", s)

	s = &IntSet{}
	s.AddAll(1, 2, 3)
	s.Difference(t)
	fmt.Printf("s difference with t: %v\n", s)

	s = &IntSet{}
	s.AddAll(1, 2, 3)
	s.SymmetricDifference(t)
	fmt.Printf("s symmetric difference with t: %v\n", s)
	fmt.Println()

	s, t = &IntSet{}, &IntSet{}
	s.AddAll(2, 3, 4, 5)
	t.AddAll(1, 2, 3)
	fmt.Printf("s: %v t: %v\n", s, t)
	s.IntersectWith(t)
	fmt.Printf("s intersect with t: %v\n", s)

	s = &IntSet{}
	s.AddAll(2, 3, 4, 5)
	s.Difference(t)
	fmt.Printf("s difference with t: %v\n", s)

	s = &IntSet{}
	s.AddAll(2, 3, 4, 5)
	s.SymmetricDifference(t)
	fmt.Printf("s symmetric difference with t: %v\n", s)
	fmt.Println()

	s, t = &IntSet{}, &IntSet{}
	s.AddAll(1, 2, 3)
	fmt.Printf("s: %v t: %v\n", s, t)
	s.IntersectWith(t)
	fmt.Printf("s intersect with t: %v\n", s)

	s = &IntSet{}
	s.AddAll(1, 2, 3)
	s.Difference(t)
	fmt.Printf("s difference with t: %v\n", s)

	s = &IntSet{}
	s.AddAll(1, 2, 3)
	s.SymmetricDifference(t)
	fmt.Printf("s symmetric difference with t: %v\n", s)
	fmt.Println()

	s, t = &IntSet{}, &IntSet{}
	t.AddAll(1, 2, 3)
	fmt.Printf("s: %v t: %v\n", s, t)
	s.IntersectWith(t)
	fmt.Printf("s intersect with t: %v\n", s)

	s = &IntSet{}
	s.Difference(t)
	fmt.Printf("s difference with t: %v\n", s)

	s = &IntSet{}
	s.SymmetricDifference(t)
	fmt.Printf("s symmetric difference with t: %v\n", s)
	fmt.Println()

	s, t = &IntSet{}, &IntSet{}
	fmt.Printf("s: %v t: %v\n", s, t)
	s.IntersectWith(t)
	fmt.Printf("s intersect with t: %v\n", s)

	s = &IntSet{}
	s.Difference(t)
	fmt.Printf("s difference with t: %v\n", s)

	s = &IntSet{}
	s.SymmetricDifference(t)
	fmt.Printf("s symmetric difference with t: %v\n", s)
	fmt.Println()
}
