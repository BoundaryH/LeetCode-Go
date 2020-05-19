package main

import "fmt"

/* 0ms */
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1
	j := len(nums) - 2
	for j >= 0 && nums[j] >= nums[i] {
		i--
		j--
	}

	if j >= 0 {
		for i < len(nums) && nums[i] > nums[j] {
			i++
		}
		i--
		nums[i], nums[j] = nums[j], nums[i]
	}

	j++
	for k := len(nums) - 1; j < k; k-- {
		nums[j], nums[k] = nums[k], nums[j]
		j++
	}
}

func main() {
	var nums []int

	nums = []int{1, 2, 3}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{5, 1, 1}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
}
