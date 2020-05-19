package main

import "fmt"

/* 0ms */
func findComplement(N int) int {
	m := 1
	for m < N {
		m = m<<1 + 1
	}
	return N ^ m
}

func main() {
	var N int

	N = 5
	fmt.Println(N, findComplement(N))

	N = 1
	fmt.Println(N, findComplement(N))

	N = 7
	fmt.Println(N, findComplement(N))

	N = 10
	fmt.Println(N, findComplement(N))

	N = 0
	fmt.Println(N, findComplement(N))
}
