package main

import "fmt"

/* 0ms */
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	max := 0
	l, r := 0, 0
	for ; r < len(s); r++ {
		if count[s[r]-'a'] < k {
			if n := longestSubstring(s[l:r], k); n > max {
				max = n
			}
			l = r + 1
		}
	}
	if l < len(s) {
		if l == 0 {
			max = len(s)
		} else if n := longestSubstring(s[l:], k); n > max {
			max = n
		}
	}
	return max
}

func main() {
	var s string
	var k int

	s = "aaabb"
	k = 3
	fmt.Println(s, k)
	fmt.Println(longestSubstring(s, k))

	s = "ababbc"
	k = 2
	fmt.Println(s, k)
	fmt.Println(longestSubstring(s, k))

	s = "ababacb"
	k = 3
	fmt.Println(s, k)
	fmt.Println(longestSubstring(s, k))

	s = "bbaaacbd"
	k = 3
	fmt.Println(s, k)
	fmt.Println(longestSubstring(s, k))
}
