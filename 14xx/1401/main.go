package main

import "fmt"

/* 0ms */
func checkOverlap(radius, x_center, y_center, x1, y1, x2, y2 int) bool {
	return inBox(x_center, y_center, x1-radius, y1, x2+radius, y2) ||
		inBox(x_center, y_center, x1, y1-radius, x2, y2+radius) ||
		dist(x_center, y_center, x1, y1) <= radius*radius ||
		dist(x_center, y_center, x2, y2) <= radius*radius ||
		dist(x_center, y_center, x1, y2) <= radius*radius ||
		dist(x_center, y_center, x2, y1) <= radius*radius
}

func dist(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}
func inBox(x, y, left, bottom, right, top int) bool {
	return x >= left && x <= right && y <= top && y >= bottom
}

func main() {
	var radius, x_center, y_center int
	var x1, y1, x2, y2 int

	radius, x_center, y_center = 1, 0, 0
	x1, y1, x2, y2 = 1, -1, 3, 1
	fmt.Println(radius, x_center, y_center, x1, y1, x2, y2)
	fmt.Println(checkOverlap(radius, x_center, y_center, x1, y1, x2, y2))

	radius, x_center, y_center = 1, 0, 0
	x1, y1, x2, y2 = -1, 0, 0, 1
	fmt.Println(radius, x_center, y_center, x1, y1, x2, y2)
	fmt.Println(checkOverlap(radius, x_center, y_center, x1, y1, x2, y2))

	radius, x_center, y_center = 1, 1, 1
	x1, y1, x2, y2 = -3, -3, 3, 3
	fmt.Println(radius, x_center, y_center, x1, y1, x2, y2)
	fmt.Println(checkOverlap(radius, x_center, y_center, x1, y1, x2, y2))

	radius, x_center, y_center = 1, 1, 1
	x1, y1, x2, y2 = 1, -3, 2, -1
	fmt.Println(radius, x_center, y_center, x1, y1, x2, y2)
	fmt.Println(checkOverlap(radius, x_center, y_center, x1, y1, x2, y2))

	radius, x_center, y_center = 31, 23, 46
	x1, y1, x2, y2 = 12, 38, 15, 41
	fmt.Println(radius, x_center, y_center, x1, y1, x2, y2)
	fmt.Println(checkOverlap(radius, x_center, y_center, x1, y1, x2, y2))
}
