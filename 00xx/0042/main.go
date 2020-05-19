package main

import "fmt"

/* 0ms */
func trap(height []int) (res int) {
	if len(height) == 0 {
		return
	}
	lMax := height[0]
	rMax := height[len(height)-1]

	l, r := 1, len(height)-2
	for l <= r {
		if lMax < rMax {
			h := height[l]
			if h > lMax {
				lMax = h
			}
			res += lMax - h
			l++
		} else {
			h := height[r]
			if h > rMax {
				rMax = h
			}
			res += rMax - h
			r--
		}
	}
	return
}

/* 4ms */
/*
func trap(height []int) (res int) {
	if len(height) == 0 {
		return
	}
	lMax := make([]int, len(height))

	m := height[0]
	for i, h := range height {
		if h > m {
			m = h
		}
		lMax[i] = m
	}

	m = height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > m {
			m = height[i]
		}
		if m < lMax[i] {
			res += m - height[i]
		} else {
			res += lMax[i] - height[i]
		}
	}
	return
}
*/

func main() {
	var height []int

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(height)
	fmt.Println(trap(height))
}
