package main

import "fmt"

/* 0ms */
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j != 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	var m, n int

	m, n = 3, 2
	fmt.Println(m, n, uniquePaths(m, n))

	m, n = 7, 3
	fmt.Println(m, n, uniquePaths(m, n))
}
