package main

import (
	"fmt"
	"math"
)

/* 8ms */
/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	left := (total - 1) / 2
	right := left

	a, b := nums1, nums2
	for left > 0 || right > 0 {
		if len(a) != 0 {
			a, b = b, a
		}

		n := b[(len(b)-1)/2]
		l, r := 0, len(a)
		for l < r {
			m := (l + r) / 2
			if a[m] <= n {
				l = m + 1
			} else {
				r = m
			}
		}
		x, y := l, (len(b)-1)/2
		if c := len(a) - x + len(b) - y - 1; c <= right {
			right -= c
			a = a[:x]
			b = b[:y+1]
			if right > 0 {
				right--
				b = b[:y]
			}
		}
		if c := x + y; c <= left {
			left -= c
			a = a[x:]
			b = b[y:]
			if left > 0 {
				left--
				b = b[1:]
			}
		}
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	if total%2 == 0 {
		if len(a) == 0 {
			return float64(b[0]+b[1]) / 2
		} else {
			return float64(a[0]+b[0]) / 2
		}
	}
	return float64(b[0])
}
*/

/* 8ms */
/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	pop := func() (n int) {
		if i == len(nums1) {
			n = nums2[j]
			j++
		} else if j == len(nums2) {
			n = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			n = nums1[i]
			i++
		} else {
			n = nums2[j]
			j++
		}
		return
	}

	if (len(nums1)+len(nums2))%2 == 1 {
		half := (len(nums1) + len(nums2)) / 2
		for t := 0; t < half; t++ {
			pop()
		}
		return float64(pop())
	}
	half := (len(nums1)+len(nums2))/2 - 1
	for t := 0; t < half; t++ {
		pop()
	}
	return (float64(pop()) + float64(pop())) / 2
}
*/

/* 8ms */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	half := (len(nums1) + len(nums2) - 1) / 2

	l, r := 0, len(nums1)
	i := (l + r) / 2
	j := half - i
	for l < r {
		if nums1[i] < nums2[j] {
			l = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j+1] {
			r = i
		} else {
			break
		}
		i = (l + r) / 2
		j = half - i
	}

	var medianLeft, medianRight float64
	if j < 0 {
		medianLeft = float64(nums1[i-1])
	} else if i < 1 {
		medianLeft = float64(nums2[j])
	} else {
		medianLeft = math.Max(float64(nums1[i-1]), float64(nums2[j]))
	}

	if (len(nums1)+len(nums2))%2 == 1 {
		return medianLeft
	} else if i == len(nums1) {
		medianRight = float64(nums2[j+1])
	} else if j+1 == len(nums2) {
		medianRight = float64(nums1[i])
	} else {
		medianRight = math.Min(float64(nums1[i]), float64(nums2[j+1]))
	}
	return (medianLeft + medianRight) / 2
}

/* 16ms */
/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	h := (len(nums) - 1) / 2
	if len(nums)%2 == 1 {
		return float64(nums[h])
	}
	return float64(nums[h]+nums[h+1]) / 2
}
*/

func main() {
	var nums1, nums2 []int

	//nums1 = []int{1, 2, 3, 4}
	//nums2 = []int{3, 5, 6}
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	//nums1 = []int{100001}
	//nums2 = []int{100000}
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//fmt.Println(findMedianSortedArrays(nums1, nums2))

	//nums1 = []int{1, 2}
	//nums2 = []int{-1, 3}
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//fmt.Println(findMedianSortedArrays(nums1, nums2))

	//nums1 = []int{1, 3}
	//nums2 = []int{2}
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//fmt.Println(findMedianSortedArrays(nums1, nums2))

	//nums1 = []int{0, 0}
	//nums2 = []int{0, 0}
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//fmt.Println(findMedianSortedArrays(nums1, nums2))
}
