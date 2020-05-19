package main

import "fmt"

/* 0ms */
func myPow(x float64, n int) (res float64) {
	if n == 0 {
		return 1
	}
	y := myPow(x, n/2)
	res = y * y
	if n&1 == 1 {
		if n > 0 {
			res *= x
		} else {
			res /= x
		}
	}
	return
}

func main() {
	var x float64
	var n int

	x = 2
	n = 10
	fmt.Println(x, n)
	fmt.Println(myPow(x, n))

	x = 2.1
	n = 3
	fmt.Println(x, n)
	fmt.Println(myPow(x, n))

	x = 2
	n = -2
	fmt.Println(x, n)
	fmt.Println(myPow(x, n))
}
