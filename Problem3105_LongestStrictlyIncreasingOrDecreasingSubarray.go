package main

import "fmt"

func abs(numm int) int {
	if numm >= 0 {
		return numm
	}
	return -numm
}

func longestMonotonicSubarray(nums []int) int {
	prev := nums[0]
	cumu := 1
	res := -1

	for _, n := range nums {
		res = max(abs(cumu), res)
		if n > prev {
			if cumu <= 0 {
				cumu = 2
			} else {
				cumu += 1
			}
		} else if n < prev {
			if cumu >= 0 {
				cumu = -2
			} else {
				cumu -= 1
			}
		} else {
			cumu = 0
		}
		prev = n
	}
	return max(abs(cumu), res)
}

func main() {
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))
}
