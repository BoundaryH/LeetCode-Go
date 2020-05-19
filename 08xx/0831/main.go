package main

import (
	"bytes"
	"fmt"
	"strings"
)

/* 0ms */
func maskPII(S string) string {
	var res bytes.Buffer
	if strings.Contains(S, "@") {
		a := strings.IndexByte(S, '@')
		res.WriteString(strings.ToLower(S[:1]))
		res.WriteString("*****")
		res.WriteString(strings.ToLower(S[a-1 : a]))
		res.WriteByte('@')
		res.WriteString(strings.ToLower(S[a+1:]))
	} else {
		nums := make([]byte, 0, 13)
		for _, c := range S {
			if c >= '0' && c <= '9' {
				nums = append(nums, byte(c))
			}
		}
		if len(nums) > 10 {
			res.WriteByte('+')
			for i := len(nums) - 10; i > 0; i-- {
				res.WriteByte('*')
			}
			res.WriteByte('-')
		}
		res.WriteString("***-***-")
		res.Write(nums[len(nums)-4:])

	}
	return res.String()
}

func main() {
	var S string

	S = "LeetCode@LeetCode.com"
	fmt.Println(S)
	fmt.Println(maskPII(S))

	S = "AB@qq.com"
	fmt.Println(S)
	fmt.Println(maskPII(S))

	S = "1(234)567-890"
	fmt.Println(S)
	fmt.Println(maskPII(S))

	S = "86-(10)12345678"
	fmt.Println(S)
	fmt.Println(maskPII(S))
}
