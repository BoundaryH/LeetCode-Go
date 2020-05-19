package main

import "fmt"

/* 0ms */
func partition(head *ListNode, x int) *ListNode {
	lh := &ListNode{}
	rh := &ListNode{}
	lt, rt := lh, rh

	n := head
	for n != nil {
		if n.Val < x {
			lt.Next = n
			lt = n
		} else {
			rt.Next = n
			rt = n
		}
		n = n.Next
	}
	lt.Next = rh.Next
	rt.Next = nil
	return lh.Next
}

func main() {
	var head *ListNode
	var x int

	head = SliceToList([]int{1, 4, 3, 2, 5, 2})
	x = 3
	fmt.Println(head)
	fmt.Println(x)
	fmt.Println(partition(head, x))
}
