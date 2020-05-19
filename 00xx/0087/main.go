package main

import "fmt"

/* 0ms */
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	count := make([]int, 256)
	for _, c := range s1 {
		count[c]++
	}
	for _, c := range s2 {
		count[c]--
		if count[c] < 0 {
			return false
		}
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		j := len(s1) - i
		if isScramble(s1[:i], s2[j:]) && isScramble(s1[i:], s2[:j]) {
			return true
		}
	}
	return false
}

func main() {
	var s1, s2 string

	s1, s2 = "great", "rgeat"
	fmt.Println(s1, s2)
	fmt.Println(isScramble(s1, s2))

	s1, s2 = "abcde", "caebd"
	fmt.Println(s1, s2)
	fmt.Println(isScramble(s1, s2))
}
