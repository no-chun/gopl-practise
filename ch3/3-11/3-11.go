package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-123456"))
}

func comma(s string) string {
	if s == "" {
		return ""
	}
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	var dot = strings.LastIndex(s, ".")
	if dot == -1 {
		dot = len(s)
	}
	for i := 0; i < dot; i++ {
		if (dot-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	buf.WriteString(s[dot:])
	return buf.String()
}
