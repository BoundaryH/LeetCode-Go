package main

import "fmt"

/* 0ms */
func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	max := 0
	l, r := 0, 0
	for ; r < len(s); r++ {
		rc := s[r] - 'A'
		count[rc]++
		if count[rc] > max {
			max = count[rc]
		}
		if r-l+1-max > k {
			count[s[l]-'A']--
			l++
		}
	}
	return r - l
}

func main() {
	var s string
	var k int

	s = "ABAB"
	k = 2
	fmt.Println(s)
	fmt.Println(k, characterReplacement(s, k))

	s = "AABABBA"
	k = 1
	fmt.Println(s)
	fmt.Println(k, characterReplacement(s, k))

	s = "ABBB"
	k = 2
	fmt.Println(s)
	fmt.Println(k, characterReplacement(s, k))

	s = "BAAAB"
	k = 2
	fmt.Println(s)
	fmt.Println(k, characterReplacement(s, k))

	s = "CCCCCCBAAAB"
	k = 2
	fmt.Println(s)
	fmt.Println(k, characterReplacement(s, k))
}
