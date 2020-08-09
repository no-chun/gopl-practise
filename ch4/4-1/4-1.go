package main

import (
	"crypto/sha256"
	"fmt"
)

func countDiff(a, b [32]uint8) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		cnt = cnt + bitDiff(a[i], b[i])
	}
	return cnt
}

func bitDiff(a uint8, b uint8) int {
	cnt := 0
	for i := uint(0); i < 8; i++ {
		mask := byte(1 << i)
		if a&mask != b&mask {
			cnt++
		}
	}
	return cnt
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("Difference bit : %d", countDiff(c1, c2))
}
