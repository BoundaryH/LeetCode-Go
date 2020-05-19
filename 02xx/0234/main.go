package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* 12ms */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	a, b := head, head
	for b != nil && b.Next != nil {
		a, b = a.Next, b.Next.Next
	}
	for a.Next, b = nil, a.Next; b != nil; {
		a, b, b.Next = b, b.Next, a
	}
	for l, r := head, a; r != nil; l, r = l.Next, r.Next {
		if l.Val != r.Val {
			return false
		}
	}
	return true
}

func main() {
	var inp *ListNode

	inp = SliceToList([]int{1, 2})
	fmt.Println(inp)
	fmt.Println(isPalindrome(inp))

	inp = SliceToList([]int{1, 2, 2, 1})
	fmt.Println(inp)
	fmt.Println(isPalindrome(inp))

	inp = SliceToList([]int{1, 2, 3, 2, 1})
	fmt.Println(inp)
	fmt.Println(isPalindrome(inp))

	inp = SliceToList([]int{1, 2, 3, 5, 4, 2, 1})
	fmt.Println(inp)
	fmt.Println(isPalindrome(inp))
}
