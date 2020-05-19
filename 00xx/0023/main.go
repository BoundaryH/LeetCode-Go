package main

import (
	"fmt"
)

/* 8ms */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	var a, b *ListNode
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		a = lists[0]
		b = lists[1]
	default:
		a = mergeKLists(lists[:n/2])
		b = mergeKLists(lists[n/2:])
	}

	head := &ListNode{}
	tail := head
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a
			tail = a
			a = a.Next
		} else {
			tail.Next = b
			tail = b
			b = b.Next
		}
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return head.Next
}

/* 12ms */
/*
func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListHeap{}
	heap.Init(h)
	for _, l := range lists {
		heap.Push(h, l)
	}

	head := &ListNode{}
	tail := head
	for h.Len() > 0 {
		n := heap.Pop(h).(*ListNode)
		tail.Next = n
		tail = n
		n = n.Next
		heap.Push(h, n)
	}
	return head.Next
}

type ListHeap []*ListNode

func (h ListHeap) Len() int {
	return len(h)
}
func (h ListHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h ListHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListHeap) Push(x interface{}) {
	n := x.(*ListNode)
	if n != nil {
		*h = append(*h, n)
	}
}

func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
*/

func main() {
	var lists []*ListNode

	lists = []*ListNode{
		SliceToList([]int{1, 4, 5}),
		SliceToList([]int{1, 3, 4}),
		SliceToList([]int{2, 6}),
	}
	for _, l := range lists {
		fmt.Println(l)
	}
	fmt.Println(mergeKLists(lists))

	lists = []*ListNode{
		SliceToList([]int{1}),
		SliceToList([]int{0}),
	}
	for _, l := range lists {
		fmt.Println(l)
	}
	fmt.Println(mergeKLists(lists))
}
