package main

import "fmt"

/* 0ms */
func isValid(s string) bool {
	stk := make([]rune, 0)
	for _, c := range s {
		if c == ')' {
			l := len(stk) - 1
			if l < 0 || stk[l] != '(' {
				return false
			}
			stk = stk[:l]

		} else if c == ']' {
			l := len(stk) - 1
			if l < 0 || stk[l] != '[' {
				return false
			}
			stk = stk[:l]

		} else if c == '}' {
			l := len(stk) - 1
			if l < 0 || stk[l] != '{' {
				return false
			}
			stk = stk[:l]
		} else {
			stk = append(stk, c)
		}
	}
	return len(stk) == 0
}

func main() {
	var s string

	s = "()"
	fmt.Println(isValid(s), s)
	s = "()[]{}"
	fmt.Println(isValid(s), s)
	s = "(]"
	fmt.Println(isValid(s), s)
	s = "([)]"
	fmt.Println(isValid(s), s)
	s = "{[]}"
	fmt.Println(isValid(s), s)
}
