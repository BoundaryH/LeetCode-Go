package main

import "fmt"

/* 0ms */
func countDigitOne(n int) int {
	count := 0
	for i := 1; i <= n; i *= 10 {
		a := n / (i * 10)
		b := n % (i * 10)
		count += a * i
		if b >= i {
			if b < i*2 {
				count += b - i + 1
			} else {
				count += i
			}
		}
	}
	return count
}

/* 0ms */
/*
func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	i := 1
	for ; i*10 <= n; i *= 10 {
	}
	a := n / i
	b := n % i
	count := a*countDigitOne(i-1) + countDigitOne(b)
	if a == 1 {
		count += b + 1
	} else if a > 1 {
		count += i
	}
	return count
}
*/

func main() {
	inp := 13
	fmt.Println(inp, countDigitOne(inp))
}
