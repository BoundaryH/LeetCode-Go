package main

import (
	"fmt"
	"strconv"
)

/* 0ms */
func restoreIpAddresses(s string) (res []string) {
	var restore func(s string, ip []byte, size int)
	restore = func(s string, ip []byte, size int) {
		if size == 4 && len(s) == 0 {
			res = append(res, string(ip[1:]))
			return
		} else if size >= 4 || len(s) == 0 {
			return
		}
		ip = append(ip, '.')
		restore(s[1:], append(ip, s[0]), size+1)
		if s[0] != '0' {
			if len(s) >= 2 {
				restore(s[2:], append(ip, s[:2]...), size+1)
			}
			if len(s) >= 3 {
				if n, err := strconv.Atoi(s[:3]); err == nil && n <= 255 {
					restore(s[3:], append(ip, s[:3]...), size+1)
				}
			}
		}
	}
	restore(s, make([]byte, 0, len(s)+4), 0)
	return
}

func main() {
	var s string

	s = "25525511135"
	fmt.Println(s)
	fmt.Println(restoreIpAddresses(s))

	s = "010010"
	fmt.Println(s)
	fmt.Println(restoreIpAddresses(s))
}
