package main

import "fmt"

/* 4ms */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	k := 1
	for ; x > k*10; k *= 10 {
	}
	for x > 0 && k != 0 && x%10 == x/k {
		x = (x % k) / 10
		k /= 100
	}
	return x == 0
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
}
