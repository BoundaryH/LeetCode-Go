package main

import "fmt"

func canJump(nums []int) bool {
	last := 0
	for i, n := range nums {
		if i > last {
			return false
		}
		if j := i + n; j > last {
			last = j
		}
	}
	return true
}

func main() {
	var nums []int

	nums = []int{2, 3, 1, 1, 4}
	fmt.Println(nums)
	fmt.Println(canJump(nums))

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(nums)
	fmt.Println(canJump(nums))
}
