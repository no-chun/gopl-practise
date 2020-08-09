package main

import "fmt"

func rotate(nums []int, n int) {
	n %= len(nums)
	tmp := append(nums, nums[:n]...)
	copy(nums, tmp[n:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(nums, 2)
	fmt.Println(nums)
}
