package main

import "fmt"

/* 16ms */
func sortArray(nums []int) []int {
	const OFFSET = 50000
	count := make([]int, 100010)
	min, max := 50000, -50000
	for _, n := range nums {
		count[n+OFFSET]++
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	k := 0
	for i := min; i <= max; i++ {
		c := count[i+OFFSET]
		for j := 0; j < c; j++ {
			nums[k] = i
			k++
		}
	}
	return nums
}

func main() {
	var nums []int

	nums = []int{5, 2, 3, 1}
	fmt.Println(nums)
	fmt.Println(sortArray(nums))

	nums = []int{5, 1, 1, 2, 0, 0}
	fmt.Println(nums)
	fmt.Println(sortArray(nums))

}
