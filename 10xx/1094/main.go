package main

import "fmt"

/* 4ms */
func carPooling(trips [][]int, capacity int) bool {
	start, end := make([]int, 1001), make([]int, 1001)
	for _, t := range trips {
		start[t[1]] += t[0]
		end[t[2]] += t[0]
	}
	for i := 0; i <= 1000; i++ {
		capacity += end[i] - start[i]
		if capacity < 0 {
			return false
		}
	}
	return true
}

func main() {
	var trips [][]int
	var capacity int

	trips = [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity = 4
	fmt.Println(trips)
	fmt.Println(capacity, carPooling(trips, capacity))

	trips = [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity = 5
	fmt.Println(trips)
	fmt.Println(capacity, carPooling(trips, capacity))

	trips = [][]int{{2, 1, 5}, {3, 5, 7}}
	capacity = 3
	fmt.Println(trips)
	fmt.Println(capacity, carPooling(trips, capacity))

	trips = [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}
	capacity = 11
	fmt.Println(trips)
	fmt.Println(capacity, carPooling(trips, capacity))

	trips = [][]int{{9, 3, 4}, {9, 1, 7}, {4, 2, 4}, {7, 4, 5}}
	capacity = 23
	fmt.Println(trips)
	fmt.Println(capacity, carPooling(trips, capacity))
}
