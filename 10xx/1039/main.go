package main

import "fmt"

/* 0ms */
func minScoreTriangulation(A []int) int {
	n := len(A)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			score := 0x7fffffff
			for k := i + 1; k < j; k++ {
				s := A[i]*A[k]*A[j] + dp[i][k] + dp[k][j]
				if s < score {
					score = s
				}
			}
			dp[i][j] = score
		}
	}
	return dp[0][n-1]
}

func main() {
	var A []int

	A = []int{1, 2, 3}
	fmt.Println(A)
	fmt.Println(minScoreTriangulation(A))

	A = []int{3, 7, 4, 5}
	fmt.Println(A)
	fmt.Println(minScoreTriangulation(A))

	A = []int{1, 3, 1, 4, 1, 5}
	fmt.Println(A)
	fmt.Println(minScoreTriangulation(A))
}
