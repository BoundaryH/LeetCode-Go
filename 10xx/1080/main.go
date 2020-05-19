package main

import "fmt"

/* 48ms */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		}
		return root
	}
	root.Left = sufficientSubset(root.Left, limit-root.Val)
	root.Right = sufficientSubset(root.Right, limit-root.Val)
	if root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func main() {
	var nums []int
	var limit int
	var root *TreeNode

	nums = []int{1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14}
	limit = 1
	root = SliceToTree(nums)
	fmt.Println(nums, limit)
	fmt.Println(root)
	fmt.Println(sufficientSubset(root, limit))

	nums = []int{5, 4, 8, 11, null, 17, 4, 7, 1, null, null, 5, 3}
	limit = 22
	root = SliceToTree(nums)
	fmt.Println(nums, limit)
	fmt.Println(root)
	fmt.Println(sufficientSubset(root, limit))

	nums = []int{5, -6, -6}
	limit = 0
	root = SliceToTree(nums)
	fmt.Println(nums, limit)
	fmt.Println(root)
	fmt.Println(sufficientSubset(root, limit))

	nums = []int{1, 2, -3, -5, null, 4, null}
	limit = -1
	root = SliceToTree(nums)
	fmt.Println(nums, limit)
	fmt.Println(root)
	fmt.Println(sufficientSubset(root, limit))
}
