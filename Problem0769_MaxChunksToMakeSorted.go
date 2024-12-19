package main

import "fmt"

func maxChunksToSorted(arr []int) int {
	prefixMax := make([]int, len(arr), len(arr)) // slice
	suffixMin := make([]int, len(arr), len(arr)) // slice

	mx := -1
	mn := 1000001

	for i := 0; i < len(arr); i++ {
		mx = max(mx, arr[i])
		mn = min(mn, arr[len(arr)-1-i])
		prefixMax[i] = mx
		suffixMin[len(arr)-1-i] = mn
	}

	chunk := 0
	for i := 0; i < len(arr); i++ {
		if i == 0 && arr[i] == 0 {
			chunk++
		} else if i > 0 && i < len(arr)-1 && prefixMax[i] <= i && suffixMin[i+1] > i {
			chunk++
		} else if i == len(arr)-1 && prefixMax[i] <= i {
			chunk++
		}
	}
	return chunk
}

func main() {
	var arr = []int{1, 0, 2, 3, 4}
	fmt.Println(maxChunksToSorted(arr))
}
