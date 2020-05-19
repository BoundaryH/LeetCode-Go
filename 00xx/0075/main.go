package main

import "fmt"

/* 0ms */
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	i := l
	for i <= r {
		if nums[i] == 0 {
			nums[i] = nums[l]
			nums[l] = 0
			l++
			i++
		} else if nums[i] == 2 {
			nums[i] = nums[r]
			nums[r] = 2
			r--
		} else {
			i++
		}
	}
}

func main() {
	var nums []int

	nums = []int{1, 2, 0}
	fmt.Println(nums)
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 2, 1, 1, 0}
	fmt.Println(nums)
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{0, 1, 0}
	fmt.Println(nums)
	sortColors(nums)
	fmt.Println(nums)
}
