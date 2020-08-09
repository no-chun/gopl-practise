package main

import "fmt"

func reverse(num *[5]int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}

func main() {
	num := [5]int{1, 2, 3, 4, 5}
	reverse(&num)
	fmt.Println(num)
}
