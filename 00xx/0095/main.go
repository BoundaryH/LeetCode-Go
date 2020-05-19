package main

import "fmt"

/* 4ms */
func generateTrees(n int) []*TreeNode {
	var generate func(i, j int) []*TreeNode
	generate = func(begin, end int) (res []*TreeNode) {
		if begin == end {
			return []*TreeNode{nil}
		}
		for i := begin; i < end; i++ {
			left := generate(begin, i)
			right := generate(i+1, end)
			for _, l := range left {
				for _, r := range right {
					res = append(res, &TreeNode{
						Val:   i,
						Left:  l,
						Right: r,
					})
				}
			}
		}
		return
	}
	if n == 0 {
		return nil
	}
	return generate(1, n+1)
}

func main() {
	var n int

	n = 3
	fmt.Println(n)
	for _, t := range generateTrees(n) {
		fmt.Println(t)
	}
}
