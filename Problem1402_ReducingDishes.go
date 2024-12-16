package main

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	res := 0
	runSum := 0
	cumSum := 0
	for _, x := range satisfaction {
		cumSum += x
		runSum += cumSum
		res = max(res, runSum)

	}
	return res
}

func main() {
	s := []int{-1, -8, 0, 5, -7} // unsorted
	fmt.Println(maxSatisfaction(s))
}
