package main

import (
	"fmt"
	"sort"
)

/* 8ms */
func merge(intervals [][]int) (res [][]int) {
	if len(intervals) == 0 {
		return
	}
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	res = make([][]int, 0, len(intervals))

	j := intervals[0]
	for _, i := range intervals {
		if j[1] < i[0] {
			res = append(res, j)
			j = i
		} else if i[1] > j[1] {
			j[1] = i[1]
		}
	}
	res = append(res, j)
	return
}

func main() {
	var intervals [][]int

	intervals = [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	fmt.Println(intervals)
	fmt.Println(merge(intervals))

	intervals = [][]int{
		[]int{1, 4},
		[]int{4, 5},
	}
	fmt.Println(intervals)
	fmt.Println(merge(intervals))
}
