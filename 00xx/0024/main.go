package main

import "fmt"

/* 0ms */
func swapPairs(head *ListNode) *ListNode {
	l := &ListNode{Next: head}
	p := l
	for p.Next != nil && p.Next.Next != nil {
		a := p.Next
		b := a.Next

		a.Next = b.Next
		b.Next = a
		p.Next = b
		p = a
	}
	return l.Next
}

func main() {
	var head *ListNode

	head = SliceToList([]int{1, 2, 3, 4})
	fmt.Println(head)
	fmt.Println(swapPairs(head))
}
