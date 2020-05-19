package main

import "fmt"

/* 12ms */
func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	row, col := len(grid), len(grid[0])
	flag := make([][]byte, row)
	for i := range flag {
		flag[i] = make([]byte, col)
	}
	c := grid[r0][c0]
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if flag[x][y] != 0 {
			return
		}
		flag[x][y] = 1
		count := 0
		if x != 0 && grid[x-1][y] == c {
			dfs(x-1, y)
			count++
		}
		if x+1 < row && grid[x+1][y] == c {
			dfs(x+1, y)
			count++
		}
		if y != 0 && grid[x][y-1] == c {
			dfs(x, y-1)
			count++
		}
		if y+1 < col && grid[x][y+1] == c {
			dfs(x, y+1)
			count++
		}
		if count != 4 {
			flag[x][y] = 2
		}
	}
	dfs(r0, c0)

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if flag[x][y] == 2 {
				grid[x][y] = color
			}
		}
	}
	return grid
}

func main() {
	var grid [][]int
	var r0, c0, color int

	grid = [][]int{{1, 1}, {1, 2}}
	r0, c0, color = 0, 0, 3
	fmt.Println(grid)
	fmt.Println(r0, c0, color)
	fmt.Println(colorBorder(grid, r0, c0, color), "\n")

	grid = [][]int{{1, 2, 2}, {2, 3, 2}}
	r0, c0, color = 0, 1, 3
	fmt.Println(grid)
	fmt.Println(r0, c0, color)
	fmt.Println(colorBorder(grid, r0, c0, color), "\n")

	grid = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	r0, c0, color = 1, 1, 2
	fmt.Println(grid)
	fmt.Println(r0, c0, color)
	fmt.Println(colorBorder(grid, r0, c0, color), "\n")
}
