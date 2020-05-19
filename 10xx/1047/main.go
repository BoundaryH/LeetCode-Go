package main

import "fmt"

/* 0ms */
func removeDuplicates(S string) string {
	stk := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		l := len(stk)
		if l == 0 || stk[l-1] != S[i] {
			stk = append(stk, S[i])
		} else {
			stk = stk[:l-1]
		}
	}
	return string(stk)
}

func main() {
	var S string

	S = "abbaca"
	fmt.Println(S, removeDuplicates(S))
}
