package main

import "fmt"

/* 0ms */
func divisorGame(N int) bool {
	return N%2 == 0
}

func main() {
	var N int

	N = 2
	fmt.Println(N, divisorGame(N))

	N = 3
	fmt.Println(N, divisorGame(N))
}
