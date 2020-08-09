package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		count[input.Text()]++
	}
	for word, cnt := range count {
		fmt.Printf("word: %s, count= %d\n", word, cnt)
	}
}
