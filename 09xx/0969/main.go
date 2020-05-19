package main

import "fmt"

/* 0ms */
func pancakeSort(A []int) []int {
	res := make([]int, 0)
	for i := len(A) - 1; i >= 0; i-- {
		maxIdx := 0
		for j := 0; j <= i; j++ {
			if A[j] >= A[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx == i {
			continue
		}
		l, r := 0, maxIdx
		res = append(res, r+1)
		for l < r {
			A[l], A[r] = A[r], A[l]
			l++
			r--
		}
		l, r = 0, i
		res = append(res, r+1)
		for l < r {
			A[l], A[r] = A[r], A[l]
			l++
			r--
		}
	}
	return res
}

func main() {
	var A []int

	A = []int{3, 2, 4, 1}
	fmt.Println(A)
	fmt.Println(pancakeSort(A))

	A = []int{1, 2, 3}
	fmt.Println(A)
	fmt.Println(pancakeSort(A))
}
