package main

import (
	"fmt"
	"sort"
)

/* 0ms */
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	dp := make([]int, k)
	return dfs(nums, dp, sum/k)
}

func dfs(nums []int, dp []int, sum int) bool {
	if len(nums) == 0 {
		return true
	}
	j := len(nums) - 1
	n := nums[j]
	for i, v := range dp {
		if v+n == sum || sum-v-n >= nums[0] {
			dp[i] += n
			if dfs(nums[:j], dp, sum) {
				return true
			}
			dp[i] -= n
		}
	}
	return false
}

func main() {
	var nums []int
	var k int

	nums = []int{4, 3, 2, 3, 5, 2, 1}
	k = 4
	fmt.Println(nums)
	fmt.Println(k, canPartitionKSubsets(nums, k))
}
