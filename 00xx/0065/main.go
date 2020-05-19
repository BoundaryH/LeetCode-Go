package main

import (
	"fmt"
	"strings"
)

/* 0ms */
func isNumber(s string) bool {
	isNum := func(s string) bool {
		for _, c := range s {
			if c > '9' || c < '0' {
				return false
			}
		}
		return len(s) != 0
	}
	dropSign := func(s string) string {
		if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
			return s[1:]
		}
		return s
	}

	s = strings.TrimSpace(s)
	if p := strings.IndexByte(s, 'e'); p != -1 {
		if a := dropSign(s[p+1:]); !isNum(a) {
			return false
		}
		s = s[:p]
	}
	s = dropSign(s)
	if p := strings.IndexByte(s, '.'); p != -1 {
		if len(s) == 1 {
			return false
		}
		if a := s[p+1:]; len(a) > 0 && !isNum(a) {
			return false
		}
		s = s[:p]
		return len(s) == 0 || isNum(s)
	}
	return isNum(s)
}

func main() {

	strs := []string{
		"+.8",
		"0",
		" 0.1 ",
		"abc",
		"1 a",
		"2e10",
		"-8.0",
		"",
		"   ",
		"e10",
		"4e",
		"2.",
		".8",
		"--1",
		"001",
		".",
		"46.e3",
		" 005047e+6",
		" 005047e+",
		" 005047e-7",
		"-.",
	}
	for _, s := range strs {
		fmt.Println(isNumber(s), "  ", s)
	}
	//fmt.Println(isNumber(strs[0]))
}
