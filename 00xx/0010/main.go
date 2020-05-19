package main

import "fmt"

/* 0ms */
func isMatch(s string, p string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(p); i++ {
		tmp := dp
		dp = make([]bool, len(s)+1)

		star := (i+1 < len(p) && p[i+1] == '*')
		dp[0] = (star && tmp[0])
		for j := range s {
			if p[i] == s[j] || p[i] == '.' {
				dp[j+1] = tmp[j] || (star && dp[j]) || (star && tmp[j+1])
			} else if star {
				dp[j+1] = tmp[j+1]
			} else {
				dp[j+1] = false
			}
		}
		if star {
			i++
		}
	}
	return dp[len(s)]
}

func main() {
	var s, p string

	s = "aa"
	p = "a"
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))

	s = "aa"
	p = "a*"
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))

	s = "ab"
	p = ".*"
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))

	s = "aab"
	p = "c*a*b"
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))

	s = "mississippi"
	p = "mis*is*p*."
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))

	s = "mississippi"
	p = "mis*is*ip*."
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))

	s = "aasdfasdfasdfasdfas"
	p = "aasdf.*asdf.*asdf.*asdf.*s"
	fmt.Println(s, p)
	fmt.Println(isMatch(s, p))
}
