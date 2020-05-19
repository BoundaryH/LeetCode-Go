package main

import "fmt"

/* 0ms */
func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	res := 1
	for n := 1 % K; n != 0; {
		n = (n*10 + 1) % K
		res++
	}
	return res
}

func main() {
	var K int

	K = 1
	fmt.Println(K, smallestRepunitDivByK(K))
	K = 2
	fmt.Println(K, smallestRepunitDivByK(K))
	K = 3
	fmt.Println(K, smallestRepunitDivByK(K))
}
