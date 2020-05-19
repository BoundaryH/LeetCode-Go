package main

import "fmt"

/* 4ms */
func isValidBST(root *TreeNode) (res bool) {
	val, init := 0, false
	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if !res || n == nil {
			return
		}
		inorder(n.Left)
		if init && n.Val <= val {
			res = false
		}
		val, init = n.Val, true
		inorder(n.Right)
	}
	res = true
	inorder(root)
	return
}

func main() {
	var root *TreeNode

	root = SliceToTree([]int{2, 1, 3})
	fmt.Println(root)
	fmt.Println(isValidBST(root))

	root = SliceToTree([]int{5, 1, 4, null, null, 3, 6})
	fmt.Println(root)
	fmt.Println(isValidBST(root))
}
