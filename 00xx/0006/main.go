package main

import "fmt"

/* 0ms */
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rows := make([][]byte, numRows)
	n, m := 0, numRows*2-2
	for _, c := range s {
		if n < numRows {
			rows[n] = append(rows[n], byte(c))
		} else {
			rows[m-n] = append(rows[m-n], byte(c))
		}
		n++
		if n == m {
			n = 0
		}
	}
	res := make([]byte, 0, len(s))
	for _, r := range rows {
		res = append(res, r...)
	}
	return string(res)
}

func main() {
	var s string
	var numRows int

	s = "LEETCODEISHIRING"
	numRows = 4
	fmt.Println(numRows, s)
	fmt.Println(convert(s, numRows))

	s = "PAYPALISHIRING"
	numRows = 3
	fmt.Println(numRows, s)
	fmt.Println(convert(s, numRows))

	s = "PAYPALISHIRING"
	numRows = 4
	fmt.Println(numRows, s)
	fmt.Println(convert(s, numRows))
}
