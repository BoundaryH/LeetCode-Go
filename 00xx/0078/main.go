package main

import "fmt"

/* 0ms */
func subsets(nums []int) (res [][]int) {
	for i := (1 << len(nums)) - 1; i >= 0; i-- {
		s := make([]int, 0, len(nums))
		m := 1 << (len(nums) - 1)
		for _, n := range nums {
			if i&m != 0 {
				s = append(s, n)
			}
			m >>= 1
		}
		res = append(res, s)
	}
	return res
}

func main() {
	var nums []int

	nums = []int{1, 2, 3}
	fmt.Println(nums)
	for _, r := range subsets(nums) {
		fmt.Println(r)
	}
	fmt.Println()
}
