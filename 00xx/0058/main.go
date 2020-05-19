package main

import "fmt"

/* 0ms */
func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	j := i
	for ; j >= 0 && s[j] != ' '; j-- {
	}
	return i - j
}

func main() {
	var s string

	s = "Hello World"
	fmt.Println(s, lengthOfLastWord(s))
}
