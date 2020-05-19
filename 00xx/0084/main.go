package main

import "fmt"

/* 4ms */
func largestRectangleArea(heights []int) (res int) {
	hStack := make([]int, 0, len(heights))
	pStack := make([]int, 0, len(heights))

	for i, h := range heights {
		idx := i
		for len(hStack) > 0 && hStack[len(hStack)-1] >= h {
			j := len(hStack) - 1
			if a := hStack[j] * (i - pStack[j]); a > res {
				res = a
			}
			idx = pStack[j]
			hStack = hStack[:j]
			pStack = pStack[:j]
		}
		hStack = append(hStack, h)
		pStack = append(pStack, idx)
	}
	for len(hStack) > 0 {
		j := len(hStack) - 1
		if a := hStack[j] * (len(heights) - pStack[j]); a > res {
			res = a
		}
		hStack = hStack[:j]
		pStack = pStack[:j]
	}
	return
}

type Stack []int

func (stk Stack) Last() int {
	return stk[len(stk)-1]
}

func (stk *Stack) Push(n int) {
	*stk = append(*stk, n)
}

func main() {
	var heights []int

	heights = []int{2, 1, 5, 6, 2, 3}
	fmt.Println(heights)
	fmt.Println(largestRectangleArea(heights))

	heights = []int{1, 1}
	fmt.Println(heights)
	fmt.Println(largestRectangleArea(heights))
}
