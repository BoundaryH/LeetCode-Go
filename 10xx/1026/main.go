package main

import "fmt"

/* 4ms */
func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, maxV, minV int) int {
	if root == nil {
		return maxV - minV
	}
	if root.Val > maxV {
		maxV = root.Val
	}
	if root.Val < minV {
		minV = root.Val
	}
	l := dfs(root.Left, maxV, minV)
	r := dfs(root.Right, maxV, minV)
	if l > r {
		return l
	}
	return r
}

func main() {
	var nums []int
	var root *TreeNode

	nums = []int{8, 3, 10, 1, 6, null, 14, null, null, 4, 7, 13}
	root = SliceToTree(nums)
	fmt.Println(root)
	fmt.Println(maxAncestorDiff(root))

}
