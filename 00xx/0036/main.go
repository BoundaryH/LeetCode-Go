package main

import "fmt"

/* 0ms */
func isValidSudoku(board [][]byte) bool {
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
			if c := board[x][y]; c != '.' {
				c -= '1'
				b := x/3*3 + y/3
				if row[x][c] || col[y][c] || box[b][c] {
					return false
				}
				row[x][c] = true
				col[y][c] = true
				box[b][c] = true
			}
		}
	}
	return true
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
	fmt.Println(isValidSudoku(board))

	board = [][]byte{
		[]byte("83..7...."),
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
	fmt.Println(isValidSudoku(board))
}
