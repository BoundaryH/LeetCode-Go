package main

import (
	"fmt"
)

/* 132ms */
func tallestBillboard(rods []int) int {
	dp := make(map[int]int)
	dp[0] = 0
	for _, n := range rods {
		tmp := make(map[int]int)
		for k, v := range dp {
			tmp[k] = max(tmp[k], v)
			tmp[k+n] = max(tmp[k+n], v)
			if k < n {
				tmp[n-k] = max(tmp[n-k], v+k)
			} else {
				tmp[k-n] = max(tmp[k-n], v+n)
			}
		}
		dp = tmp
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var rods []int

	rods = []int{1, 2, 3, 6}
	fmt.Println(rods)
	fmt.Println(tallestBillboard(rods))

	rods = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(rods)
	fmt.Println(tallestBillboard(rods))

	rods = []int{1, 2}
	fmt.Println(rods)
	fmt.Println(tallestBillboard(rods))

	rods = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 50, 50, 50, 150, 150, 150, 100, 100, 100, 123}
	fmt.Println(rods)
	fmt.Println(tallestBillboard(rods))

	rods = []int{52, 61, 69, 53, 64, 69, 63, 78, 66, 65, 68, 78, 65, 68, 81, 800, 800, 800, 800, 800}
	fmt.Println(rods)
	fmt.Println(tallestBillboard(rods))
}
