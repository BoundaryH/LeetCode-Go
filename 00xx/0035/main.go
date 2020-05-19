package main

import "fmt"

/* 4ms */
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	var nums []int
	var target int

	nums = []int{1, 3, 5, 6}
	target = 5
	fmt.Println(nums, target, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(nums, target, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(nums, target, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 0
	fmt.Println(nums, target, searchInsert(nums, target))
}
