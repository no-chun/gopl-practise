package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (res int) {
	defer func() {
		res = 100
		_ = recover()
	}()
	panic("error")
}
