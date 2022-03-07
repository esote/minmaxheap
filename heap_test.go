package minmaxheap

import (
	"math/rand"
	"sort"
	"testing"
)

type myHeap []int

func (h myHeap) Len() int           { return len(h) }
func (h myHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h myHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *myHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h myHeap) verify(t *testing.T, i int) {
	t.Helper()
	n := h.Len()
	l := 2*i + 1
	r := 2*i + 2
	if l < n {
		if isMinLevel(i) {
			if h.Less(l, i) {
				t.Errorf("heap invariant violated [%d] = %d > [%d] = %d", i, h[i], l, h[l])
				return
			}
		} else {
			if h.Less(i, l) {
				t.Errorf("heap invariant violated [%d] = %d > [%d] = %d", l, h[l], i, h[i])
				return
			}
		}
		h.verify(t, l)
	}
	if r < n {
		if isMinLevel(i) {
			if h.Less(r, i) {
				t.Errorf("heap invariant violated [%d] = %d > [%d] = %d", i, h[i], r, h[r])
				return
			}
		} else {
			if h.Less(i, r) {
				t.Errorf("heap invariant violated [%d] = %d > [%d] = %d", r, h[r], i, h[i])
				return
			}
		}
		h.verify(t, r)
	}
}

func TestInit0(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(0) // all elements are the same
	}
	Init(h)
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := Pop(h).(int)
		h.verify(t, 0)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestInit0Max(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(0) // all elements are the same
	}
	Init(h)
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := PopMax(h).(int)
		h.verify(t, 0)
		if x != 0 {
			t.Errorf("%d.th popmax got %d; want %d", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(i) // all elements are different
	}
	Init(h)
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := Pop(h).(int)
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestInit1Max(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(i) // all elements are different
	}
	Init(h)
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x := PopMax(h).(int)
		h.verify(t, 0)
		if x != 20-i+1 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 20-i+1)
		}
	}
}
func TestInit2(t *testing.T) {
	testcases := []myHeap{
		{6, 10, 13, 3, 12, 8, 12, 2, 12, 16},
	}
	for _, tc := range testcases {
		Init(&tc)
		tc.verify(t, 0)
	}
}

func Test(t *testing.T) {
	h := new(myHeap)
	h.verify(t, 0)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}
	Init(h)
	h.verify(t, 0)

	for i := 10; i > 0; i-- {
		Push(h, i)
		h.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		x := Pop(h).(int)
		if i < 20 {
			Push(h, 20+i)
		}
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestMax(t *testing.T) {
	h := new(myHeap)
	h.verify(t, 0)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}
	Init(h)
	h.verify(t, 0)

	for i := 10; i > 0; i-- {
		Push(h, i)
		h.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		x := PopMax(h).(int)
		if i > 20 {
			Push(h, i-20)
		}
		h.verify(t, 0)
		if x != 20-i+1 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 20-i+1)
		}
	}
}

func TestRandomSorted(t *testing.T) {
	rand.Seed(0)

	const n = 1_000
	h := new(myHeap)
	for i := 0; i < n; i++ {
		*h = append(*h, rand.Intn(n/2))
	}

	Init(h)
	h.verify(t, 0)

	var ints []int
	for h.Len() > 0 {
		ints = append(ints, Pop(h).(int))
		h.verify(t, 0)
	}
	if !sort.IntsAreSorted(ints) {
		t.Fatal("min pop order invalid")
	}
}

func TestRandomSortedMax(t *testing.T) {
	rand.Seed(0)

	const n = 1_000
	h := new(myHeap)
	for i := 0; i < n; i++ {
		*h = append(*h, rand.Intn(n/2))
	}

	Init(h)
	h.verify(t, 0)

	var ints []int
	for h.Len() > 0 {
		ints = append(ints, PopMax(h).(int))
		h.verify(t, 0)
	}
	if !sort.IsSorted(sort.Reverse(sort.IntSlice(ints))) {
		t.Fatal("max pop order invalid")
	}
}
func TestRemove0(t *testing.T) {
	h := new(myHeap)
	for i := 0; i < 10; i++ {
		Push(h, i)
	}
	h.verify(t, 0)

	for h.Len() > 0 {
		i := h.Len() - 1
		want := (*h)[i]
		x := Remove(h, i).(int)
		if x != want {
			t.Errorf("Remove(%d) got %d; want %d", i, x, want)
		}
		h.verify(t, 0)
	}
}

func TestRemove1(t *testing.T) {
	h := new(myHeap)
	for i := 0; i < 10; i++ {
		Push(h, i)
	}
	h.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		x := Remove(h, 0).(int)
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		h.verify(t, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	h := new(myHeap)
	for i := 0; i < N; i++ {
		Push(h, i)
	}
	h.verify(t, 0)

	m := make(map[int]bool)
	for h.Len() > 0 {
		m[Remove(h, (h.Len()-1)/2).(int)] = true
		h.verify(t, 0)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	h := make(myHeap, 0, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			Push(&h, 0) // all elements are the same
		}
		for h.Len() > 0 {
			Pop(&h)
		}
	}
}

func TestFix(t *testing.T) {
	rand.Seed(0)

	h := new(myHeap)
	h.verify(t, 0)

	for i := 200; i > 0; i -= 10 {
		Push(h, i)
	}
	h.verify(t, 0)

	if (*h)[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", (*h)[0])
	}
	(*h)[0] = 210
	Fix(h, 0)
	h.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			(*h)[elem] *= 2
		} else {
			(*h)[elem] /= 2
		}
		Fix(h, elem)
		h.verify(t, 0)
	}
}
