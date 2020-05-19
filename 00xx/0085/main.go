package main

import "fmt"

/* 0ms */
func maximalRectangle(matrix [][]byte) (res int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	col := len(matrix[0])
	height := make([]int, col)
	hStack := make([]int, 0, col)
	pStack := make([]int, 0, col)

	for x, row := range matrix {
		hStack = hStack[:0]
		pStack = pStack[:0]

		for y := range row {
			if matrix[x][y] == '0' {
				height[y] = 0
				hStack = hStack[:0]
				pStack = pStack[:0]
			} else {
				height[y]++
				h := height[y]
				p := y
				for len(hStack) > 0 && hStack[len(hStack)-1] > h {
					i := len(hStack) - 1
					p = pStack[i]
					hStack = hStack[:i]
					pStack = pStack[:i]
				}
				hStack = append(hStack, h)
				pStack = append(pStack, p)
				for k := range hStack {
					if a := (y - pStack[k] + 1) * hStack[k]; a > res {
						res = a
					}
				}
			}
		}
	}
	return
}

func main() {
	var matrix [][]byte

	matrix = [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	for _, r := range matrix {
		fmt.Println(string(r))
	}
	fmt.Println(maximalRectangle(matrix))

	/*
		matrix = [][]byte{
			[]byte("01101"),
			[]byte("11010"),
			[]byte("01110"),
			[]byte("11111"),
			[]byte("00000"),
		}
		for _, r := range matrix {
			fmt.Println(string(r))
		}
		fmt.Println(maximalRectangle(matrix))

		matrix = [][]byte{
			[]byte("0110111111111111110"),
			[]byte("1011111111111111111"),
			[]byte("1101111111110111111"),
			[]byte("1111111111111011111"),
			[]byte("1111111111111101111"),
			[]byte("1110111011111111101"),
			[]byte("1011111111111101111"),
			[]byte("1111111111111110111"),
			[]byte("0011111111111110111"),
			[]byte("1101111111011111111"),
			[]byte("1111111110111111111"),
			[]byte("0110111011111111111"),
			[]byte("1111011111111101111"),
			[]byte("1111111111111111111"),
			[]byte("1111111111111111111"),
			[]byte("1111111111111111101"),
			[]byte("1111111101101101111"),
			[]byte("1111110111111110111"),
		}
		for _, r := range matrix {
			fmt.Println(string(r))
		}
		fmt.Println(maximalRectangle(matrix))
	*/
}
