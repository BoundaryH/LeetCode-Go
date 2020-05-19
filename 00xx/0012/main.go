package main

import (
	"fmt"
)

/* 4ms */
func intToRoman(num int) (s string) {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i, v := range value {
		for num >= v {
			s += symbol[i]
			num -= v
		}
	}
	return
}

func main() {
	var num int

	num = 1994
	fmt.Println(num)
	fmt.Println(intToRoman(num))
}
