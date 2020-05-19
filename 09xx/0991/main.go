package main

import "fmt"

/* 0ms */
func brokenCalc(X int, Y int) int {
	count := 0
	for X != Y {
		if Y < X {
			count += X - Y
			Y = X
		} else if Y%2 == 0 {
			count++
			Y /= 2
		} else {
			count += 2
			Y = (Y + 1) / 2
		}
	}
	return count
}

func main() {
	var X, Y int

	X, Y = 2, 3
	fmt.Println(X, Y, brokenCalc(X, Y))

	X, Y = 5, 8
	fmt.Println(X, Y, brokenCalc(X, Y))

	X, Y = 3, 10
	fmt.Println(X, Y, brokenCalc(X, Y))

	X, Y = 1024, 1
	fmt.Println(X, Y, brokenCalc(X, Y))
}
