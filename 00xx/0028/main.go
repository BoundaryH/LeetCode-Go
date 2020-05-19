package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	/* 0ms */
	return strings.Index(haystack, needle)
	/* 0ms */
	//return kmp(haystack, needle)
}

func next(needle string) []int {
	next := make([]int, len(needle))
	next[0] = -1
	i, j := 0, -1
	for i < len(needle)-1 {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func kmp(s string, sub string) int {
	if len(sub) == 0 {
		return 0
	}
	next := next(sub)

	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if j == -1 || s[i] == sub[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(sub) {
		return i - j
	}
	return -1
}

func main() {
	var haystack string
	var needle string

	haystack = "hello"
	needle = "ll"
	fmt.Println(haystack, needle, strStr(haystack, needle))

	haystack = "aaaaa"
	needle = "bba"
	fmt.Println(haystack, needle, strStr(haystack, needle))
}
