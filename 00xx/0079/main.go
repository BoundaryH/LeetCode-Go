package main

import "fmt"

/* 4ms */
func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	mark := make([][]bool, row)
	for i := range mark {
		mark[i] = make([]bool, col)
	}

	var find func(x, y, i int) bool
	find = func(x, y, i int) bool {
		if i == len(word) {
			return true
		}
		if x < 0 || y < 0 || x >= row || y >= col ||
			mark[x][y] || board[x][y] != word[i] {
			return false
		}
		mark[x][y] = true
		i++
		res := find(x+1, y, i) || find(x-1, y, i) ||
			find(x, y+1, i) || find(x, y-1, i)
		mark[x][y] = false
		return res
	}

	for x, r := range board {
		for y := range r {
			if find(x, y, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	var board [][]byte
	var word string

	board = [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	word = "ABCCED"
	for _, i := range board {
		fmt.Println(string(i))
	}
	fmt.Println(word, exist(board, word))

	word = "ABCB"
	fmt.Println(word, exist(board, word))

	word = "SEE"
	fmt.Println(word, exist(board, word))
}
