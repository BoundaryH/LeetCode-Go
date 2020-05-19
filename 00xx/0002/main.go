package main

import "fmt"

/* 8ms */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	sum := 0
	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			tail.Next = &ListNode{Val: sum - 10}
			sum = 1
		} else {
			tail.Next = &ListNode{Val: sum}
			sum = 0
		}
		tail = tail.Next
	}
	return head.Next
}

func main() {
	l1 := SliceToList([]int{2, 4, 3})
	l2 := SliceToList([]int{5, 6, 4})
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}
