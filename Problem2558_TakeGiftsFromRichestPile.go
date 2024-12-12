import (
	"container/heap"
	"fmt"
	"math"
)

// Define a type for the max-heap. heap.Interface methods.
type IntHeap []int

// Implement the Len() method required by heap.Interface
func (h IntHeap) Len() int {
	return len(h)
}

// Implement the Less() method required by heap.Interface. Max Heap (h[i] > h[j])
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
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

func floor_sq(n int) int {
	sqrtValue := math.Sqrt(float64(n))
	floorValue := math.Floor(sqrtValue)
	new_n := int(floorValue)
	return new_n
}

func pickGifts(gifts []int, k int) int64 {
	var res int64

	h := &IntHeap{}
	*h = append(*h, gifts...) // Copy the contents of gifts into the heap
	heap.Init(h)              // Initialize the heap structure

	for i := 0; i < k; i++ {
		popped := heap.Pop(h)

		g, ok := popped.(int)
		if !ok {
			return -1
		}
		if g == 1 {
			return int64(len(gifts))
		}
		new_g := floor_sq(g)
		heap.Push(h, new_g)
	}

	for h.Len() > 0 {
		popped := heap.Pop(h)
		g, ok := popped.(int)
		if !ok {
			return -1
		}
		fmt.Println(g)
		res += int64(g)
	}
	return res
}

func main() {
  var gifts = []int{25, 64, 9, 4, 100}
  k := 4
  fmt.Println(pickGifts(gifts, k))
}
