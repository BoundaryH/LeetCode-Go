package main

import "fmt"

/* 0ms */
func totalNQueens(n int) (res int) {
	col := make([]bool, n)
	diag1 := make([]bool, n*2)
	diag2 := make([]bool, n*2)

	var dfs func([]int)
	dfs = func(q []int) {
		x := len(q)
		if x == n {
			res++
			return
		}
		for y := 0; y < n; y++ {
			d1 := x + y
			d2 := x - y + n
			if col[y] || diag1[d1] || diag2[d2] {
				continue
			}
			col[y] = true
			diag1[d1] = true
			diag2[d2] = true
			dfs(append(q, y))
			col[y] = false
			diag1[d1] = false
			diag2[d2] = false

		}
	}
	dfs(make([]int, 0, n))
	return
}

func main() {
	var n int

	n = 4
	fmt.Println(n, totalNQueens(n))
}
