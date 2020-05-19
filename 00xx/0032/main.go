package main

import "fmt"

/* 0ms */
func longestValidParentheses(s string) (res int) {
	stk := make([]int, 0, len(s))
	start := -1
	for i, c := range s {
		if c == '(' {
			stk = append(stk, i)
		} else if len(stk) > 0 {
			stk = stk[:len(stk)-1]
			l := i - start
			if len(stk) != 0 {
				l = i - stk[len(stk)-1]
			}
			if l > res {
				res = l
			}
		} else {
			start = i
		}
	}
	return
}

func main() {
	var s string

	s = "(()"
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = ")()())"
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = "((()())"
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = ")()())"
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = "()(()()()"
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = "())()()"
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = ")(((((()())()()))()(()))("
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

	s = "(())(()(((((((((("
	fmt.Println(s)
	fmt.Println(longestValidParentheses(s))

}
