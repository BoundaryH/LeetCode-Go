package main

import "fmt"

/* 0ms */
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	a := len(str1)
	b := len(str2)
	for b != 0 {
		a, b = b, a%b
	}
	return str1[:a]
}

/* 0ms */
/*
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	if len(str2) == 0 {
		return str1
	}
	if str1[:len(str2)] == str2 {
		return gcdOfStrings(str1[len(str2):], str2)
	}
	return ""
}
*/

func main() {
	var str1 string
	var str2 string

	str1 = "ABCABC"
	str2 = "ABC"
	fmt.Println(str1, str2)
	fmt.Println(gcdOfStrings(str1, str2))

	str1 = "ABABAB"
	str2 = "ABAB"
	fmt.Println(str1, str2)
	fmt.Println(gcdOfStrings(str1, str2))

	str1 = "LEET"
	str2 = "CODE"
	fmt.Println(str1, str2)
	fmt.Println(gcdOfStrings(str1, str2))
}
