package main

import "fmt"

/* 4ms */
func deleteDuplicates(head *ListNode) *ListNode {
	start := &ListNode{Next: head}
	i, j := start, head
	for j != nil {
		k := j.Next
		for k != nil && k.Val == j.Val {
			k = k.Next
		}
		if k == j.Next {
			i = j
			j = k
		} else {
			i.Next = k
			j = k
		}
	}
	return start.Next
}

func main() {
	var head *ListNode

	head = SliceToList([]int{1, 2, 3, 3, 4, 4, 5})
	fmt.Println(head)
	fmt.Println(deleteDuplicates(head))

	head = SliceToList([]int{1, 1, 1, 2, 3})
	fmt.Println(head)
	fmt.Println(deleteDuplicates(head))
}
