package main

import "fmt"

/* 0ms */
func permute(nums []int) (res [][]int) {
	var dfs func(int)
	dfs = func(p int) {
		if p == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}
		for i := p; i < len(nums); i++ {
			nums[i], nums[p] = nums[p], nums[i]
			dfs(p + 1)
			nums[i], nums[p] = nums[p], nums[i]
		}
	}
	dfs(0)
	return
}

func main() {
	var nums []int

	nums = []int{1, 2, 3}
	fmt.Println(nums)
	for _, i := range permute(nums) {
		fmt.Println("  ", i)
	}
}
