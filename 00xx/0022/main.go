package main

import "fmt"

/* 0ms */
func generateParenthesis(n int) (res []string) {
	var find func([]byte, int, int)
	find = func(s []byte, left, right int) {
		if left == 0 && right == 0 {
			res = append(res, string(s))
		}
		if left > 0 {
			find(append(s, '('), left-1, right)
		}
		if right > left {
			find(append(s, ')'), left, right-1)
		}
	}

	find(make([]byte, 0, n*2), n, n)
	return
}

func main() {
	var n int

	n = 2
	fmt.Println(n)
	for _, r := range generateParenthesis(n) {
		fmt.Println(r)
	}

	n = 3
	fmt.Println(n)
	for _, r := range generateParenthesis(n) {
		fmt.Println(r)
	}
}
