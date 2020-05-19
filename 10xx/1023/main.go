package main

import "fmt"

/* 0ms */
func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = match(q, pattern)
	}
	return res
}

func match(querie, pattern string) bool {
	i, j := 0, 0
	for ; i < len(querie) && j < len(pattern); i++ {
		if querie[i] == pattern[j] {
			j++
		} else if querie[i] < 'a' || querie[i] > 'z' {
			return false
		}
	}
	if j != len(pattern) {
		return false
	}
	for ; i < len(querie); i++ {
		if querie[i] < 'a' || querie[i] > 'z' {
			return false
		}
	}
	return true
}

func main() {
	var queries []string
	var pattern string

	queries = []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern = "FB"
	fmt.Println(pattern, queries)
	fmt.Println(camelMatch(queries, pattern))

	queries = []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern = "FoBa"
	fmt.Println(pattern, queries)
	fmt.Println(camelMatch(queries, pattern))

	queries = []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern = "FoBaT"
	fmt.Println(pattern, queries)
	fmt.Println(camelMatch(queries, pattern))
}
