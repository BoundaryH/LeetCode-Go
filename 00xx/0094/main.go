package main

import "fmt"

/* 0ms */
func inorderTraversal(root *TreeNode) (res []int) {
	stk := make([]*TreeNode, 0)
	for n := root; n != nil; n = n.Left {
		stk = append(stk, n)
	}
	for len(stk) > 0 {
		n := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, n.Val)

		n = n.Right
		for ; n != nil; n = n.Left {
			stk = append(stk, n)
		}
	}
	return
}

/* 0ms */
/*
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return
}
*/

func main() {
	var root *TreeNode

	root = SliceToTree([]int{1, null, 2, 3})
	fmt.Println(root)
	fmt.Println(inorderTraversal(root))

	root = SliceToTree([]int{1, null, 2, null, null, 3})
	fmt.Println(root)
	fmt.Println(inorderTraversal(root))

	root = SliceToTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(root)
	fmt.Println(inorderTraversal(root))
}
