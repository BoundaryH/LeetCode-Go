package main

import "fmt"

/* 8ms */
func prefixesDivBy5(A []int) []bool {
	B := make([]bool, len(A))
	mod := 0
	for i, v := range A {
		mod = (mod*2 + v) % 5
		B[i] = (mod == 0)
	}
	return B
}

func main() {
	var A []int

	A = []int{0, 1, 1}
	fmt.Println(A)
	fmt.Println(prefixesDivBy5(A))

	A = []int{1, 1, 1}
	fmt.Println(A)
	fmt.Println(prefixesDivBy5(A))

	A = []int{0, 1, 1, 1, 1, 1}
	fmt.Println(A)
	fmt.Println(prefixesDivBy5(A))

	A = []int{1, 1, 1, 0, 1}
	fmt.Println(A)
	fmt.Println(prefixesDivBy5(A))
}
