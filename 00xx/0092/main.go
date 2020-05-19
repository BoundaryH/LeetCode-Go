package main

import "fmt"

/* 0ms */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	l := &ListNode{Next: head}
	a := l
	for i := 1; i < m; i++ {
		a = a.Next
	}
	b := a.Next
	start, end := a, b
	for i := m; i <= n; i++ {
		b.Next, a, b = a, b, b.Next
	}
	start.Next = a
	end.Next = b
	return l.Next
}

func main() {
	var head *ListNode
	var m, n int

	head = SliceToList([]int{1, 2, 3, 4, 5})
	m, n = 2, 4
	fmt.Println(head)
	fmt.Println(m, n)
	fmt.Println(reverseBetween(head, m, n))
}
