package main

import "fmt"

/* 0ms */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	node := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}
	return head.Next
}

func main() {
	l1 := SliceToList([]int{1, 2, 4})
	l2 := SliceToList([]int{1, 3, 4})
	ret := mergeTwoLists(l1, l2)
	fmt.Println(ret)
}
