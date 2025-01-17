package main

import "fmt"

func xor(a int, b int) int {
	return a ^ b
}

func reduce(nums []int, op func(int, int) int, initVal int) int {
	result := initVal
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			result = op(result, nums[i])
		}
	}
	return result
}

func doesValidArrayExist(derived []int) bool {
	return reduce(derived, xor, 0) == 0
}

func main() {
	fmt.Println(doesValidArrayExist([]int{1, 2, 3}))
}
