package main

import "fmt"

/* 28ms */
func numPairsDivisibleBy60(time []int) int {
	count := make([]int, 60)
	p := 0
	for _, n := range time {
		a := n % 60
		p += count[(60-a)%60]
		count[a]++
	}
	return p
}

func main() {
	var time []int

	time = []int{30, 20, 150, 100, 40}
	fmt.Println(numPairsDivisibleBy60(time), time)

	time = []int{60, 60, 60}
	fmt.Println(numPairsDivisibleBy60(time), time)
}
