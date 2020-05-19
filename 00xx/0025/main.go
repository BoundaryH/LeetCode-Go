package main

import "fmt"

/* 4ms */
func reverseKGroup(head *ListNode, k int) *ListNode {
	l := &ListNode{Next: head}
	for p := l; ; {
		q := p
		for i := 0; i < k && q != nil; i++ {
			q = q.Next
		}
		if q == nil {
			break
		}

		a := q.Next
		b := p.Next
		p.Next = q
		p = b
		for i := 0; i < k; i++ {
			b.Next, a, b = a, b, b.Next
		}
	}
	return l.Next
}

func main() {
	var head *ListNode
	var k int

	head = SliceToList([]int{1, 2, 3, 4, 5})
	k = 2
	fmt.Println(k, head)
	fmt.Println(reverseKGroup(head, k))
}
