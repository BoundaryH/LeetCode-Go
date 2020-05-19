package main

import "fmt"

/* 52ms */
func maxScoreSightseeingPair(A []int) int {
	sum := 0
	maxLeft := A[0]
	for i := 1; i < len(A); i++ {
		s := A[i] - i + maxLeft
		if s > sum {
			sum = s
		}
		l := A[i] + i
		if l > maxLeft {
			maxLeft = l
		}
	}
	return sum
}

func main() {
	var A []int

	A = []int{8, 1, 5, 2, 6}
	fmt.Println(maxScoreSightseeingPair(A), A)

}
