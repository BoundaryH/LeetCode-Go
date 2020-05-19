package main

import "fmt"

/* 0ms */
func minWindow(s string, t string) (res string) {
	if len(t) == 0 {
		return ""
	}
	count := make([]int, 256)
	target := make([]bool, 256)
	d := 0
	for _, c := range t {
		target[c] = true
		if count[c] == 0 {
			d++
		}
		count[c]--
	}
	res = s
	l := 0
	for r, c := range s {
		if !target[c] {
			continue
		}
		count[c]++
		if count[c] == 0 {
			d--
		}
		for l <= r && (!target[s[l]] || count[s[l]] > 0) {
			count[s[l]]--
			l++
		}
		if d == 0 && r-l+1 < len(res) {
			res = s[l : r+1]
		}
	}
	if len(res) == len(s) && d != 0 {
		return ""
	}
	return
}

func main() {
	var s, t string

	s = "ADOBECODEBANC"
	t = "ABC"
	fmt.Println(s, t)
	fmt.Println(minWindow(s, t))

	s = "ab"
	t = "b"
	fmt.Println(s, t)
	fmt.Println(minWindow(s, t))

	s = "a"
	t = "aa"
	fmt.Println(s, t)
	fmt.Println(minWindow(s, t))
}
