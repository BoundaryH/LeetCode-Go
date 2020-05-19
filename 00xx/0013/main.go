package main

import "fmt"

/* 0ms */
func romanToInt(s string) int {
	char2num := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	sum := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		n := char2num[s[i]]
		if n < prev {
			sum -= n
		} else {
			sum += n
		}
		prev = n
	}
	return sum
}

func main() {
	var s string

	s = "III"
	fmt.Printf("%s  -> %d\n", s, romanToInt(s))
	s = "IV"
	fmt.Printf("%s  -> %d\n", s, romanToInt(s))
	s = "IX"
	fmt.Printf("%s  -> %d\n", s, romanToInt(s))
	s = "LVIII"
	fmt.Printf("%s  -> %d\n", s, romanToInt(s))
	s = "MCMXCIV"
	fmt.Printf("%s  -> %d\n", s, romanToInt(s))
}
