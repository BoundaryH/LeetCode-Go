package main

import "fmt"

/* 32ms */
func prevPermOpt1(A []int) []int {
	i := len(A) - 2
	for ; i >= 0 && A[i] <= A[i+1]; i-- {
	}
	if i < 0 {
		return A
	}
	j := len(A) - 1
	for ; A[j] >= A[i] || A[j-1] == A[j]; j-- {
	}
	A[i], A[j] = A[j], A[i]
	return A
}

func main() {
	var A []int

	A = []int{3, 2, 1}
	fmt.Println(A)
	fmt.Println(prevPermOpt1(A), "\n")

	A = []int{1, 1, 5}
	fmt.Println(A)
	fmt.Println(prevPermOpt1(A), "\n")

	A = []int{1, 9, 4, 6, 7}
	fmt.Println(A)
	fmt.Println(prevPermOpt1(A), "\n")

	A = []int{3, 1, 1, 3}
	fmt.Println(A)
	fmt.Println(prevPermOpt1(A), "\n")
}
