package main

import (
	"fmt"
	"sort"
)

/* 4ms */
func threeSumClosest(nums []int, target int) (res int) {
	res = nums[0] + nums[1] + nums[2]
	min := abs(res - target)

	sort.Ints(nums)
	size := len(nums)
	for i := 0; i < size; {
		l, r := i+1, size-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			diff := abs(sum - target)
			if diff < min {
				min = diff
				res = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
		for n := nums[i]; i < size && nums[i] == n; i++ {
		}
	}
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var nums []int
	var target int

	nums = []int{-1, 2, 1, -4}
	target = 1
	fmt.Println(nums)
	fmt.Println(target, threeSumClosest(nums, target))

	nums = []int{1, 1, -1, -1, 3}
	target = 3
	fmt.Println(nums)
	fmt.Println(target, threeSumClosest(nums, target))

	nums = []int{1, 2, 4, 8, 16, 32, 64, 128}
	target = 82
	fmt.Println(nums)
	fmt.Println(target, threeSumClosest(nums, target))
}
