package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}
	breathFirst(keys)
}

func breathFirst(keys []string) {
	seen := make(map[string]bool)
	for len(keys) > 0 {
		items := keys
		keys = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				fmt.Printf("%s\n", item)
				keys = append(keys, prereqs[item]...)
			}
		}
	}
}
