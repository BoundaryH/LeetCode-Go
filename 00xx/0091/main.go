package main

import "fmt"

/* 0ms */
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			dp[i+1] = dp[i]
		}
		n := int((s[i-1]-'0')*10 + s[i] - '0')
		if s[i-1] != '0' && n <= 26 && n > 0 {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(s)]
}

func main() {
	var s string

	s = "12"
	fmt.Println(s, numDecodings(s))

	s = "226"
	fmt.Println(s, numDecodings(s))

	s = "01"
	fmt.Println(s, numDecodings(s))
}
