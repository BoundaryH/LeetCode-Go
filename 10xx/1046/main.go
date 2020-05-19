package main

import (
	"container/heap"
	"fmt"
)

/* 0ms */
func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)
		heap.Push(&h, a-b)
	}
	return heap.Pop(&h).(int)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var stones []int

	stones = []int{2, 7, 4, 1, 8, 1}
	fmt.Println(stones)
	fmt.Println(lastStoneWeight(stones))
}
