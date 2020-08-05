package test

import (
	"gopl-practise/ch2/popcount"
	"testing"
)

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
