package main

import "fmt"

/* 0ms */
func plusOne(digits []int) []int {
	a := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += a
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			a = 0
			break
		}
	}
	if a > 0 {
		digits = append([]int{a}, digits...)
	}
	return digits
}

func main() {
	var digits []int

	digits = []int{1, 2, 3}
	fmt.Print(digits)
	fmt.Println(plusOne(digits))
	digits = []int{4, 3, 2, 1}
	fmt.Print(digits)
	fmt.Println(plusOne(digits))
	digits = []int{9, 9, 9, 9}
	fmt.Print(digits)
	fmt.Println(plusOne(digits))

}
