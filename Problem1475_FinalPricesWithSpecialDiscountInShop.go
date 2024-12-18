package main

import (
	"fmt"
)

type stack []int

func (s *stack) Push(i int) {
	*s = append(*s, i)
}

func (s *stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) Len() int {
	return len(*s)
}

func finalPrices(prices []int) []int {
	slice := make([]int, len(prices), len(prices))
	var st stack

	for i := len(prices) - 1; i >= 0; i-- {
		slice[i] = prices[i]
		if st.Len() > 0 {
			for st.Len() > 0 && st[st.Len()-1] > prices[i] {
				st.Pop()
			}
			if st.Len() > 0 {
				slice[i] = prices[i] - st[st.Len()-1]
			}
		}
		st.Push(prices[i])
	}
	return slice
}

func main() {
	s := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(s))
}
