package main

import "fmt"

/* 0ms */
func firstMissingPositive(nums []int) int {
	mask := make([]bool, len(nums)+1)
	for _, n := range nums {
		if n > 0 && n < len(mask) {
			mask[n] = true
		}
	}
	mask[0] = true
	for i, m := range mask {
		if !m {
			return i
		}
	}
	return len(mask)
}

func main() {
	var nums []int

	nums = []int{3, 4, -1, 1}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{7, 8, 9, 11, 12}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{1, 2, 0}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{1}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{1, 1}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{0, 1, 2}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))

	nums = []int{2, 2}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))
}
