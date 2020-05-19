package main

import "fmt"

/* 0ms */
func spiralOrder(matrix [][]int) (res []int) {
	if len(matrix) == 0 {
		return res
	}
	row, col := len(matrix), len(matrix[0])
	res = make([]int, 0, row*col)
	layer := (row + 1) / 2
	if row > col {
		layer = (col + 1) / 2
	}

	for i := 0; i < layer; i++ {
		for x, y := i, i; y <= col-i-1; y++ {
			res = append(res, matrix[x][y])
		}
		for x, y := i+1, col-i-1; x <= row-i-1; x++ {
			res = append(res, matrix[x][y])
		}
		if row-2*i == 1 {
			break
		}
		for x, y := row-i-1, col-i-2; y >= i; y-- {
			res = append(res, matrix[x][y])
		}
		if col-2*i == 1 {
			break
		}
		for x, y := row-i-2, i; x > i; x-- {
			res = append(res, matrix[x][y])
		}
	}
	return
}

func main() {
	matrix := func(row, col int) (m [][]int) {
		m = make([][]int, row)
		for x := 0; x < row; x++ {
			for y := 0; y < col; y++ {
				m[x] = append(m[x], x*col+y+1)
			}
		}
		return m
	}

	var row, col int

	row, col = 3, 3
	for _, r := range matrix(row, col) {
		fmt.Println(r)
	}
	fmt.Println(spiralOrder(matrix(row, col)))
	fmt.Println()

	row, col = 5, 3
	for _, r := range matrix(row, col) {
		fmt.Println(r)
	}
	fmt.Println(spiralOrder(matrix(row, col)))
	fmt.Println()

	row, col = 3, 4
	for _, r := range matrix(row, col) {
		fmt.Println(r)
	}
	fmt.Println(spiralOrder(matrix(row, col)))
	fmt.Println()
}
