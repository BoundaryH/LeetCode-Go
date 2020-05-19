package main

import "fmt"

/* 4ms */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	l, r := 0, row*col-1
	for l < r {
		m := (l + r) / 2
		n := matrix[m/col][m%col]
		if n > target {
			r = m
		} else if n < target {
			l = m + 1
		} else {
			return true
		}
	}
	return matrix[l/col][l%col] == target
}

func main() {
	var matrix [][]int
	var target int

	matrix = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	target = 3
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println(target)
	fmt.Println(searchMatrix(matrix, target))

	matrix = [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	target = 13
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println(target)
	fmt.Println(searchMatrix(matrix, target))
}
