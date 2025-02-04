package main

import "fmt"

func maxAscendingSum(nums []int) int {
	prev := -1
	res := -1
	cumu := 0

	for _, n := range nums {
		res = max(cumu, res)
		if n > prev {
			cumu += n
		} else {
			cumu = n
		}
		prev = n
	}
	return max(cumu, res)

}

func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}
