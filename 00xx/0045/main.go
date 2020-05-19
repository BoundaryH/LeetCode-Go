package main

import "fmt"

/* 8ms */
func jump(nums []int) int {
	step := make([]int, len(nums))
	j := 1
	for i, n := range nums {
		for j < len(nums) && j <= i+n {
			step[j] = step[i] + 1
			j++
		}
	}
	return step[len(step)-1]
}

func main() {
	var nums []int

	nums = []int{2, 3, 1, 1, 4}
	fmt.Println(nums)
	fmt.Println(jump(nums))
}
