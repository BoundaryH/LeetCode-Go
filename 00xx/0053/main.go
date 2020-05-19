package main

import "fmt"

/* 4ms */
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0
	for _, n := range nums {
		if sum > 0 {
			sum += n
		} else {
			sum = n
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	var nums []int

	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(nums)
	fmt.Println(maxSubArray(nums))
}
