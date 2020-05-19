package main

import "fmt"

/* 0ms */
func maxUncrossedLines(A []int, B []int) int {
	dp := make([]int, len(B))
	tmp := make([]int, len(B))
	for i := 0; i < len(A); i++ {
		tmp[0] = dp[0]
		if A[i] == B[0] {
			tmp[0] = 1
		}
		for j := 1; j < len(B); j++ {
			tmp[j] = max(tmp[j-1], dp[j])
			if A[i] == B[j] {
				tmp[j] = max(tmp[j], dp[j-1]+1)
			}
		}
		dp, tmp = tmp, dp
	}
	return dp[len(B)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var A, B []int

	A = []int{1, 4, 2}
	B = []int{1, 2, 4}
	fmt.Println(A, B)
	fmt.Println(maxUncrossedLines(A, B))

	A = []int{2, 5, 1, 2, 5}
	B = []int{10, 5, 2, 1, 5, 2}
	fmt.Println(A, B)
	fmt.Println(maxUncrossedLines(A, B))

	A = []int{1, 3, 7, 1, 7, 5}
	B = []int{1, 9, 2, 5, 1}
	fmt.Println(A, B)
	fmt.Println(maxUncrossedLines(A, B))

	A = []int{1, 2, 2, 2, 1, 2}
	B = []int{1}
	fmt.Println(A, B)
	fmt.Println(maxUncrossedLines(A, B))
}
