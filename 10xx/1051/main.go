package main

import (
	"fmt"
	"sort"
)

/* 0ms */
func heightChecker(heights []int) int {
	tmp := make([]int, len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)
	count := 0
	for i, v := range tmp {
		if v != heights[i] {
			count++
		}
	}
	return count
}

func main() {
	var heights []int

	heights = []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(heights), heights)
}
