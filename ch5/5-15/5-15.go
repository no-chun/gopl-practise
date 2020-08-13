package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(max(1, 2, 3, 9, 5, 6))
	fmt.Println(min(2, 4, 5, 1, 7, 9))
	fmt.Println(max_(1, 2, 3, 9, 5, 6))
	fmt.Println(min_(2, 4, 5, 1, 7, 9))
}

func max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("there are no arguments！")
	}
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum, nil
}

func min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("there are no arguments！")
	}
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum, nil
}

func max_(first int, rest ...int) int {
	maxNum := first
	for _, num := range rest {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func min_(first int, rest ...int) int {
	minNum := first
	for _, num := range rest {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}
