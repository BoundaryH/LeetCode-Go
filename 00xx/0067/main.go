package main

import "fmt"

/* 0ms */
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	c := make([]byte, len(a)+1)
	i, j, k := len(a)-1, len(b)-1, len(c)-1
	sum := byte(0)
	for j >= 0 {
		sum += a[i] - '0' + b[j] - '0'
		c[k] = (sum & 1) + '0'
		sum >>= 1
		i--
		j--
		k--
	}
	for i >= 0 {
		sum += a[i] - '0'
		c[k] = (sum & 1) + '0'
		sum >>= 1
		i--
		k--
	}
	if sum > 0 {
		c[k] = '1'
		k--
	}
	return string(c[k+1:])
}

func main() {
	var a, b string

	a, b = "11", "1"
	fmt.Println(a, b)
	fmt.Println(addBinary(a, b))

	a, b = "1010", "1011"
	fmt.Println(a, b)
	fmt.Println(addBinary(a, b))
}
