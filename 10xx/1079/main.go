package main

import "fmt"

/* 0ms */
func numTilePossibilities(tiles string) int {
	count := make([]int, 26)
	for _, c := range tiles {
		count[c-'A']++
	}
	return dfs(count, -1) - 1
}

func dfs(count []int, last int) int {
	sum := 1
	for i, j := range count {
		if j == 0 || i == last {
			continue
		}
		for k := 1; k <= j; k++ {
			count[i]--
			sum += dfs(count, i)
		}
		count[i] = j
	}
	return sum
}

func main() {
	var tiles string

	tiles = "AAB"
	fmt.Println(numTilePossibilities(tiles), tiles)

	tiles = "AAABBC"
	fmt.Println(numTilePossibilities(tiles), tiles)
}
