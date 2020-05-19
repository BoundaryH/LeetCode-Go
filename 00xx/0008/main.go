package main

import "fmt"

/* 0ms */
func myAtoi(str string) int {
	var res int32
	const INT_MAX = (1 << 31) - 1
	const INT_MIN = (1 << 31) * -1
	minus := false
	i := 0
	for ; i < len(str) && str[i] == ' '; i++ {
	}
	if i < len(str) {
		if str[i] == '-' {
			minus = true
			i++
		} else if str[i] == '+' {
			i++
		}
	}
	for ; i < len(str); i++ {
		n := int32(str[i]) - '0'
		if n < 0 || n > 9 {
			break
		}
		if res < INT_MIN/10 || INT_MIN-res*10+n > 0 {
			res = INT_MIN
			break
		}
		res = res*10 - n
	}
	if !minus {
		if res == INT_MIN {
			return INT_MAX
		}
		res = -res
	}
	return int(res)
}

func main() {
	var str string

	str = "42"
	fmt.Println(str)
	fmt.Println(myAtoi(str))

	str = "     -42"
	fmt.Println(str)
	fmt.Println(myAtoi(str))

	str = "4193 with words"
	fmt.Println(str)
	fmt.Println(myAtoi(str))

	str = "words and 987"
	fmt.Println(str)
	fmt.Println(myAtoi(str))

	str = "abddefg"
	fmt.Println(str)
	fmt.Println(myAtoi(str))
}
