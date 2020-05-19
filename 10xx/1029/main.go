package main

import (
	"fmt"
	"sort"
)

/* 0ms */
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] > costs[j][1]-costs[j][0]
	})
	sum := 0
	for _, c := range costs[:len(costs)/2] {
		sum += c[0]
	}
	for _, c := range costs[len(costs)/2:] {
		sum += c[1]
	}
	return sum
}

func main() {
	var costs [][]int

	costs = [][]int{[]int{10, 20}, []int{30, 200}, []int{400, 50}, []int{30, 20}}
	fmt.Println(twoCitySchedCost(costs), costs)
}
