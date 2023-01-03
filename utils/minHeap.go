package utils

import "container/heap"

// MinHeap is implementation of binary min-heap
// see https://maupanelo.com/posts/how-to-write-a-binary-heap-in-golang/
type MinHeap[T any, N Number] struct {
	adapter *minHeapAdapter[T, N]
}

func (h MinHeap[T, N]) Push(item T, value N) {
	hi := heapItem[T, N]{item: item, value: value}
	heap.Push(h.adapter, hi)
}

func (h MinHeap[T, N]) Pop() T {
	item, _ := h.PopWithValue()
	return item
}

func (h MinHeap[T, N]) PopWithValue() (T, N) {
	item := heap.Pop(h.adapter).(heapItem[T, N])
	return item.item, item.value
}

func (h MinHeap[T, N]) Len() int {
	return h.adapter.Len()
}

func (h MinHeap[T, N]) Empty() bool {
	return h.adapter.Len() == 0
}

func NewMinHeap[T any, N Number]() MinHeap[T, N] {
	return MinHeap[T, N]{adapter: nil}
}

func NewMinHeapInt[T any]() MinHeap[T, int] {
	return MinHeap[T, int]{adapter: &minHeapAdapter[T, int]{}}
}

type heapItem[T any, N Number] struct {
	item  T
	value N
}

type minHeapAdapter[T any, N Number] []heapItem[T, N]

func (h minHeapAdapter[T, N]) Len() int           { return len(h) }
func (h minHeapAdapter[T, N]) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minHeapAdapter[T, N]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeapAdapter[T, N]) Push(item any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, item.(heapItem[T, N]))
}

func (h *minHeapAdapter[T, N]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
