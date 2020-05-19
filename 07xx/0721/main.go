package main

import (
	"fmt"
	"sort"
)

/* 56ms */
func accountsMerge(accounts [][]string) [][]string {
	set := make([]int, len(accounts))
	hashMap := make(map[string]int)
	for i, a := range accounts {
		set[i] = i
		for j := 1; j < len(a); j++ {
			if v, ok := hashMap[a[j]]; !ok {
				hashMap[a[j]] = i
			} else {
				for v != set[v] {
					v = set[v]
				}
				set[v] = i
			}
		}
	}
	res := make([][]string, 0)
	idx := make(map[int]int)
	for s, i := range hashMap {
		for i != set[i] {
			i = set[i]
		}
		if n, ok := idx[i]; ok {
			res[n] = append(res[n], s)
		} else {
			idx[i] = len(res)
			res = append(res, []string{accounts[i][0], s})
		}
	}
	for i := 0; i < len(res); i++ {
		sort.Strings(res[i][1:])
	}
	return res
}

func main() {
	var accounts [][]string

	accounts = [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"}}
	for _, row := range accounts {
		fmt.Println(row)
	}
	fmt.Println()
	for _, row := range accountsMerge(accounts) {
		fmt.Println(row)
	}

}
