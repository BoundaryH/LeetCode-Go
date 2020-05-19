package main

import (
	"bytes"
	"fmt"
)

/* 0ms */
func fullJustify(words []string, maxWidth int) (res []string) {
	space := make([]byte, maxWidth)
	for i := range space {
		space[i] = ' '
	}
	concate := func(words []string, width int) string {
		var line bytes.Buffer
		line.Grow(maxWidth)
		line.WriteString(words[0])

		if len(words) == 1 {
			line.Write(space[:maxWidth-len(words[0])])
			return line.String()
		}

		a := (maxWidth-width)/(len(words)-1) + 1
		b := (maxWidth-width)%(len(words)-1) + 1
		i := 1
		for ; i < b; i++ {
			line.Write(space[:a+1])
			line.WriteString(words[i])
		}
		for ; i < len(words); i++ {
			line.Write(space[:a])
			line.WriteString(words[i])
		}
		return line.String()
	}

	row := make([]string, 0)
	rw := -1
	for _, w := range words {
		if rw+len(w)+1 > maxWidth {
			res = append(res, concate(row, rw))
			row = nil
			rw = -1
		}
		row = append(row, w)
		rw += len(w) + 1
	}
	if len(row) > 0 {
		line := bytes.NewBufferString(row[0])
		line.Grow(maxWidth)
		for i := 1; i < len(row); i++ {
			line.WriteByte(' ')
			line.WriteString(row[i])
		}
		line.Write(space[:maxWidth-line.Len()])
		res = append(res, line.String())
	}
	return res
}

func main() {
	var words []string
	var maxWidth int

	maxWidth = 16
	words = []string{
		"This", "is", "an", "example", "of", "text", "justification.",
	}
	fmt.Println(words)
	fmt.Println(maxWidth)
	for _, s := range fullJustify(words, maxWidth) {
		fmt.Println(s, "|", len(s))
	}
	fmt.Println()

	maxWidth = 16
	words = []string{
		"What", "must", "be", "acknowledgment", "shall", "be",
	}
	fmt.Println(words)
	fmt.Println(maxWidth)
	for _, s := range fullJustify(words, maxWidth) {
		fmt.Println(s, "|", len(s))
	}
	fmt.Println()

	maxWidth = 20
	words = []string{
		"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do",
	}
	fmt.Println(words)
	fmt.Println(maxWidth)
	for _, s := range fullJustify(words, maxWidth) {
		fmt.Println(s, "|", len(s))
	}
	fmt.Println()

}
