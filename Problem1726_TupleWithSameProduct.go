package main

import "fmt"

func tupleSameProduct(nums []int) int {
	cntr := make(map[int]int)
	lenn := len(nums)
	for i := 0; i < lenn; i++ {
		for j := i + 1; j < lenn; j++ {
			prod := nums[i] * nums[j]
			if val, ok := cntr[prod]; ok {
				cntr[prod] = val + 1
			} else {
				cntr[prod] = 1
			}
		}
	}
	res := 0
	for _, v := range cntr {
		if v > 1 {
			res += (4 * v * (v - 1))
		}
	}
	return res
}

func main() {
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10}))  // 16
}
