package main

import "fmt"

/* 4ms */
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	var nums []int
	var n int

	nums = []int{1, 1, 2}
	n = removeDuplicates(nums)
	fmt.Println(nums[:n])

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n = removeDuplicates(nums)
	fmt.Println(nums[:n])
}
