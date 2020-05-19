package main

import "fmt"

/* 0ms */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] != 0 {
		return 0
	}
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	dp[0][0] = 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] != 0 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[row-1][col-1]
}

func main() {
	var obstacleGrid [][]int

	obstacleGrid = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	for _, i := range obstacleGrid {
		fmt.Println(i)
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
