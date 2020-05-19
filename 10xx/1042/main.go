package main

import "fmt"

/* 56ms */
func gardenNoAdj(N int, paths [][]int) []int {
	pathMap := make([][]int, N)
	res := make([]int, N)
	for _, p := range paths {
		i := p[0] - 1
		j := p[1] - 1
		if i < j {
			pathMap[j] = append(pathMap[j], i)
		} else {
			pathMap[i] = append(pathMap[i], j)
		}
	}
	for i := 0; i < N; i++ {
		flag := [5]bool{}
		for _, j := range pathMap[i] {
			if res[j] != 0 {
				flag[res[j]] = true
			}
		}
		for c := 1; c < 5; c++ {
			if !flag[c] {
				res[i] = c
				break
			}
		}
	}
	return res
}

func main() {
	var N int
	var paths [][]int

	N = 3
	paths = [][]int{{1, 2}, {2, 3}, {3, 1}}
	fmt.Println(N, paths)
	fmt.Println(gardenNoAdj(N, paths))

	N = 4
	paths = [][]int{{1, 2}, {3, 4}}
	fmt.Println(N, paths)
	fmt.Println(gardenNoAdj(N, paths))

	N = 4
	paths = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}
	fmt.Println(N, paths)
	fmt.Println(gardenNoAdj(N, paths))

	N = 1
	paths = [][]int{}
	fmt.Println(N, paths)
	fmt.Println(gardenNoAdj(N, paths))
}
