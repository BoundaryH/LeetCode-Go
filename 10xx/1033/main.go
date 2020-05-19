package main

import "fmt"

/* 0ms */
func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	res := []int{1, c - a - 2}
	if b-a == 1 && c-b == 1 {
		res[0] = 0
	} else if b-a > 2 && c-b > 2 {
		res[0] = 2
	}
	return res
}

func main() {
	var (
		a int
		b int
		c int
	)

	a = 1
	b = 2
	c = 5
	fmt.Println(a, b, c, numMovesStones(a, b, c))

	a = 4
	b = 3
	c = 2
	fmt.Println(a, b, c, numMovesStones(a, b, c))

	a = 3
	b = 5
	c = 1
	fmt.Println(a, b, c, numMovesStones(a, b, c))
}
