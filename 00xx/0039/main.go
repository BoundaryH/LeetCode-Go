package main

import "fmt"

/* 0ms */
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	arr := []int{}
	return combin(candidates, target, arr, res)
}

func combin(candidates []int, target int, arr []int, res [][]int) [][]int {
	if target == 0 {
		return append(res, append([]int{}, arr...))
	}
	for i, n := range candidates {
		if n <= target {
			res = combin(candidates[i:], target-n, append(arr, n), res)
		}
	}
	return res
}

func main() {
	var candidates []int
	var target int
	var res [][]int

	candidates = []int{2, 3, 6, 7}
	target = 7
	res = combinationSum(candidates, target)
	fmt.Println(candidates, target)
	for _, r := range res {
		fmt.Println("  ", r)
	}

	candidates = []int{2, 3, 5}
	target = 8
	res = combinationSum(candidates, target)
	fmt.Println(candidates, target)
	for _, r := range res {
		fmt.Println("  ", r)
	}

	candidates = []int{3, 2}
	target = 18
	res = combinationSum(candidates, target)
	fmt.Println(candidates, target)
	for _, r := range res {
		fmt.Println("  ", r)
	}
}
