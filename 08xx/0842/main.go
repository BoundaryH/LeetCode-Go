package main

import (
	"fmt"
	"strconv"
)

/* 0ms */
func splitIntoFibonacci(S string) []int {
	const INT_MAX = 0x7fffffff
	for i := len(S) - 1; i > 0; i-- {
		if S[i] == '0' && len(S)-i > 1 {
			continue
		}
		a, err := strconv.ParseInt(S[i:], 10, 64)
		if err != nil || a > INT_MAX {
			break
		}

		for j := i - 1; j > 0; j-- {
			if S[j] == '0' && i-j > 1 {
				continue
			}
			b, err := strconv.ParseInt(S[j:i], 10, 64)
			if err != nil || b > a {
				break
			}

			if res := next(S[:j], int(b), int(a)); len(res) > 0 {
				res = append(res, int(b))
				res = append(res, int(a))
				return res
			}
		}
	}
	return []int{}
}

func next(S string, b, a int) []int {
	c := a - b
	s := strconv.Itoa(c)
	if i := len(S) - len(s); i >= 0 && S[i:] == s {
		if i == 0 {
			return []int{c}
		} else if res := next(S[:i], c, b); len(res) > 0 {
			return append(res, c)
		}
	}
	return []int{}
}

func main() {
	var S string

	S = "123456579"
	fmt.Println(S)
	fmt.Println(splitIntoFibonacci(S))

	S = "11235813"
	fmt.Println(S)
	fmt.Println(splitIntoFibonacci(S))

	S = "112358130"
	fmt.Println(S)
	fmt.Println(splitIntoFibonacci(S))

	S = "0123"
	fmt.Println(S)
	fmt.Println(splitIntoFibonacci(S))

	S = "1101111"
	fmt.Println(S)
	fmt.Println(splitIntoFibonacci(S))

	S = "0000"
	fmt.Println(S)
	fmt.Println(splitIntoFibonacci(S))
}
