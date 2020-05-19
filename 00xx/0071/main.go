package main

import (
	"bytes"
	"fmt"
	"strings"
)

/* 0ms */
func simplifyPath(path string) string {
	stk := make([]string, 0)
	for _, s := range strings.Split(path, "/") {
		if s == ".." {
			if l := len(stk); l > 0 {
				stk = stk[:l-1]
			}
		} else if len(s) > 0 && s != "." {
			stk = append(stk, s)
		}
	}
	if len(stk) == 0 {
		return "/"
	}
	var res bytes.Buffer
	for _, s := range stk {
		res.WriteByte('/')
		res.WriteString(s)
	}
	return res.String()
}

func main() {
	ps := []string{
		"/home/",
		"/../",
		"/home//foo/",
		"/a/./b/../../c/",
		"/a/../../b/../c//.//",
		"/a//b////c/d//././/..",
	}

	for _, path := range ps {
		fmt.Println(path)
		fmt.Println(simplifyPath(path))
	}
}
