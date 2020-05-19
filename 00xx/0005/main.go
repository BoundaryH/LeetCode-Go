package main

import "fmt"

/* 0ms */
func longestPalindrome(s string) string {
	var res string
	s2 := make([]rune, 0, len(s)*2+1)
	s2 = append(s2, '_')
	for _, ch := range s {
		s2 = append(s2, ch)
		s2 = append(s2, '_')
	}
	a := make([]int, len(s2))
	p, b := 0, 0
	for i := range s2 {
		a[i] = 1
		if i < b {
			a[i] = min(b-i, a[p*2-i])
		}
		for l, r := i-a[i], i+a[i]; l >= 0 && r < len(s2) && s2[l] == s2[r]; {
			a[i]++
			l--
			r++
		}
		if i+a[i] > b {
			b = i + a[i]
			p = i
		}
		if a[i]-1 > len(res) {
			x := (i - a[i] + 1) / 2
			res = s[x : x+a[i]-1]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/* 4ms */
/*
func longestPalindrome(s string) string {
	var res string
	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if p := s[l+1 : r]; len(p) > len(res) {
			res = p
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if p := s[l+1 : r]; len(p) > len(res) {
			res = p
		}
	}
	return res
}
*/

func main() {
	var s string

	s = "babad"
	fmt.Println(s)
	fmt.Println(longestPalindrome(s))

	s = "cbbd"
	fmt.Println(s)
	fmt.Println(longestPalindrome(s))
}
