package main

import "fmt"

/* 8ms */
func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	firstCol := false
	for i := 0; i < row; i++ {
		firstCol = (firstCol || matrix[i][0] == 0)
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if firstCol {
			matrix[i][0] = 0
		}
	}
}

func main() {
	var matrix [][]int

	matrix = [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println()
	setZeroes(matrix)
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println()

	matrix = [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println()
	setZeroes(matrix)
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println()
}
