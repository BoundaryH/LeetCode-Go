package main

import "fmt"

/* 4ms */
func insert(intervals [][]int, newInterval []int) (res [][]int) {
	res = make([][]int, 0, len(intervals))
	for i, t := range intervals {
		if t[1] < newInterval[0] {
			res = append(res, intervals[i])
		} else if t[0] > newInterval[1] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			newInterval = nil
			break
		} else {
			if t[0] < newInterval[0] {
				newInterval[0] = t[0]
			}
			if t[1] > newInterval[1] {
				newInterval[1] = t[1]
			}
		}
	}
	if newInterval != nil {
		res = append(res, newInterval)
	}
	return
}

func main() {
	var intervals [][]int
	var newInterval []int

	intervals = [][]int{
		[]int{1, 3},
		[]int{6, 9},
	}
	newInterval = []int{2, 5}
	fmt.Println(intervals)
	fmt.Println(newInterval)
	fmt.Println(insert(intervals, newInterval))
	fmt.Println()

	intervals = [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}
	newInterval = []int{4, 8}
	fmt.Println(intervals)
	fmt.Println(newInterval)
	fmt.Println(insert(intervals, newInterval))
	fmt.Println()
}
