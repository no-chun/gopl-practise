package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
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

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for mask := uint(0); mask < 64; mask++ {
			if word&(1<<mask) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	n := &IntSet{}
	copy(n.words, s.words)
	return n
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	Len1 := s.Len()
	Len2 := t.Len()
	var Len int
	if Len1 < Len2 {
		Len = Len1
	} else {
		Len = Len2
	}
	for i := 0; i < Len; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:Len]
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
	for len(s.words) != 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
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
	for len(s.words) != 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

func (s *IntSet) Elems() []int {
	var elems []int
	x := 0
	for _, word := range s.words {
		for mask := uint(0); mask < 64; mask++ {
			if word&(1<<mask) != 0 {
				elems = append(elems, x)
				x++
			}
		}
	}
	return elems
}
