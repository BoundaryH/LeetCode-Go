package main

import "fmt"

/* 0ms */
func rotate(matrix [][]int) {
	n := len(matrix)
	for x := 0; x < (n+1)/2; x++ {
		for y := 0; y < n/2; y++ {
			x2 := n - x - 1
			y2 := n - y - 1
			matrix[x][y], matrix[y2][x], matrix[x2][y2], matrix[y][x2] =
				matrix[y2][x], matrix[x2][y2], matrix[y][x2], matrix[x][y]

		}
	}
}

func main() {
	var matrix [][]int
	var n int

	n = 3
	matrix = nil
	for x := 0; x < n; x++ {
		matrix = append(matrix, nil)
		for y := 0; y < n; y++ {
			matrix[x] = append(matrix[x], x*n+y+1)
		}
	}

	for _, r := range matrix {
		fmt.Println(r)
	}
	rotate(matrix)
	fmt.Println()
	for _, r := range matrix {
		fmt.Println(r)
	}

}
