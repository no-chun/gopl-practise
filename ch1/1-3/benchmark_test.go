package main

import (
	"os"
	"strings"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args, " ")
	}
}
