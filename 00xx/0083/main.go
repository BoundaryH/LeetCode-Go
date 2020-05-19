package main

import "fmt"

/* 0ms */
func deleteDuplicates(head *ListNode) *ListNode {
	for i := head; i != nil; i = i.Next {
		if i.Next != nil && i.Val == i.Next.Val {
			j := i.Next
			for j != nil && i.Val == j.Val {
				j = j.Next
			}
			i.Next = j
		}
	}
	return head
}

func main() {
	var head *ListNode

	head = SliceToList([]int{1, 1, 2})
	fmt.Println(head)
	fmt.Println(deleteDuplicates(head))

	head = SliceToList([]int{1, 1, 2, 3, 3})
	fmt.Println(head)
	fmt.Println(deleteDuplicates(head))
}
