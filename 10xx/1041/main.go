package main

import "fmt"

/* 0ms */
func isRobotBounded(instructions string) bool {
	x, y, d := 0, 0, 0
	move := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, i := range instructions {
		if i == 'L' {
			d = (d + 3) % 4
		} else if i == 'R' {
			d = (d + 1) % 4
		} else {
			x += move[d][0]
			y += move[d][1]
		}
	}
	return (x == 0 && y == 0) || (d != 0)
}

func main() {
	var instructions string

	instructions = "GGLLGG"
	fmt.Println(isRobotBounded(instructions), instructions)

	instructions = "GG"
	fmt.Println(isRobotBounded(instructions), instructions)

	instructions = "GL"
	fmt.Println(isRobotBounded(instructions), instructions)
}
