package main

import "fmt"

/* 8ms */
func maxArea(height []int) (res int) {
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			if h := (r - l) * height[l]; h > res {
				res = h
			}
			l++
		} else {
			if h := (r - l) * height[r]; h > res {
				res = h
			}
			r--
		}
	}
	return
}

func main() {
	var height []int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(height)
	fmt.Println(maxArea(height))
}
