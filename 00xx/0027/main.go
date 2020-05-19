package main

import "fmt"

/* 0ms */
func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	var nums []int
	var val int
	var n int

	val = 3
	nums = []int{3, 2, 2, 3}
	fmt.Print(nums, " val=", val, " -> ")
	n = removeElement(nums, val)
	fmt.Println(nums[:n])

	val = 2
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Print(nums, " val=", val, " -> ")
	n = removeElement(nums, val)
	fmt.Println(nums[:n])
}
