package main

import "fmt"

/* 0ms */
func soupServings(N int) float64 {
	if N >= 6000 {
		return 1
	}
	N = (N + 24) / 25
	dp := make([][]float64, N+1)
	for i := range dp {
		dp[i] = make([]float64, N+1)
	}
	op := [][]int{{4, 0}, {3, 1}, {2, 2}, {1, 3}}
	dp[N][N] = 1
	for i := N; i > 0; i-- {
		for j := N; j > 0; j-- {
			add := dp[i][j] / 4
			for k := 0; k < 4; k++ {
				x, y := i-op[k][0], j-op[k][1]
				if x < 0 {
					x = 0
				}
				if y < 0 {
					y = 0
				}
				dp[x][y] += add
			}
		}
	}
	sum := dp[0][0] / 2
	for i := 1; i <= N; i++ {
		sum += dp[0][i]
	}
	return sum
}

func main() {
	var N int

	N = 50
	fmt.Println(N, soupServings(N))
}
