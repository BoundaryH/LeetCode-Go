package main

import "fmt"

/* 0ms */
func getPermutation(n int, k int) string {
	f := make([]int, n+1)
	c := make([]byte, n)
	f[0] = 1
	for i := 1; i < len(f); i++ {
		f[i] = f[i-1] * i
		c[i-1] = byte(i)
	}

	s := make([]byte, n)
	k--
	i, j := n-1, 0
	for i >= 0 {
		m := k / f[i]
		k -= m * f[i]
		s[j] = c[m] + '0'
		c = append(c[:m], c[m+1:]...)
		i--
		j++
	}
	return string(s)
}

func main() {
	var n, k int

	n, k = 3, 1
	fmt.Println(n, k)
	fmt.Println(getPermutation(n, k))

	n, k = 4, 9
	fmt.Println(n, k)
	fmt.Println(getPermutation(n, k))
}
