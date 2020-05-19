package main

import "fmt"

/* 8ms */
func combine(n int, k int) (res [][]int) {
	var cmb func(i int, a []int)
	cmb = func(i int, a []int) {
		if len(a) == k {
			res = append(res, append([]int{}, a...))
			return
		}
		for ; i <= n; i++ {
			cmb(i+1, append(a, i))
		}
	}
	cmb(1, make([]int, 0, k))
	return
}

func main() {
	var n, k int

	n, k = 4, 2
	fmt.Println(n, k)
	for _, r := range combine(n, k) {
		fmt.Println(r)
	}

	n, k = 5, 5
	fmt.Println(n, k)
	for _, r := range combine(n, k) {
		fmt.Println(r)
	}
}
