package main

import "fmt"

/* 32ms */
func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0
	for i, v := range customers {
		if i < X {
			sum += v
		} else {
			sum += v * (1 - grumpy[i])
		}
	}
	maxSum := sum
	for i := X; i < len(customers); i++ {
		j := i - X
		sum -= customers[j] * grumpy[j]
		sum += customers[i] * grumpy[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	var customers, grumpy []int
	var X int

	customers = []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy = []int{0, 1, 0, 1, 0, 1, 0, 1}
	X = 3
	fmt.Println(customers)
	fmt.Println(grumpy)
	fmt.Println(X, maxSatisfied(customers, grumpy, X))

}
