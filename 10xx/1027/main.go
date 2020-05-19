package main

import "fmt"

/* 12ms */
func longestArithSeqLength(A []int) int {
	n := len(A)
	dp := make([][]int, n)
	idx := make([]int, 10002)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	res := 2
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := A[i]*2 - A[j]
			if a < 0 || a >= len(idx) || idx[a] == 0 {
				dp[i][j] = 2
			} else {
				dp[i][j] = dp[idx[a]-1][i] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
		idx[A[i]] = i + 1
	}
	return res
}

func main() {
	var A []int

	A = []int{3, 6, 9, 12}
	fmt.Println(A)
	fmt.Println(longestArithSeqLength(A))

	A = []int{9, 4, 7, 2, 10}
	fmt.Println(A)
	fmt.Println(longestArithSeqLength(A))

	A = []int{20, 1, 15, 3, 10, 5, 8}
	fmt.Println(A)
	fmt.Println(longestArithSeqLength(A))

	A = []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}
	fmt.Println(A)
	fmt.Println(longestArithSeqLength(A))

}
