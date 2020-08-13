package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"a", "b", "c", "d", "e"}
	fmt.Println(join("--", strs...))
}

func join(sep string, strs ...string) string {
	var builder strings.Builder
	for i, s := range strs {
		if i == 0 {
			builder.WriteString(s)
		} else {
			builder.WriteString(sep)
			builder.WriteString(s)
		}
	}
	return builder.String()
}
