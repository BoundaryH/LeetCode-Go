package main

import (
	"fmt"
	"sort"
)

/* 12ms */
func longestStrChain(words []string) int {
	lenMap := make(map[string]int)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	maxLen := 0
	for _, w := range words {
		n := 1
		for i := 0; i < len(w); i++ {
			if j := lenMap[w[:i]+w[i+1:]] + 1; j > n {
				n = j
			}
		}
		if n > maxLen {
			maxLen = n
		}
		lenMap[w] = n
	}
	return maxLen
}

func main() {
	var words []string

	words = []string{"a", "b", "ba", "bca", "bda", "bdca"}
	fmt.Println(words)
	fmt.Println(longestStrChain(words))
}
