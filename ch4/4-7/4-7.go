package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	b := []byte("1 2 3 一 二 三")
	reverseUTF8(b)
	fmt.Printf("%s\n", b)
}
