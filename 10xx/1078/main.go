package main

import (
	"fmt"
	"strings"
)

/* 0ms */
func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	res := []string{}
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			res = append(res, words[i])
		}
	}
	return res
}

func main() {
	var text, first, second string

	text = "alice is a good girl she is a good student"
	first = "a"
	second = "good"
	fmt.Println(text)
	fmt.Println(first, second)
	fmt.Println(findOcurrences(text, first, second))

	text = "we will we will rock you"
	first = "we"
	second = "will"
	fmt.Println(text)
	fmt.Println(first, second)
	fmt.Println(findOcurrences(text, first, second))
}
