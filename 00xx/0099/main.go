package main

import "fmt"

/* 8ms */
func recoverTree(root *TreeNode) {
	var p, q, max *TreeNode

	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		if max != nil && max.Val > n.Val {
			p = max
			q = n
		} else {
			max = n
		}
		inorder(n.Right)
	}
	inorder(root)
	p.Val, q.Val = q.Val, p.Val
}

func main() {
	var root *TreeNode

	root = SliceToTree([]int{1, 3, null, null, 2})
	fmt.Println(root)
	recoverTree(root)
	fmt.Println(root)

	root = SliceToTree([]int{3, 1, 4, null, null, 2})
	fmt.Println(root)
	recoverTree(root)
	fmt.Println(root)
}
