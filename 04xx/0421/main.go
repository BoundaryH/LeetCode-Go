package main

import "fmt"

/* 44ms */
/*
type Node struct {
	left  *Node
	right *Node
}

func insert(node *Node, n int, mask int) {
	if mask == 0 {
		return
	}
	if n&mask == 0 {
		if node.left == nil {
			node.left = &Node{nil, nil}
		}
		insert(node.left, n, mask>>1)
	} else {
		if node.right == nil {
			node.right = &Node{nil, nil}
		}
		insert(node.right, n, mask>>1)
	}
}

func search(node *Node, n int, mask int) int {
	if mask == 0 {
		return 0
	}
	if node.right == nil || (node.left != nil && n&mask != 0) {
		return search(node.left, n, mask>>1)
	} else {
		return search(node.right, n, mask>>1) + mask
	}
}

func findMaximumXOR(nums []int) int {
	root := Node{nil, nil}
	insert(&root, nums[0], 1<<30)

	max := 0
	for i := 1; i < len(nums); i++ {
		j := nums[i] ^ search(&root, nums[i], 1<<30)
		if j > max {
			max = j
		}
		insert(&root, nums[i], 1<<30)
	}
	return max
}
*/

// a^b = c    =>    a^c = a^a^b = b
/* 44ms */
func findMaximumXOR(nums []int) int {
	mask := 0
	max := 0
	for m := 1 << 30; m != 0; m >>= 1 {
		mask += m
		prefix := make(map[int]bool)
		for _, n := range nums {
			prefix[n&mask] = true
		}

		i := max + m
		for _, n := range nums {
			if prefix[i^(n&mask)] {
				max = i
				break
			}
		}
	}
	return max
}

func main() {
	var nums []int

	nums = []int{3, 10, 5, 25, 2, 8}
	fmt.Println(nums)
	fmt.Println(findMaximumXOR(nums))
}
