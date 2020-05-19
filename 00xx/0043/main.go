package main

import "fmt"

/* 0ms */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	mul := make([]int, len1+len2)

	for i := len1 - 1; i >= 0; i-- {
		k := len2 + i
		for j := len2 - 1; j >= 0; j-- {
			mul[k] += int(num1[i]-'0') * int(num2[j]-'0')
			k--
		}
	}

	sum := 0
	res := make([]byte, len1+len2)
	for i := len1 + len2 - 1; i >= 0; i-- {
		sum += mul[i]
		res[i] = byte(sum%10 + '0')
		sum /= 10
	}
	if res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}

func main() {
	var num1, num2 string

	num1 = "123"
	num2 = "456"
	fmt.Println(num1, num2)
	fmt.Println(multiply(num1, num2))

	num1 = "123"
	num2 = "199"
	fmt.Println(num1, num2)
	fmt.Println(multiply(num1, num2))
}
