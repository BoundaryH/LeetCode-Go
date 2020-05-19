package main

import "fmt"

/* 4ms */
func regionsBySlashes(grid []string) int {
	row, col := len(grid), len(grid[0])
	const (
		U = iota
		D
		L
		R
	)
	idx := func(x, y, z int) int {
		return (x*col+y)<<2 + z
	}
	res := row * col * 4
	links := make([]int, res)
	for i := range links {
		links[i] = i
	}
	union := func(a, b int) int {
		for links[a] != a {
			a = links[a]
		}
		for links[b] != b {
			b = links[b]
		}
		if a != b {
			res--
		}
		links[b] = a
		return a
	}

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			i := idx(x, y, 0)
			u, d, l, r := i, i+1, i+2, i+3
			if x > 0 {
				u = union(idx(x-1, y, D), u)
			}
			if y > 0 {
				l = union(idx(x, y-1, R), l)
			}
			if x+1 < row {
				d = union(idx(x+1, y, U), d)
			}
			if y+1 < col {
				r = union(idx(x, y+1, L), r)
			}
			if grid[x][y] == '/' {
				union(u, l)
				union(d, r)
			} else if grid[x][y] == '\\' {
				union(u, r)
				union(d, l)
			} else {
				union(union(u, d), union(l, r))
			}
		}
	}
	return res
}

func main() {
	var grid []string

	grid = []string{" /", "/ "}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(regionsBySlashes(grid))

	grid = []string{" /", "  "}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(regionsBySlashes(grid))

	grid = []string{"\\/", "/\\"}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(regionsBySlashes(grid))

	grid = []string{"/\\", "\\/"}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(regionsBySlashes(grid))

	grid = []string{"//", "/ "}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(regionsBySlashes(grid))

}
