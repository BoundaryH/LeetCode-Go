package main

import (
	"fmt"
	"sort"
)

/* 0ms */
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	arr := []int{}
	sort.Ints(candidates)
	return combin(candidates, target, arr, res)
}

func combin(candidates []int, target int, arr []int, res [][]int) [][]int {
	if target == 0 {
		return append(res, append([]int{}, arr...))
	}
	for i := 0; i < len(candidates); i++ {
		n := candidates[i]
		if n <= target {
			res = combin(candidates[i+1:], target-n, append(arr, n), res)
		}
		for i+1 < len(candidates) && candidates[i+1] == n {
			i++
		}
	}
	return res
}

func main() {
	var candidates []int
	var target int
	var res [][]int

	candidates = []int{10, 1, 2, 7, 6, 1, 5}
	target = 8
	res = combinationSum2(candidates, target)
	fmt.Println(candidates, target)
	for _, r := range res {
		fmt.Println("  ", r)
	}

	candidates = []int{2, 5, 2, 1, 2}
	target = 5
	res = combinationSum2(candidates, target)
	fmt.Println(candidates, target)
	for _, r := range res {
		fmt.Println("  ", r)
	}

}
