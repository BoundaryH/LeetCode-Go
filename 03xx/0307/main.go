package main

import "fmt"

type NumArray struct {
	ptr  int
	tree []int
}

/* 40ms */
func Constructor(nums []int) NumArray {
	count := 2
	for n := len(nums); n > 1; n = (n + 1) >> 1 {
		count += n
	}
	obj := NumArray{
		ptr:  count - len(nums),
		tree: make([]int, count),
	}
	i := obj.ptr
	for _, n := range nums {
		obj.tree[i] = n
		i++
	}
	for i := len(obj.tree) - 1; i > 0; i-- {
		obj.tree[i>>1] += obj.tree[i]
	}
	return obj
}

func (this *NumArray) Update(i int, val int) {
	i += this.ptr
	diff := val - this.tree[i]
	for i > 0 {
		this.tree[i] += diff
		i >>= 1
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	l, r := i+this.ptr, j+this.ptr
	for l <= r {
		if l&1 == 1 {
			sum += this.tree[l]
			l++
		}
		if r&1 == 0 {
			sum += this.tree[r]
			r--
		}
		l >>= 1
		r >>= 1
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

func main() {
	var obj NumArray
	var nums []int

	nums = []int{1, 3, 5}
	fmt.Println(nums)
	obj = Constructor(nums)
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))

	nums = []int{9, -8}
	fmt.Println(nums)
	obj = Constructor(nums)
	obj.Update(0, 3)
	fmt.Println(obj.SumRange(1, 1))
	fmt.Println(obj.SumRange(0, 1))
	obj.Update(1, -3)
	fmt.Println(obj.SumRange(0, 1))

	nums = []int{-28, 39, 53, 65, 11, -56, -65, -39, -43, 97}
	fmt.Println(nums)
	obj = Constructor(nums)
	fmt.Println(obj.SumRange(5, 6))
	obj.Update(9, 27)
	fmt.Println(obj.SumRange(2, 3))
	fmt.Println(obj.SumRange(6, 7))
	obj.Update(1, -82)
	obj.Update(3, -72)
	fmt.Println(obj.SumRange(3, 7))
	fmt.Println(obj.SumRange(1, 8))
	obj.Update(5, 13)
	obj.Update(4, -67)
}
