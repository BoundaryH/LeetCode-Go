package main

import "fmt"

/* 20ms */
func pivotIndex(nums []int) int {
	rSum, lSum := 0, 0
	for _, n := range nums {
		rSum += n
	}
	for i, n := range nums {
		rSum -= n
		if lSum == rSum {
			return i
		}
		lSum += n
	}
	return -1
}

func main() {
	var nums []int

	nums = []int{1, 7, 3, 6, 5, 6}
	fmt.Println(nums)
	fmt.Println(pivotIndex(nums))

	nums = []int{1, 2, 3}
	fmt.Println(nums)
	fmt.Println(pivotIndex(nums))
}
