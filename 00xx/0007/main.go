package main

import "fmt"

/* 0ms */
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	if y > 0x7fffffff || y < 0x80000000*-1 {
		return 0
	}
	return y
}

func main() {
	x := 0
	fmt.Printf("%12d  ->  %12d\n", x, reverse(x))
	x = 123
	fmt.Printf("%12d  ->  %12d\n", x, reverse(x))
	x = -1563847412
	fmt.Printf("%12d  ->  %12d\n", x, reverse(x))
	x = 0x80000000
	fmt.Printf("%12d  ->  %12d\n", x, reverse(x))
	x = -2147483412
	fmt.Printf("%12d  ->  %12d\n", x, reverse(x))
	x = 483847412
	fmt.Printf("%12d  ->  %12d\n", x, reverse(x))
	fmt.Println((1 << 31) - 1)
	fmt.Println(0x7fffffff)
	fmt.Println(0x80000000 * -1)
}
