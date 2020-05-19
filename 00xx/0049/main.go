package main

import (
	"fmt"
)

/* 16 ms */
func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, s := range strs {
		count := make([]rune, 26)
		for _, c := range s {
			count[c-'a']++
		}
		key := string(count)
		group[key] = append(group[key], s)
	}
	res := make([][]string, 0, len(group))
	for _, i := range group {
		res = append(res, i)
	}
	return res
}

func main() {
	var strs []string

	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(strs)
	for _, i := range groupAnagrams(strs) {
		fmt.Println(i)
	}
}
