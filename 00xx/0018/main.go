package main

import (
	"fmt"
	"sort"
)

/* 4ms */
func fourSum(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	size := len(nums)

	var a, b, c, d int
	for a = 0; a < size-3; a++ {
		n1 := nums[a]
		if a > 0 && n1 == nums[a-1] {
			continue
		}
		for b = a + 1; b < size-2; b++ {
			n2 := nums[b]
			if b > a+1 && n2 == nums[b-1] {
				continue
			}
			if n1+n2+nums[b+1]+nums[b+2] <= target &&
				n1+n2+nums[size-2]+nums[size-1] >= target {
				c = b + 1
				d = size - 1
				for c < d {
					n3, n4 := nums[c], nums[d]
					sum := n1 + n2 + n3 + n4
					if sum == target {
						res = append(res, []int{n1, n2, n3, n4})
					}
					if sum < target {
						for ; c < d && nums[c] == n3; c++ {
						}
					} else {
						for ; c < d && nums[d] == n4; d-- {
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	var nums []int
	var target int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0
	fmt.Println(nums)
	fmt.Println(target)
	for _, r := range fourSum(nums, target) {
		fmt.Println("  ", r)
	}

	nums = []int{-1, 0, 1, 2, -1, -4}
	target = -1
	fmt.Println(nums)
	fmt.Println(target)
	for _, r := range fourSum(nums, target) {
		fmt.Println("  ", r)
	}

	nums = []int{5, 5, 3, 5, 1, -5, 1, -2}
	target = 4
	fmt.Println(nums)
	fmt.Println(target)
	for _, r := range fourSum(nums, target) {
		fmt.Println("  ", r)
	}
}
