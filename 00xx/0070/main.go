package main

import "fmt"

/* 0ms */
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	for n := 0; n < 20; n++ {
		fmt.Println(n, climbStairs(n))
	}
}
