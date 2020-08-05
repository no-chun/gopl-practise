package main

import (
	"bufio"
	"fmt"
	"gopl-practise/ch2/2-1"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			conv(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	}
}

func conv(temp string) {
	t, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.FToC(f)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}
