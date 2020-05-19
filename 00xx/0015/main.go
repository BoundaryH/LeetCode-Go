package main

import (
	"fmt"
	"sort"
)

/* 36ms */
func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	size := len(nums)
	for i := 0; i < size; {
		for l, r := i+1, size-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			}
			if sum < 0 {
				for n := nums[l]; l < r && nums[l] == n; l++ {
				}
			} else {
				for n := nums[r]; l < r && nums[r] == n; r-- {
				}
			}
		}
		for n := nums[i]; i < size && nums[i] == n; i++ {
		}
	}
	return
}

func main() {
	var nums []int

	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(nums)
	for _, r := range threeSum(nums) {
		fmt.Println("   ", r)
	}

	nums = []int{-1, 0, 1, 0}
	fmt.Println(nums)
	for _, r := range threeSum(nums) {
		fmt.Println("   ", r)
	}

	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(nums)
	for _, r := range threeSum(nums) {
		fmt.Println("   ", r)
	}
}
