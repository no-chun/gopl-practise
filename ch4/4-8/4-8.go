package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

const (
	letter  string = "letter"
	number  string = "number"
	graphic string = "graphic"
	space   string = "space"
	symbol  string = "symbol"
)

func main() {
	count := make(map[string]int, 5)
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		switch {
		case unicode.IsLetter(r):
			count[letter]++
		case unicode.IsNumber(r):
			count[number]++
		case unicode.IsGraphic(r):
			count[graphic]++
		case unicode.IsSpace(r):
			count[space]++
		case unicode.IsSymbol(r):
			count[symbol]++
		}
	}
	for class, cnt := range count {
		fmt.Printf("class: %s, count = %d\n", class, cnt)
	}
}
