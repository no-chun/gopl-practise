package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	s := "hello\n\nworld\n\n"
	var wc WordCount
	_, _ = fmt.Fprintf(&wc, s)
	fmt.Println(wc)
	var lc LineCount
	_, _ = fmt.Fprintf(&lc, s)
	fmt.Println(lc)
}

type WordCount int

type LineCount int

func (c *WordCount) Write(w []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(w))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return len(w), nil
}

func (c *LineCount) Write(w []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(w))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(w), nil
}
