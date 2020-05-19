package main

import "fmt"

/* 0ms */
func removeOuterParentheses(S string) string {
	ret := make([]rune, 0, len(S))
	n := 0
	for _, c := range S {
		if c == '(' {
			if n != 0 {
				ret = append(ret, c)
			}
			n++
		} else {
			n--
			if n != 0 {
				ret = append(ret, c)
			}
		}
	}
	return string(ret)
}

func main() {
	var S string

	S = "(()())(())"
	fmt.Println(S)
	fmt.Println(removeOuterParentheses(S))

	S = "(()())(())(()(()))"
	fmt.Println(S)
	fmt.Println(removeOuterParentheses(S))

	S = "()()"
	fmt.Println(S)
	fmt.Println(removeOuterParentheses(S))

}
