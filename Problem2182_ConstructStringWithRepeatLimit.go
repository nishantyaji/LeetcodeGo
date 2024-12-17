package main

import (
	"fmt"
	"strings"
)

type CounterPair struct {
	Character rune
	Count     int
}

type stack []CounterPair

func (s *stack) Push(cp CounterPair) {
	*s = append(*s, cp)
}

func (s *stack) Pop() CounterPair {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) Len() int {
	return len(*s)
}

func repeatLimitedString(s string, repeatLimit int) string {
	aInt := int('a')
	slice := make([]CounterPair, 26, 26)
	for _, c := range s {
		index := int(c) - aInt
		slice[index].Character = c
		slice[index].Count += 1
	}
	var st stack
	for _, cp := range slice {
		if cp.Count > 0 {
			st.Push(cp)
		}
	}

	var builder strings.Builder // Create a new builder
	// Use string builder because usual string appending (s += "a") is not efficient
	count := 0

	for i := 0; i < len(s); i++ {
		if count == repeatLimit {
			large := st.Pop()
			if st.Len() > 0 {
				seco := st.Pop()
				builder.WriteByte('a' + byte(int(seco.Character)-aInt))
				seco.Count -= 1
				if seco.Count > 0 {
					st.Push(seco)
				}
				count = 0
			} else {
				break
			}
			st.Push(large)
		} else {
			large := st.Pop()
			builder.WriteByte('a' + byte(int(large.Character)-aInt))
			large.Count -= 1
			if large.Count > 0 {
				st.Push(large)
				count++
			} else {
				count = 0
			}
		}
	}
	return builder.String()
}

func main() {

	fmt.Println(repeatLimitedString("cczazcc", 3))
	fmt.Println(repeatLimitedString("aababab", 2))
	fmt.Println(repeatLimitedString("ccccccccc", 3))
}
