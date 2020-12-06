package minmaxheap

import "testing"

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h intHeap) verify(t *testing.T) {
	n := len(h)
	for i := range h {
		min := isMinLevel(i)
		// children should larger (min) or smaller (!min)
		for j := lchild(i); j < n && j <= rchild(i); j++ {
			if h.Less(j, i) == min {
				t.Fatalf("heap invariant invalidated [%d]%d > [%d]%d == %t", i, h[i], j, h[j], min)
				return
			}
		}
		// grandchildren should be larger (min) or smaller (!min)
		for j := lchild(lchild(i)); j < n && j <= rchild(rchild(i)); j++ {
			if h.Less(j, i) == min {
				t.Fatalf("heap invariant invalidated [%d]%d > [%d]%d == %t", i, h[i], j, h[j], min)
				return
			}
		}
	}
}

func Test(t *testing.T) {
	// Test adapted from heap_test.go in "container/heap".
	h := new(intHeap)
	h.verify(t)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}
	Init(h)
	h.verify(t)

	for i := 10; i > 0; i-- {
		Push(h, i)
		h.verify(t)
	}

	for i := 1; h.Len() > 0; i++ {
		x := PopMin(h).(int)
		if i < 20 {
			Push(h, 20+i)
		}
		h.verify(t)
		if x != i {
			t.Fatalf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestMin(t *testing.T) {
	h := new(intHeap)
	for i := 0; i <= 20; i++ {
		h.Push(i)
	}
	Init(h)

	for i := 0; i <= 20; i++ {
		if x := PopMin(h).(int); x != i {
			t.Fatalf("%d.th pop got %d; want %d", i, x, i)
		}
		h.verify(t)
	}
}

func TestMax(t *testing.T) {
	h := new(intHeap)
	for i := 0; i <= 20; i++ {
		h.Push(i)
	}
	Init(h)

	for i := 20; i >= 0; i-- {
		if x := PopMax(h).(int); x != i {
			t.Fatalf("%d.th pop got %d; want %d", i, x, i)
		}
		h.verify(t)
	}
}
