package main

import "fmt"

/* 0ms */
func lengthOfLongestSubstring(s string) int {
	res := 0
	idx := make([]int, 256)
	for i := range idx {
		idx[i] = -1
	}
	j := -1
	for i, c := range s {
		if idx[c] > j {
			j = idx[c]
		}
		if r := i - j; r > res {
			res = r
		}
		idx[c] = i
	}
	return res
}

func main() {
	var s string

	s = "abcdefg"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))

	s = "bbbbb"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}
