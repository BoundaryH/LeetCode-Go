package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type node struct {
	v int
	x int
	y int
}

/* 0ms */
func verticalTraversal(root *TreeNode) [][]int {
	nodes := []*node{}
	nodes = dfs(root, 0, 0, nodes)
	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.x < b.x ||
			(a.x == b.x && a.y > b.y) ||
			(a.x == b.x && a.y == b.y && a.v < b.v)
	})
	res := [][]int{}
	for i := 0; i < len(nodes); {
		col := []int{}
		x := nodes[i].x
		for i < len(nodes) && nodes[i].x == x {
			col = append(col, nodes[i].v)
			i++
		}
		res = append(res, col)
	}
	return res
}

func dfs(root *TreeNode, x, y int, nodes []*node) []*node {
	if root == nil {
		return nodes
	}
	nodes = append(nodes, &node{root.Val, x, y})
	nodes = dfs(root.Left, x-1, y-1, nodes)
	nodes = dfs(root.Right, x+1, y-1, nodes)
	return nodes
}

func main() {
	var root *TreeNode

	root = SliceToTree([]int{3, 9, 20, null, null, 15, 7})
	fmt.Println(root)
	fmt.Println(verticalTraversal(root))

	root = SliceToTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(root)
	fmt.Println(verticalTraversal(root))

	root = SliceToTree([]int{0, 2, 1, 3, null, null, null, 4, 5, null, 7, 6, null, 10, 8, 11, 9})
	fmt.Println(root)
	fmt.Println(verticalTraversal(root))
}
