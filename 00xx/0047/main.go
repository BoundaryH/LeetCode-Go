package main

import (
	"fmt"
)

/* 0ms */
func permuteUnique(nums []int) (res [][]int) {
	var dfs func(p int)
	dfs = func(p int) {
		if p == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}
		m := make(map[int]bool)
		for i := p; i < len(nums); i++ {
			if !m[nums[i]] {
				m[nums[i]] = true
				nums[i], nums[p] = nums[p], nums[i]
				dfs(p + 1)
				nums[i], nums[p] = nums[p], nums[i]
			}
		}
	}
	dfs(0)
	return
}

func main() {
	var nums []int

	nums = []int{1, 1, 2}
	fmt.Println(nums)
	for _, i := range permuteUnique(nums) {
		fmt.Println("  ", i)
	}
}
