package main

import "fmt"

/* 0ms */
func letterCombinations(digits string) (res []string) {
	if len(digits) == 0 {
		return
	}
	keys := []string{
		"", "abc", "def",
		"ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz"}

	var dfs func(string, []byte)
	dfs = func(d string, s []byte) {
		if len(d) == 0 {
			res = append(res, string(s))
			return
		}
		for _, c := range keys[d[0]-'0'-1] {
			dfs(d[1:], append(s, byte(c)))
		}
	}
	dfs(digits, make([]byte, 0, len(digits)))
	return
}

func main() {
	var digits string

	digits = "23"
	fmt.Println(digits)
	for _, s := range letterCombinations(digits) {
		fmt.Println("    ", s)
	}
}
