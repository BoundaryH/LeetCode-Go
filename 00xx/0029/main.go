package main

import "fmt"

/* 4ms */
func divide(dividend int, divisor int) int {
	const INT_MIN = (1 << 31) * -1
	const INT_MAX = (1 << 31) - 1
	if dividend == 0 {
		return 0
	}
	sign := (dividend > 0 && divisor < 0) ||
		(dividend < 0 && divisor > 0)
	a, b := int32(dividend), int32(divisor)
	res := int32(0)
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}
	p := b
	for p > a && p >= INT_MIN/2 {
		p *= 2
	}
	for p <= b {
		res = res * 2
		if a <= p {
			res--
			a -= p
		}
		p /= 2
	}
	if !sign {
		if res == INT_MIN {
			res = INT_MAX
		} else {
			res = -res
		}
	}
	return int(res)
}

/* 4ms */
/*
func divide(dividend int, divisor int) int {
	const INT_MIN = (1 << 31) * -1
	const INT_MAX = (1 << 31) - 1
	r := dividend / divisor
	if r > INT_MAX {
		r = INT_MAX
	} else if r < INT_MIN {
		r = INT_MIN
	}
	return r
}
*/

func main() {
	var dividend, divisor int

	dividend = 10
	divisor = 3
	fmt.Println(dividend, divisor, divide(dividend, divisor))
}
