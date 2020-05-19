package main

import "fmt"

/* 0ms */
func generateMatrix(n int) (res [][]int) {
	res = make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i, j := 0, 0; j < n*n; i++ {
		for y := i; y < n-i; y++ {
			j++
			res[i][y] = j
		}
		for x := i + 1; x < n-i; x++ {
			j++
			res[x][n-i-1] = j
		}
		for y := n - i - 2; y >= i; y-- {
			j++
			res[n-i-1][y] = j
		}
		for x := n - i - 2; x > i; x-- {
			j++
			res[x][i] = j
		}
	}
	return
}

func main() {
	var n int

	n = 3
	fmt.Println(n)
	for _, r := range generateMatrix(n) {
		fmt.Println(r)
	}

}
