package main

import "fmt"

/* 0ms */
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(s1); i++ {
		dp[i+1][0] = (dp[i][0] && s1[i] == s3[i])
	}
	for j := 0; j < len(s2); j++ {
		dp[0][j+1] = (dp[0][j] && s2[j] == s3[j])
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			c := s3[i+j-1]
			dp[i][j] = (dp[i-1][j] && s1[i-1] == c) ||
				(dp[i][j-1] && s2[j-1] == c)
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	var s1, s2, s3 string

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbcbcac"
	fmt.Println(s1, s2)
	fmt.Println(s3)
	fmt.Println(isInterleave(s1, s2, s3))

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	fmt.Println(s1, s2)
	fmt.Println(s3)
	fmt.Println(isInterleave(s1, s2, s3))
}
