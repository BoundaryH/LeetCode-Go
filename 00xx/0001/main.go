package main

import "fmt"

/* 4ms */
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := hashMap[target-v]; ok {
			return []int{j, i}
		} else {
			hashMap[v] = i
		}
	}
	return []int{0, 0}
}

/* 36ms */
/*
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
*/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
