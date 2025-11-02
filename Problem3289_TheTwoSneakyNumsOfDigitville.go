// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	test := []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}
	res := getSneakyNumbers(test)
	for _, r := range res {
		fmt.Println(r)
	}
}

func getSneakyNumbers(nums []int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; ok {
			mp[nums[i]] = mp[nums[i]] + 1
		} else {
			mp[nums[i]] = 1
		}
	}
	res := make([]int, 2, 2)
	j := 0
	for key, value := range mp {
		if value > 1 {
			res[j] = key
			value += 1
			j += 1
		}
	}
	return res
}
