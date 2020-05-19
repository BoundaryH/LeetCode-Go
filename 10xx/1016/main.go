package main

import (
	"fmt"
	"strings"
)

/* 0ms */
func queryString(S string, N int) bool {
	b := 1
	for n := N >> 1; n != 0; n >>= 1 {
		b <<= 1
	}
	for i := 0; i < b; i++ {
		n := N - i
		if !strings.Contains(S, fmt.Sprintf("%b", n)) {
			return false
		}
	}
	return true
}

func main() {
	var S string
	var N int

	S = "0110"
	N = 3
	fmt.Println(queryString(S, N), S, N)

	S = "0110"
	N = 4
	fmt.Println(queryString(S, N), S, N)
}
