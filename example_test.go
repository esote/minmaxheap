package minmaxheap_test

import (
	"fmt"

	heap "github.com/esote/minmaxheap"
)

// IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Example() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)

	fmt.Println("min:", heap.Pop(h))
	fmt.Println("max:", heap.PopMax(h))
	// Output:
	// min: 1
	// max: 5
}
