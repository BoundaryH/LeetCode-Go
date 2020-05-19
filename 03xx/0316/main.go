package main

import "fmt"

/* 0ms */
func removeDuplicateLetters(s string) string {
	count := [26]int{}
	flag := [26]bool{}
	for _, c := range s {
		count[c-'a']++
	}
	stk := make([]rune, 0)
	for _, ch := range s {
		count[ch-'a']--
		if flag[ch-'a'] {
			continue
		}
		for len(stk) > 0 {
			last := stk[len(stk)-1]
			if last > ch && count[last-'a'] > 0 {
				flag[last-'a'] = false
				stk = stk[:len(stk)-1]
			} else {
				break
			}
		}
		stk = append(stk, ch)
		flag[ch-'a'] = true
	}
	return string(stk)
}

func main() {
	var s string

	s = "cdadabcc"
	fmt.Println(s, removeDuplicateLetters(s))

	s = "abcd"
	fmt.Println(s, removeDuplicateLetters(s))

	s = "ecbacba"
	fmt.Println(s, removeDuplicateLetters(s))

	s = "leetcode"
	fmt.Println(s, removeDuplicateLetters(s))

}
