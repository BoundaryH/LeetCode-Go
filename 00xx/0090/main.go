package main

import (
	"fmt"
	"sort"
)

/* 0ms */
func subsetsWithDup(nums []int) (res [][]int) {
	sort.Ints(nums)

	var subset func(n, s []int)
	subset = func(n, s []int) {
		res = append(res, append([]int{}, s...))
		for i := range n {
			if i == 0 || n[i] != n[i-1] {
				subset(n[i+1:], append(s, n[i]))
			}
		}
	}
	subset(nums, make([]int, 0, len(nums)))
	return
}

func main() {
	var nums []int

	nums = []int{1, 2, 2}
	fmt.Println(nums)
	for _, i := range subsetsWithDup(nums) {
		fmt.Println(i)
	}
	fmt.Println()
}
