package main

import "fmt"

/* 0ms */
func findSubstringInWraproundString(p string) int {
	if len(p) == 0 {
		return 0
	}
	length := [26]int{}
	i, j := 0, 1
	update := func() {
		for ; i < j; i++ {
			c := p[i] - 'a'
			if j-i > length[c] {
				length[c] = j - i
			}
		}
	}
	for ; j < len(p); j++ {
		if p[j-1]+1 != p[j] && (p[j-1] != 'z' || p[j] != 'a') {
			update()
		}
	}
	update()
	sum := 0
	for _, n := range length {
		sum += n
	}
	return sum
}

func main() {
	var p string

	p = "a"
	fmt.Println(p)
	fmt.Println(findSubstringInWraproundString(p))

	p = "cac"
	fmt.Println(p)
	fmt.Println(findSubstringInWraproundString(p))

	p = "zab"
	fmt.Println(p)
	fmt.Println(findSubstringInWraproundString(p))
}
