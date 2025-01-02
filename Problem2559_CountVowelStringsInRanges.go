package main

import "fmt"

func isVowel(s string) int {
	mp := make(map[rune]int)
	mp['a'] = 1
	mp['e'] = 1
	mp['i'] = 1
	mp['o'] = 1
	mp['u'] = 1
	if _, ok1 := mp[rune(s[0])]; ok1 {
		if _, ok2 := mp[rune(s[len(s)-1])]; ok2 {
			return 1
		}
	}
	return 0
}

func mapSlice(ss []string, fn func(string) int) []int {
	result := make([]int, len(ss), len(ss))
	for i, s := range ss {
		result[i] = fn(s)
	}
	return result
}

func add(a int, b int) int {
	return a + b
}

func accumulateWithInitVal(nums []int, op func(int, int) int, initVal int) []int {
	result := make([]int, len(nums)+1, len(nums)+1)
	if len(nums) > 0 {
		result[0] = initVal
		for i := 1; i <= len(nums); i++ {
			result[i] = op(result[i-1], nums[i-1])
		}
	}
	return result
}

func accumulateWithoutInitVal(nums []int, op func(int, int) int) []int {
	result := make([]int, len(nums))
	if len(nums) > 0 {
		result[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			result[i] = op(result[i-1], nums[i])
		}
	}
	return result
}

func vowelStrings(words []string, queries [][]int) []int {
	acc := accumulateWithInitVal(mapSlice(words, isVowel), add, 0)
	res := make([]int, len(queries), len(queries))
	for i, q := range queries {
		res[i] = acc[q[1]+1] - acc[q[0]]
	}
	return res
}

func main() {
	queries := [][]int{
		{0, 2},
		{1, 4},
		{1, 1},
	}
	test := []string{"aba", "bcb", "ece", "aa", "e"}
	res := vowelStrings(test, queries)
	for _, r := range res {
		fmt.Println(r)
	}
}
