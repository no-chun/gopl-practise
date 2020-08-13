package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("foo---foo----foo", f))
}

func expand(s string, f func(string) string) string {
	for i := strings.Index(s, "foo"); i != -1; i = strings.Index(s, "foo") {
		s = s[:i] + f("foo") + s[i+3:]
	}
	return s
}

func f(s string) string {
	return "***"
}
