package main

import (
	"fmt"
)

/* 16ms */
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	maxDict := max(r0, R-r0) + max(c0, C-c0)
	l := make([][][]int, maxDict+1)
	for x := 0; x < R; x++ {
		for y := 0; y < C; y++ {
			d := abs(x-r0) + abs(y-c0)
			l[d] = append(l[d], []int{x, y})
		}
	}
	res := make([][]int, 0, R*C)
	for _, i := range l {
		res = append(res, i...)
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* 24ms */
/*
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	cells := make([][]int, 0)
	for x := 0; x < R; x++ {
		for y := 0; y < C; y++ {
			cells = append(cells, []int{x, y})
		}
	}
	sort.Slice(cells, func(i, j int) bool {
		return abs(cells[i][0]-r0)+abs(cells[i][1]-c0) < abs(cells[j][0]-r0)+abs(cells[j][1]-c0)
	})
	return cells
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
*/

func main() {
	var R int
	var C int
	var r0 int
	var c0 int

	R, C, r0, c0 = 1, 2, 0, 0
	fmt.Printf("R=%d C=%d r0=%d c0=%d\n", R, C, r0, c0)
	fmt.Println(allCellsDistOrder(R, C, r0, c0))

	R, C, r0, c0 = 2, 2, 0, 1
	fmt.Printf("R=%d C=%d r0=%d c0=%d\n", R, C, r0, c0)
	fmt.Println(allCellsDistOrder(R, C, r0, c0))

	R, C, r0, c0 = 2, 3, 1, 2
	fmt.Printf("R=%d C=%d r0=%d c0=%d\n", R, C, r0, c0)
	fmt.Println(allCellsDistOrder(R, C, r0, c0))
}
