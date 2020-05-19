package main

import "fmt"

/* 0ms */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	x := nums[0]
	i, c := 0, 0
	for _, n := range nums {
		if n != x {
			x = n
			c = 1
			nums[i] = n
			i++
		} else if c < 2 {
			c++
			nums[i] = n
			i++
		}
	}
	return i
}

func main() {
	var nums []int

	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(nums)
	fmt.Println(nums[:removeDuplicates(nums)])

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(nums)
	fmt.Println(nums[:removeDuplicates(nums)])
}
