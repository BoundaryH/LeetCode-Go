package main

import "fmt"

/* 0ms */
func videoStitching(clips [][]int, T int) int {
	clipMap := make([]int, T+1)
	for _, c := range clips {
		start, end := c[0], c[1]
		if start > T {
			start = T
		}
		if end > T {
			end = T
		}
		if clipMap[start] < end {
			clipMap[start] = end
		}
	}
	dp := make([]int, T+1)
	for i := 0; i <= T; i++ {
		dp[i] = T + 1
	}
	dp[0] = 0
	for i := 0; i <= T; i++ {
		sum := dp[i] + 1
		j := clipMap[i]
		if j == 0 {
			continue
		}
		for ; j > i && dp[j] > sum; j-- {
			dp[j] = sum
		}
	}

	if dp[T] == T+1 {
		return -1
	}
	return dp[T]
}

func main() {
	var clips [][]int
	var T int

	clips = [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	T = 10
	fmt.Println(clips)
	fmt.Println(T, videoStitching(clips, T))

	clips = [][]int{{0, 1}, {1, 2}}
	T = 5
	fmt.Println(clips)
	fmt.Println(T, videoStitching(clips, T))

	clips = [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	T = 9
	fmt.Println(clips)
	fmt.Println(T, videoStitching(clips, T))

	clips = [][]int{{0, 4}, {2, 8}}
	T = 5
	fmt.Println(clips)
	fmt.Println(T, videoStitching(clips, T))
}
