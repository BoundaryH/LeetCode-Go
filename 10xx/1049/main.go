package main

import "fmt"

/* 0ms */
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, n := range stones {
		sum += n
	}
	half := sum / 2
	dp := make([]int, half+1)
	for _, i := range stones {
		for j := half; j-i >= 0; j-- {
			if dp[j-i]+i > dp[j] {
				dp[j] = dp[j-i] + i
			}
		}
	}
	return sum - dp[half]*2
}

func main() {
	var stones []int

	stones = []int{2, 7, 4, 1, 8, 1}
	fmt.Println(stones)
	fmt.Println(lastStoneWeightII(stones))

	stones = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 14, 23, 37, 61, 98}
	fmt.Println(stones)
	fmt.Println(lastStoneWeightII(stones))
}
