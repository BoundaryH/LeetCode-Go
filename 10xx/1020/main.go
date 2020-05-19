package main

import "fmt"

/* 52ms */
func numEnclaves(A [][]int) int {
	var x, y int
	row, col := len(A), len(A[0])

	y = col - 1
	for x := 0; x < row; x++ {
		dfs(A, x, 0)
		dfs(A, x, y)
	}
	x = row - 1
	for y := 0; y < col; y++ {
		dfs(A, 0, y)
		dfs(A, x, y)
	}
	count := 0
	for x = 0; x < row; x++ {
		for y = 0; y < col; y++ {
			if A[x][y] == 1 {
				count++
			}
		}
	}
	return count
}

func dfs(A [][]int, x, y int) {
	if x < 0 || x >= len(A) || y < 0 || y >= len(A[0]) || A[x][y] != 1 {
		return
	}
	A[x][y] = 2
	dfs(A, x+1, y)
	dfs(A, x-1, y)
	dfs(A, x, y+1)
	dfs(A, x, y-1)
}

func main() {
	var A [][]int

	A = [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	for _, r := range A {
		fmt.Println(r)
	}
	fmt.Println(numEnclaves(A))

	A = [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	for _, r := range A {
		fmt.Println(r)
	}
	fmt.Println(numEnclaves(A))

	A = [][]int{{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0}, {1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0}, {1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0}, {0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0}, {1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1}}
	for _, r := range A {
		fmt.Println(r)
	}
	fmt.Println(numEnclaves(A))
}
