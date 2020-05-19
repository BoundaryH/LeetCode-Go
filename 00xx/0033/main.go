package main

import "fmt"

/* 0ms */
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}
		if nums[m] < nums[r] {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}

func main() {
	var nums []int
	var target int

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	fmt.Println(nums)
	fmt.Println(target, search(nums, target))

	nums = []int{5, 1, 3}
	target = 5
	fmt.Println(nums)
	fmt.Println(target, search(nums, target))
}
