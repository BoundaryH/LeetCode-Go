package main

import "fmt"

/* 0ms */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minStr := strs[0]
	for _, s := range strs {
		if len(s) < len(minStr) {
			minStr = s
		}
	}
	for i, c := range minStr {
		for _, s := range strs {
			if s[i] != byte(c) {
				return minStr[:i]
			}
		}
	}
	return minStr
}

func main() {
	var strs []string

	strs = []string{"flower", "flow", "flight"}
	fmt.Println(strs)
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(strs)
	fmt.Println(longestCommonPrefix(strs))
}
