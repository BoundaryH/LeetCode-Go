package main

import "fmt"

/* 0ms */
func countAndSay(n int) string {
	a := []byte{'1'}
	b := []byte{}
	for i := 2; i <= n; i++ {
		a = append(a, '#')
		j := 0
		k := 0
		for ; k < len(a); k++ {
			if a[j] != a[k] {
				b = append(b, byte(k-j+'0'))
				b = append(b, a[j])
				j = k
			}
		}
		a = b
		b = []byte{}
	}
	return string(a)
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(countAndSay(i))
	}
}
