package main

import (
	"container/heap"
	"fmt"
)

// Define a struct to represent a pair (tuple)
type Pair struct {
	Amount int
	Index  int
}

// Define a type for the max-heap.
// It needs to implement the heap.Interface methods.
type MinHeap []Pair

// Implement the Len() method required by heap.Interface
func (h MinHeap) Len() int {
	return len(h)
}

// Implement the Less() method required by heap.Interface
// This is a max-heap, so we want the larger number to come first.
func (h MinHeap) Less(i, j int) bool {
	if h[i].Amount == h[j].Amount {
		return h[i].Index < h[j].Index
	}
	return h[i].Amount < h[j].Amount
}

// Implement the Swap() method required by heap.Interface
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Implement the Push() method required by heap.Interface
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

// Implement the Pop() method required by heap.Interface
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	h := &MinHeap{}
	heap.Init(h)

	for i, x := range nums {
		p := Pair{Amount: x, Index: i}
		heap.Push(h, p)
	}

	for i := 0; i < k; i++ {
		temp := heap.Pop(h)
		t, ok := temp.(Pair)
		if !ok {
			return nums
		}
		t2 := Pair{Amount: t.Amount * multiplier, Index: t.Index}
		heap.Push(h, t2)
	}

	slice := make([]int, len(nums), len(nums))
	for h.Len() > 0 {
		popped := heap.Pop(h)
		g, ok := popped.(Pair)
		if !ok {
			return nums
		}
		slice[g.Index] = g.Amount
	}
	return slice
}

func main() {
	var arr = []int{2, 1, 3, 5, 6}

	fmt.Println(getFinalState(arr, 5, 2))
}
