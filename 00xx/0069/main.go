package main

import "fmt"

/* 0ms */
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 1, x
	for l < r {
		m := l + (r-l+1)/2
		if x/m < m {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func main() {
	var x int

	x = 4
	fmt.Println(x, mySqrt(x))

	x = 8
	fmt.Println(x, mySqrt(x))

	x = 2147395599
	fmt.Println(x, mySqrt(x))

}
