package main

import "fmt"

/* 32ms */
func shipWithinDays(weights []int, D int) int {
	l, r := 0, 0
	for _, w := range weights {
		r += w
		if l < w {
			l = w
		}
	}
	for l != r {
		m := (l + r) / 2
		sum := 0
		d := 1
		for _, w := range weights {
			sum += w
			if sum > m {
				sum = w
				d++
			}
		}
		if d > D {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	var weights []int
	var D int

	weights = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D = 5
	fmt.Println(weights)
	fmt.Println(D, shipWithinDays(weights, D))

	weights = []int{3, 2, 2, 4, 1, 4}
	D = 3
	fmt.Println(weights)
	fmt.Println(D, shipWithinDays(weights, D))

	weights = []int{1, 2, 3, 1, 1}
	D = 4
	fmt.Println(weights)
	fmt.Println(D, shipWithinDays(weights, D))
}
