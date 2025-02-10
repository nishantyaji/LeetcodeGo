package main

import "fmt"

func countBadPairs(nums []int) int64 {
	d := make(map[int]int64)
	var res int64
	ln := int64(len(nums))

	for i, n := range nums {
		temp := n - i
		if _, ok := d[temp]; !ok {
			d[temp] = 1
		} else {
			d[temp] = d[temp] + 1
		}
	}
	for _, v := range d {
		res += ((ln - v) * v)
	}
	return res / 2
}

func main() {
	fmt.Println(countBadPairs([]int{4, 1, 3, 3})) // 5
}
