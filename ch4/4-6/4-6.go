package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func removeSpace(b []byte) []byte {
	for i := 0; i < len(b); {
		str, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(str) {
			str2, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(str2) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}
	return b
}

func main() {
	b := []byte("a b  c  d e  f")
	b = removeSpace(b)
	fmt.Println(string(b))
}
