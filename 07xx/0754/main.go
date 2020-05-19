package main

import "fmt"

/* 0ms */
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	for n := 1; ; n++ {
		sub := n*(n+1)/2 - target
		if sub >= 0 && sub&1 == 0 {
			return n
		}
	}
}

func main() {
	var i int
	i = 3
	fmt.Println(i, reachNumber(i))

	i = 2
	fmt.Println(i, reachNumber(i))

	i = 5
	fmt.Println(i, reachNumber(i))
}
