package main

import "fmt"

/* 0ms */
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(root *TreeNode, num int) int {
	const MOD = 1000000007
	num = (num*2 + root.Val) % MOD
	if root.Left == nil && root.Right == nil {
		return num
	}
	sum := 0
	if root.Left != nil {
		sum += dfs(root.Left, num)
	}
	if root.Right != nil {
		sum += dfs(root.Right, num)
	}
	return sum % MOD
}

func main() {
	var nums []int
	var root *TreeNode

	nums = []int{1, 0, 1, 0, 1, 0, 1}
	root = SliceToTree(nums)
	fmt.Println(root)
	fmt.Println(nums)
	fmt.Println(sumRootToLeaf(root))
}
