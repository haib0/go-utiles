package libs

import (
	"container/heap"
	"sort"
)

type MySlice []int

type hp struct {
	sort.IntSlice
	MySlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func RunHeap() {
	h := &hp{}
	heap.Push(h, 0)
}
