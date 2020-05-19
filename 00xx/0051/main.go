package main

import "fmt"

/* 4ms */
func solveNQueens(n int) (res [][]string) {
	col := make([]bool, n)
	diag1 := make([]bool, n*2)
	diag2 := make([]bool, n*2)

	var dfs func([]int)
	dfs = func(q []int) {
		x := len(q)
		if x == n {
			b := make([]string, n)
			for i := 0; i < n; i++ {
				row := make([]byte, n)
				for j := 0; j < n; j++ {
					if j == q[i] {
						row[j] = 'Q'
					} else {
						row[j] = '.'
					}
				}
				b[i] = string(row)
			}
			res = append(res, b)
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
	fmt.Println(n)
	for _, i := range solveNQueens(n) {
		for _, j := range i {
			fmt.Println(j)
		}
		fmt.Println()
	}
}
