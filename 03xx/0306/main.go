package main

import "fmt"

/* 0ms */
func isAdditiveNumber(num string) bool {
	for i := len(num) - 1; i > 0; i-- {
		a := num[i:]
		if len(a) > 1 && a[0] == '0' {
			continue
		}
		for j := i - 1; j > 0 && i-j <= len(a); j-- {
			b := num[j:i]
			if len(b) > 1 && b[0] == '0' {
				continue
			}
			if next(num[:j], a, b) {
				return true
			}
		}
	}
	return false
}

func next(num, a, b string) bool {
	if len(num) == 0 {
		return true
	}
	c := strSub(a, b)
	i := len(num) - len(c)
	return i >= 0 && num[i:] == c && next(num[:i], b, c)
}

func strSub(a, b string) string {
	c := make([]byte, len(a))

	i, j := len(a)-1, len(b)-1
	s := byte(0)
	for ; i >= 0; i-- {
		c[i] = a[i] - s
		if j >= 0 {
			c[i] -= b[j] - '0'
			j--
		}
		if c[i] < '0' {
			c[i] += 10
			s = 1
		} else {
			s = 0
		}
	}
	if s != 0 {
		return "-"
	}
	for i++; i < len(c)-1 && c[i] == '0'; i++ {
	}
	return string(c[i:])
}

func main() {
	var num string

	num = "112358"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "199100199"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "123456579"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "11235813"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "112358130"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "0123"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "1101111"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "0000"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "1023"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "8917"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))

	num = "121202436"
	fmt.Println(num)
	fmt.Println(isAdditiveNumber(num))
}
