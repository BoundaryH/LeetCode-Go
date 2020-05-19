package main

import "fmt"

/* 4ms */
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				a := dp[i-1][j] + 1
				b := dp[i][j-1] + 1
				c := dp[i-1][j-1] + 1
				if a < b {
					dp[i][j] = a
				} else {
					dp[i][j] = b
				}
				if c < dp[i][j] {
					dp[i][j] = c
				}
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	var word1, word2 string

	word1 = "horse"
	word2 = "ros"
	fmt.Println(word1)
	fmt.Println(word2)
	fmt.Println(minDistance(word1, word2))

	word1 = "intention"
	word2 = "execution"
	fmt.Println(word1)
	fmt.Println(word2)
	fmt.Println(minDistance(word1, word2))
}
