package main

import "fmt"

/* 0ms */
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

func main() {
	var n int

	for n = 0; n < 10; n++ {
		fmt.Println(n, numTrees(n))
	}
}
