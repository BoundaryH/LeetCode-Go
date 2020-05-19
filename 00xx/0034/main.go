package main

import "fmt"

/* 8ms */
func searchRange(nums []int, target int) (res []int) {
	res = []int{-1, -1}
	if len(nums) == 0 {
		return
	}
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if nums[l] != target {
		return
	}
	res[0] = l
	r = len(nums) - 1
	for l < r {
		m := (l + r + 1) / 2
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	res[1] = l
	return
}

func main() {
	var nums []int
	var target int

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	fmt.Println(nums)
	fmt.Println(target, searchRange(nums, target))

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(nums)
	fmt.Println(target, searchRange(nums, target))
}
