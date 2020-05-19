package main

import "fmt"

/* 0ms */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := &ListNode{Next: head}
	a, b := list, list
	for i := 0; i < n+1; i++ {
		b = b.Next
	}
	for b != nil {
		a = a.Next
		b = b.Next
	}
	a.Next = a.Next.Next
	return list.Next
}

func main() {
	var head *ListNode
	var n int

	head = SliceToList([]int{1, 2, 3, 4, 5})
	n = 2
	fmt.Println(head)
	fmt.Println(n)
	fmt.Println(removeNthFromEnd(head, n))
}
