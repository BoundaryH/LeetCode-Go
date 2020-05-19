package main

import "fmt"

/* 8ms */
func recoverFromPreorder(S string) *TreeNode {
	depth := make([]int, 0)
	nums := make([]int, 0)
	for i := 0; i < len(S); {
		j := i
		for ; S[j] == '-'; j++ {
		}
		depth = append(depth, j-i)
		n := 0
		for ; j < len(S) && S[j] != '-'; j++ {
			n = n*10 + int(S[j]-'0')
		}
		nums = append(nums, n)
		i = j
	}
	root := &TreeNode{nums[0], nil, nil}
	dfs(root, depth, nums, 1, 1)
	return root
}

func dfs(node *TreeNode, depth []int, nums []int, idx int, d int) int {
	if idx < len(nums) && depth[idx] == d {
		node.Left = &TreeNode{nums[idx], nil, nil}
		idx = dfs(node.Left, depth, nums, idx+1, d+1)
	}
	if idx < len(nums) && depth[idx] == d {
		node.Right = &TreeNode{nums[idx], nil, nil}
		idx = dfs(node.Right, depth, nums, idx+1, d+1)
	}
	return idx
}

func main() {
	var S string

	S = "1-2--3--4-5--6--7"
	fmt.Println(S)
	fmt.Println(recoverFromPreorder(S))

	S = "1-2--3---4-5--6---7"
	fmt.Println(S)
	fmt.Println(recoverFromPreorder(S))

	S = "1-401--349---90--88"
	fmt.Println(S)
	fmt.Println(recoverFromPreorder(S))

}
