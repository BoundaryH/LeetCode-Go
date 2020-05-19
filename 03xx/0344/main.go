package main

import "fmt"

/* 40ms */
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	var str string
	var buf []byte

	str = "hello"
	fmt.Println(str)
	buf = []byte(str)
	reverseString(buf)
	fmt.Println(string(buf))

	str = "Hannah"
	fmt.Println(str)
	buf = []byte(str)
	reverseString(buf)
	fmt.Println(string(buf))
}
