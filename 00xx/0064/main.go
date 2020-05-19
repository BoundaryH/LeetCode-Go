package main

import "fmt"

/* 8ms */
func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dp[i][j] = grid[i][j]
			if i != 0 && j != 0 {
				if dp[i][j-1] < dp[i-1][j] {
					dp[i][j] += dp[i][j-1]
				} else {
					dp[i][j] += dp[i-1][j]
				}
			} else if i != 0 {
				dp[i][j] += dp[i-1][j]
			} else if j != 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[row-1][col-1]
}

func main() {
	var grid [][]int

	grid = [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	for _, r := range grid {
		fmt.Println(r)
	}
	fmt.Println(minPathSum(grid))
}
