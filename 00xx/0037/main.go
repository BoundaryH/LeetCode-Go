package main

import "fmt"

/* 0ms */
func solveSudoku(board [][]byte) {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	box := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		box[i] = make([]bool, 9)
	}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if n := board[x][y]; n != '.' {
				n = n - '1'
				row[x][n] = true
				col[y][n] = true
				box[x/3*3+y/3][n] = true
			}
		}
	}

	var search func(x, y int) bool
	search = func(x, y int) bool {
		for x < 9 && board[x][y] != '.' {
			y++
			if y == 9 {
				x++
				y = 0
			}
		}
		if x == 9 {
			return true
		}
		b := x/3*3 + y/3
		for i := 0; i < 9; i++ {
			if !row[x][i] && !col[y][i] && !box[b][i] {
				board[x][y] = byte(i + '1')
				row[x][i] = true
				col[y][i] = true
				box[b][i] = true
				if search(x, y) {
					return true
				}
				board[x][y] = '.'
				row[x][i] = false
				col[y][i] = false
				box[b][i] = false
			}
		}
		return false
	}
	search(0, 0)
}

func main() {
	var board [][]byte

	board = [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	for _, r := range board {
		fmt.Println(string(r))
	}
	solveSudoku(board)
	fmt.Println()
	for _, r := range board {
		fmt.Println(string(r))
	}

}
