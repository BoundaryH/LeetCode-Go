package main

import "fmt"

/* 0ms */
func isBoomerang(points [][]int) bool {
	a := points[0]
	b := points[1]
	c := points[2]
	return (a[0]-c[0])*(b[1]-c[1]) != (b[0]-c[0])*(a[1]-c[1])
}

func main() {
	var points [][]int

	points = [][]int{[]int{1, 1}, []int{2, 3}, []int{3, 2}}
	fmt.Println(points)
	fmt.Println(isBoomerang(points))

	points = [][]int{[]int{1, 1}, []int{2, 2}, []int{3, 3}}
	fmt.Println(points)
	fmt.Println(isBoomerang(points))

	points = [][]int{[]int{1, 1}, []int{1, 1}, []int{1, 1}}
	fmt.Println(points)
	fmt.Println(isBoomerang(points))
}
