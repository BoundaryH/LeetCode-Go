package main

import "fmt"

/* 0ms */
func smallestSubsequence(text string) string {
	count := [26]int{}
	flag := [26]bool{}
	for _, c := range text {
		count[c-'a']++
	}
	stk := make([]rune, 0)
	for _, ch := range text {
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
	var text string

	text = "cdadabcc"
	fmt.Println(text, smallestSubsequence(text))

	text = "abcd"
	fmt.Println(text, smallestSubsequence(text))

	text = "ecbacba"
	fmt.Println(text, smallestSubsequence(text))

	text = "leetcode"
	fmt.Println(text, smallestSubsequence(text))

}
