package main

import "fmt"

/* 188ms */
func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	sum = dfs(root.Right, sum)
	sum += root.Val
	root.Val = sum
	sum = dfs(root.Left, sum)
	return sum
}

func main() {
	var nums []int
	var root *TreeNode

	nums = []int{5, 2, 13}
	root = SliceToTree(nums)
	fmt.Println(root)
	fmt.Println(convertBST(root))

	nums = []int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}
	root = SliceToTree(nums)
	fmt.Println(root)
	fmt.Println(convertBST(root))

}
