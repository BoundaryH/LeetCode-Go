package main

import "fmt"

/* 4ms */
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return true
		}
		if nums[m] == nums[r] {
			r--
		} else if nums[m] < nums[r] {
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
	return false
}

func main() {
	var nums []int
	var target int

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target = 0
	fmt.Println(nums)
	fmt.Println(target, search(nums, target))

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	fmt.Println(nums)
	fmt.Println(target, search(nums, target))
}
