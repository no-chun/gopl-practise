package test

import (
	"gopl-practise/ch2/popcount"
	"testing"
)

func PopCountByShifting(x uint64) int {
	cnt := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			cnt++
		}
	}
	return cnt
}

func PopCountByClearing(x uint64) int {
	cnt := 0
	for x != 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount1(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}
