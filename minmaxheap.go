// Package minmaxheap provides min-max heap operations for any type that
// implements heap.Interface. A min-max heap, also called a double-ended
// priority queue, has logarithmic-time removal of minimum and maximum elements.
//
// Min-max heap implementation from the 1986 paper "Min-Max Heaps and
// Generalized Priority Queues" by Atkinson et. al.
// https://doi.org/10.1145/6617.6621.
package minmaxheap

// TODO: implement Remove from "container/heap".
// TODO: implement Fix from "container/heap".
// TODO: handle overflow in rchild, lchild, level.

import (
	"container/heap"
	"math/bits"
)

// Interface from the heap package, so that code that imports minmaxheap does
// not have to import "container/heap".
type Interface = heap.Interface

func level(i int) int {
	return bits.Len(uint(i)+1) - 1
}

func isMinLevel(i int) bool {
	return level(i)%2 == 0
}

func parent(i int) int {
	return (i - 1) / 2
}

func hasParent(i int) bool {
	return i > 0
}

func hasGrandparent(i int) bool {
	return i > 2
}

func hasChildren(i, n int) bool {
	return lchild(i) < n
}

func lchild(i int) int {
	return i*2 + 1
}

func rchild(i int) int {
	return i*2 + 2
}

func down(h Interface, i, n int) {
	downMinMax(h, i, n, isMinLevel(i))
}

func downMinMax(h Interface, m, n int, min bool) {
	for hasChildren(m, n) {
		child := true
		i := m
		// min of children
		m = lchild(i)
		if rchild(i) < n && h.Less(rchild(i), m) == min {
			m = rchild(i)
		}
		// min of grandchildren
		// grandchildren are contiguous i*4+{3,4,5,6}
		for j := lchild(lchild(i)); j < n && j <= rchild(rchild(i)); j++ {
			if h.Less(j, m) == min {
				m = j
				child = false
			}
		}
		if h.Less(m, i) != min {
			return
		}
		h.Swap(m, i)
		if child {
			return
		}
		if h.Less(parent(m), m) == min {
			h.Swap(m, parent(m))
		}
	}
}

func up(h Interface, i int) {
	min := isMinLevel(i)
	if hasParent(i) && h.Less(i, parent(i)) != min {
		h.Swap(i, parent(i))
		upMinMax(h, parent(i), !min)
	} else {
		upMinMax(h, i, min)
	}
}

func upMinMax(h Interface, i int, min bool) {
	for hasGrandparent(i) && h.Less(i, parent(parent(i))) == min {
		h.Swap(i, parent(parent(i)))
		i = parent(parent(i))
	}
}

// Init establishes heap ordering. The complexity is O(7n/3) = O(n) where n =
// h.Len().
func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap. The complexity is O(0.5log(n+1)) =
// O(log(n)) where n = h.Len().
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

// PopMin removes and returns the minimum element from the heap. The complexity
// is O(2.5log(n)) = O(log(n)) where n = h.Len().
func PopMin(h Interface) interface{} {
	// Minimum element is at h[0]
	n := h.Len()
	h.Swap(0, n-1)
	down(h, 0, n-1)
	return h.Pop()
}

// PopMax removes and returns the maximum element from the heap. The complexity
// is O(2.5log(n)) = O(log(n)) where n = h.Len().
func PopMax(h Interface) interface{} {
	n := h.Len()
	if n <= 2 {
		// n=1: pop root h[0]
		// n=2: child of root is smallest, pop h[lchild(0)] = h[1]
		return h.Pop()
	}
	i := lchild(0)
	if !h.Less(rchild(0), i) {
		i = rchild(0)
	}
	h.Swap(i, n-1)
	down(h, i, n-1)
	return h.Pop()
}
