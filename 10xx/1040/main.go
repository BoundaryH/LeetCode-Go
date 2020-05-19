package main

import (
	"fmt"
	"sort"
)

/* 16ms */
func numMovesStonesII(stones []int) []int {
	n := len(stones)
	sort.Ints(stones)
	res := []int{stones[n-1] - stones[0], 0}

	if a, b := stones[n-1]-stones[1], stones[n-2]-stones[0]; a > b {
		res[1] = a - n + 2
	} else {
		res[1] = b - n + 2
	}

	for l, r := 0, 1; l <= r && r < n; {
		a := stones[r] - stones[l] + 1
		if a < n {
			r++
		} else if a != n && l == 0 && r == n-1 {
			l++
		} else {
			move := l + n - r - 1
			if a > n {
				move++
			}
			if move < res[0] {
				res[0] = move
			}
			l++
		}
	}
	return res
}

func main() {
	var stones []int

	stones = []int{7, 4, 9}
	fmt.Println(stones)
	fmt.Println(numMovesStonesII(stones))

	stones = []int{6, 5, 4, 3, 10}
	fmt.Println(stones)
	fmt.Println(numMovesStonesII(stones))

	stones = []int{100, 101, 104, 102, 103}
	fmt.Println(stones)
	fmt.Println(numMovesStonesII(stones))

	stones = []int{6, 5, 4, 3, 0}
	fmt.Println(stones)
	fmt.Println(numMovesStonesII(stones))

	stones = []int{100, 101, 104, 102, 103}
	fmt.Println(stones)
	fmt.Println(numMovesStonesII(stones))

	stones = []int{7, 4, 9}
	fmt.Println(stones)
	fmt.Println(numMovesStonesII(stones))
}
