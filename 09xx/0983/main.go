package main

import "fmt"

/* 0ms */
func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)

	dayCost, weekCost, monthCost := costs[0], costs[1], costs[2]
	for i, j := days[0], 0; i <= lastDay; i++ {
		if i != days[j] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = dp[i-1] + dayCost
		if i < 7 {
			dp[i] = min(dp[i], weekCost)
			dp[i] = min(dp[i], monthCost)
		} else if i < 30 {
			dp[i] = min(dp[i], dp[i-7]+weekCost)
			dp[i] = min(dp[i], monthCost)
		} else {
			dp[i] = min(dp[i], dp[i-7]+weekCost)
			dp[i] = min(dp[i], dp[i-30]+monthCost)
		}
		j++
	}
	return dp[lastDay]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	var days, costs []int

	days = []int{1, 4, 6, 7, 8, 20}
	costs = []int{2, 7, 15}
	fmt.Println(days)
	fmt.Println(costs, mincostTickets(days, costs))

	days = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs = []int{2, 7, 15}
	fmt.Println(days)
	fmt.Println(costs, mincostTickets(days, costs))
}
