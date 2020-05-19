package main

import "fmt"

/* 4ms */
func maxSumAfterPartitioning(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	dp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		maxV := A[i]
		for j := 0; j < K && i-j >= 0; j++ {
			if maxV < A[i-j] {
				maxV = A[i-j]
			}
			sum := maxV * (j + 1)
			if i != j {
				sum += dp[i-j-1]
			}
			if dp[i] < sum {
				dp[i] = sum
			}
		}
	}
	return dp[len(A)-1]
}

func main() {
	var A []int
	var K int

	A = []int{1, 15, 7, 9, 2, 5, 10}
	K = 3
	fmt.Println(K, A)
	fmt.Println(maxSumAfterPartitioning(A, K))
}
