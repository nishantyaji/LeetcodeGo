package main

import (
	"container/heap"
	"fmt"
)

// Define a type for the max-heap.
// It needs to implement the heap.Interface methods.
type IntHeap []int

// Implement the Len() method required by heap.Interface
func (h IntHeap) Len() int {
	return len(h)
}

// Implement the Less() method required by heap.Interface
// This is a max-heap, so we want the larger number to come first.
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Implement the Swap() method required by heap.Interface
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Implement the Push() method required by heap.Interface
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Implement the Pop() method required by heap.Interface
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findScore(nums []int) int64 {
	var res int64

	var flags [1000000]int

	h := &IntHeap{}
	heap.Init(h) // Initialize the heap structure

	for i := 0; i < len(nums); i++ {
		heap.Push(h, 1000001*nums[i]+i)
	}

	count := 0
	for count < len(nums) {
		temp := heap.Pop(h)

		t, ok := temp.(int)
		if !ok {
			return -1
		}
		n := t / 1000001
		idx := t % 1000001
		if flags[idx] == 0 {
			res += (int64)(n)
			flags[idx] = 1
			count++

			if idx-1 >= 0 && flags[idx-1] == 0 {
				flags[idx-1] = 1
				count++
			}
			if idx+1 < len(nums) && flags[idx+1] == 0 {
				flags[idx+1] = 1
				count++
			}
		}
	}
	return res

}

func main() {
	var arr = []int{2, 1, 3, 4, 5, 2}
	fmt.Println(findScore(arr))  // 7
}
