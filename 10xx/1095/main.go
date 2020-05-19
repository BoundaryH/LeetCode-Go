package main

import "fmt"

/* 4ms */
func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	for l < r {
		m := (l + r) / 2
		if m == 0 {
			m++
		} else if m == mountainArr.length()-1 {
			m--
		}
		a, b, c := mountainArr.get(m-1), mountainArr.get(m), mountainArr.get(m+1)
		if a < b && b > c {
			l, r = m, m
			break
		}
		if b < c {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	peak := l
	l, r = 0, peak
	for l <= r {
		m := (l + r) / 2
		n := mountainArr.get(m)
		if n < target {
			l = m + 1
		} else if n > target {
			r = m - 1
		} else {
			return m
		}
	}
	l, r = peak, mountainArr.length()-1
	for l <= r {
		m := (l + r) / 2
		n := mountainArr.get(m)
		if n > target {
			l = m + 1
		} else if n < target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

type MountainArray struct {
	count int
	nums  []int
}

func (this *MountainArray) get(index int) int {
	this.count++
	return this.nums[index]
}
func (this *MountainArray) length() int {
	return len(this.nums)
}

func main() {
	var target int
	var arr MountainArray

	target = 3
	arr = MountainArray{nums: []int{1, 2, 3, 4, 5, 3, 1}}
	fmt.Println(arr.nums)
	fmt.Println(target, findInMountainArray(target, &arr), arr.count)

	target = 3
	arr = MountainArray{nums: []int{0, 1, 2, 4, 2, 1}}
	fmt.Println(arr.nums)
	fmt.Println(target, findInMountainArray(target, &arr), arr.count)

	target = 2
	arr = MountainArray{nums: []int{1, 2, 3, 4, 5, 3, 1}}
	fmt.Println(arr.nums)
	fmt.Println(target, findInMountainArray(target, &arr), arr.count)

	target = 2
	arr = MountainArray{nums: []int{1, 5, 2}}
	fmt.Println(arr.nums)
	fmt.Println(target, findInMountainArray(target, &arr), arr.count)

}
