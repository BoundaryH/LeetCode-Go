package main

import "fmt"

/* 8ms */
/*
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, n := range nums {
		for i := sum - n; i >= 0; i-- {
			if dp[i] {
				dp[i+n] = true
			}
		}
	}
	return dp[sum]
}
*/

type NC struct {
	num   int
	count int
}

/* 0ms */
func canPartition(nums []int) bool {
	sum := 0
	numMap := make(map[int]int)
	for _, n := range nums {
		sum += n
		numMap[n]++
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	numCount := make([]*NC, 0)
	for n, c := range numMap {
		numCount = append(numCount, &NC{n, c})
	}
	return dfs(numCount, sum)
}

func dfs(numCount []*NC, sum int) bool {
	if sum == 0 {
		return true
	}
	if len(numCount) == 0 {
		return false
	}
	for i, nc := range numCount {
		if nc.count == 0 {
			continue
		}
		nc.count--
		if sum >= nc.num && dfs(numCount[i:], sum-nc.num) {
			return true
		}
		nc.count++
	}
	return false
}

func main() {
	var nums []int

	nums = []int{1, 5, 11, 5}
	fmt.Println(nums)
	fmt.Println(canPartition(nums))

	nums = []int{1, 2, 3, 5}
	fmt.Println(nums)
	fmt.Println(canPartition(nums))

	nums = []int{2, 2, 3, 5}
	fmt.Println(nums)
	fmt.Println(canPartition(nums))
}
