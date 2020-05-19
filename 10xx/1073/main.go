package main

import "fmt"

/* 4ms */
func addNegabinary(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2)+3)
	i, j, k := len(arr1)-1, len(arr2)-1, len(res)-1
	for add := 0; i >= 0 || j >= 0 || add != 0; k-- {
		sum := -1 * add
		if i >= 0 {
			sum += arr1[i]
			i--
		}
		if j >= 0 {
			sum += arr2[j]
			j--
		}
		res[k] = sum & 1
		add = sum >> 1
	}
	for k++; k < len(res)-1 && res[k] == 0; k++ {
	}
	return res[k:]
}

func main() {
	var arr1, arr2 []int

	arr1 = []int{1, 1, 1, 1, 1}
	arr2 = []int{1, 0, 1}
	fmt.Println(arr1, arr2)
	fmt.Println(addNegabinary(arr1, arr2))

	arr1 = []int{1}
	arr2 = []int{1}
	fmt.Println(arr1, arr2)
	fmt.Println(addNegabinary(arr1, arr2))

	arr1 = []int{1}
	arr2 = []int{1, 1}
	fmt.Println(arr1, arr2)
	fmt.Println(addNegabinary(arr1, arr2))

	arr1 = []int{1}
	arr2 = []int{1, 0, 1}
	fmt.Println(arr1, arr2)
	fmt.Println(addNegabinary(arr1, arr2))
}
