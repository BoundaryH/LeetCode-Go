package main

import "fmt"

/* 0ms */
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	lSum := rangeSum(A, L)
	mSum := rangeSum(A, M)

	sum := 0
	lMax := leftMax(lSum)
	rMax := rightMax(mSum)
	for i := L; len(A)-i >= M; i++ {
		s := lMax[i-1] + rMax[i+M-1]
		if s > sum {
			sum = s
		}
	}

	lMax = leftMax(mSum)
	rMax = rightMax(lSum)
	for i := M; len(A)-i >= L; i++ {
		s := lMax[i-1] + rMax[i+L-1]
		if s > sum {
			sum = s
		}
	}
	return sum
}

func rangeSum(A []int, l int) []int {
	sum := make([]int, len(A))
	s := 0
	for i := 0; i < l; i++ {
		s += A[i]
	}
	sum[l-1] = s
	for i := l; i < len(A); i++ {
		s += A[i] - A[i-l]
		sum[i] = s
	}
	return sum
}

func leftMax(A []int) []int {
	lMax := make([]int, len(A))
	max := 0
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
		lMax[i] = max
	}
	return lMax
}

func rightMax(A []int) []int {
	rMax := make([]int, len(A))
	max := 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] > max {
			max = A[i]
		}
		rMax[i] = max
	}
	return rMax
}

func main() {
	var A []int
	var L, M int

	A = []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
	L, M = 1, 2
	fmt.Println(A)
	fmt.Println(L, M, maxSumTwoNoOverlap(A, L, M))

	A = []int{3, 8, 1, 3, 2, 1, 8, 9, 0}
	L, M = 3, 2
	fmt.Println(A)
	fmt.Println(L, M, maxSumTwoNoOverlap(A, L, M))

	A = []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}
	L, M = 4, 3
	fmt.Println(A)
	fmt.Println(L, M, maxSumTwoNoOverlap(A, L, M))

	A = []int{4, 0, 1}
	L, M = 2, 1
	fmt.Println(A)
	fmt.Println(L, M, maxSumTwoNoOverlap(A, L, M))
}
