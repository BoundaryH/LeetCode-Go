package main

import "fmt"

/* 0ms */
func merge(nums1 []int, m int, nums2 []int, n int) {
	a, b, c := m-1, n-1, m+n-1
	for a >= 0 && b >= 0 {
		if nums1[a] > nums2[b] {
			nums1[c] = nums1[a]
			a--
		} else {
			nums1[c] = nums2[b]
			b--
		}
		c--
	}
	for b >= 0 {
		nums1[c] = nums2[b]
		b--
		c--
	}
}

func main() {
	var nums1, nums2 []int
	var m, n int

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	m, n = 3, 3
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(m, n)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
