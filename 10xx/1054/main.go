package main

import (
	"fmt"
)

/* 56ms */
func rearrangeBarcodes(barcodes []int) []int {
	count := make([]int, 10002)
	m := 0
	for _, v := range barcodes {
		count[v]++
		if count[v] > count[m] {
			m = v
		}
	}
	idx := 0
	for i := 0; i < count[m]; i++ {
		barcodes[idx] = m
		idx += 2
	}
	count[m] = 0
	for v, c := range count {
		for i := 0; i < c; i++ {
			if idx >= len(barcodes) {
				idx = 1
			}
			barcodes[idx] = v
			idx += 2
		}
	}
	return barcodes
}

func main() {
	var barcodes []int

	barcodes = []int{1, 1, 1, 2, 2, 2}
	fmt.Println(barcodes)
	fmt.Println(rearrangeBarcodes(barcodes))

	barcodes = []int{1, 1, 1, 1, 2, 2, 3, 3}
	fmt.Println(barcodes)
	fmt.Println(rearrangeBarcodes(barcodes))

	barcodes = []int{7, 7, 7, 8, 5, 7, 5, 5, 5, 8}
	fmt.Println(barcodes)
	fmt.Println(rearrangeBarcodes(barcodes))

}
