package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* 4ms */
func removeElements(head *ListNode, val int) *ListNode {
	res := &ListNode{}
	tail := res
	for n := head; n != nil; n = n.Next {
		if n.Val != val {
			tail.Next = n
			tail = n
		}
	}
	tail.Next = nil
	return res.Next
}

func main() {
	inp := SliceToList([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	fmt.Println(inp, val)
	res := removeElements(inp, val)
	fmt.Println(res)
}
