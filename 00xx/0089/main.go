package main

import (
	"fmt"
	"strconv"
)

/* 0ms */
func grayCode(n int) (res []int) {
	res = make([]int, 1, 1<<n)
	for i := 0; i < n; i++ {
		mask := (1 << i)
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]|mask)
		}
	}
	return
}

func main() {
	var n int

	n = 0
	fmt.Println(n)
	for _, i := range grayCode(n) {
		fmt.Printf("%0"+strconv.Itoa(n)+"b\n", i)
	}
	fmt.Println()

	n = 1
	fmt.Println(n)
	for _, i := range grayCode(n) {
		fmt.Printf("%0"+strconv.Itoa(n)+"b\n", i)
	}
	fmt.Println()

	n = 2
	fmt.Println(n)
	for _, i := range grayCode(n) {
		fmt.Printf("%0"+strconv.Itoa(n)+"b\n", i)
	}
	fmt.Println()

	n = 3
	fmt.Println(n)
	for _, i := range grayCode(n) {
		fmt.Printf("%0"+strconv.Itoa(n)+"b\n", i)
	}
	fmt.Println()
}
