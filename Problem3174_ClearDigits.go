package main

import (
	"fmt"
	"unicode"
)

type stack []rune

func (s *stack) Push(i rune) {
	*s = append(*s, i)
}

func (s *stack) Pop() rune {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) Len() int {
	return len(*s)
}

func clearDigits(s string) string {
	var st stack
	for _, c := range s {
		if unicode.IsDigit(c) {
			if st.Len() > 0 {
				st.Pop()
			}
		} else {
			st.Push(c)
		}

	}
	return string(st)
}

func main() {
	fmt.Println(clearDigits("abc12d"))
}
