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

func xorAllNums(nums1 []int, nums2 []int) int {
	x1 := reduce(nums1, xor, 0)
	x2 := reduce(nums2, xor, 0)
	x1temp := 0
	x2temp := 0
	if len(nums2)%2 == 1 {
		x1temp = x1
	}
	if len(nums1)%2 == 1 {
		x2temp = x2
	}
	return x1temp ^ x2temp
}

func main() {

	nums1 := []int{2, 1, 3}
	nums2 := []int{10, 2, 5, 0}
	fmt.Println(xorAllNums(nums1, nums2))
}
