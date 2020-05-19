package main

import "fmt"

/* 0ms */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	size := 1
	for tail.Next != nil {
		tail = tail.Next
		size++
	}

	tail.Next = head
	for i := (size - k%size); i > 0; i-- {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}

func main() {
	var head *ListNode
	var k int

	head = SliceToList([]int{1, 2, 3, 4, 5})
	k = 2
	fmt.Println(head)
	fmt.Println(k)
	fmt.Println(rotateRight(head, k))
	fmt.Println()

	head = SliceToList([]int{0, 1, 2})
	k = 4
	fmt.Println(head)
	fmt.Println(k)
	fmt.Println(rotateRight(head, k))
	fmt.Println()
}
