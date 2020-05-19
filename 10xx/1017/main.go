package main

import "fmt"

/* 0ms */
func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	bin := make([]byte, 40)
	i := len(bin) - 1
	j := 1
	for N != 0 {
		if N&1 == 1 {
			bin[i] = '1'
			N -= j
		} else {
			bin[i] = '0'
		}
		N >>= 1
		j *= -1
		i--
	}
	return string(bin[i+1:])
}

func main() {
	var N int

	for N = 2; N < 20; N++ {
		fmt.Println(N, baseNeg2(N))
	}
}
