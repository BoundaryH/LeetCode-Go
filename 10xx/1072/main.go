package main

import "fmt"

/* 192ms */
func maxEqualRowsAfterFlips(matrix [][]int) int {
	strCount := make(map[string]int)
	for _, row := range matrix {
		s1 := make([]byte, len(row))
		s2 := make([]byte, len(row))
		for i, n := range row {
			s1[i] = byte(n + '0')
			s2[i] = byte(1 - n + '0')
		}
		strCount[string(s1)]++
		strCount[string(s2)]++
	}
	res := 0
	for _, n := range strCount {
		if n > res {
			res = n
		}
	}
	return res
}

func main() {
	var matrix [][]int

	matrix = [][]int{{0, 1}, {1, 1}}
	fmt.Println(matrix)
	fmt.Println(maxEqualRowsAfterFlips(matrix))

	matrix = [][]int{{0, 1}, {1, 0}}
	fmt.Println(matrix)
	fmt.Println(maxEqualRowsAfterFlips(matrix))

	matrix = [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}
	fmt.Println(matrix)
	fmt.Println(maxEqualRowsAfterFlips(matrix))
}
