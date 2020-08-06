package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}

func comma(s string) string {
	if s == "" {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
