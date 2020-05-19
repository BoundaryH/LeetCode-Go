package main

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	buf := bytes.NewBufferString("")
	for p := head; p != nil; p = p.Next {
		if p != head {
			fmt.Fprintf(buf, " -> %d", p.Val)
		} else {
			fmt.Fprintf(buf, "%d", p.Val)
		}
	}
	return buf.String()
}

func SliceToList(nums []int) *ListNode {
	head := ListNode{0, nil}
	node := &head
	for _, v := range nums {
		node.Next = &ListNode{v, nil}
		node = node.Next
	}
	return head.Next
}

func ListToSlice(head *ListNode) []int {
	node := head
	nums := make([]int, 0)
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	return nums
}
