package collections

import (
	"container/heap"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
)

// MinHeap is implementation of binary min-heap
// see https://maupanelo.com/posts/how-to-write-a-binary-heap-in-golang/
type MinHeap[T comparable, N utils.Number] struct {
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
	item := heap.Pop(h.adapter).(*heapItem[T, N])
	return item.item, item.value
}

func (h MinHeap[T, N]) Fix(item T, value N) {
	// lookup heap item for it's index
	hi := h.adapter.item2heapItem[item]
	hi.value = value
	heap.Fix(h.adapter, hi.index)
}

func (h MinHeap[T, N]) Len() int {
	return h.adapter.Len()
}

func (h MinHeap[T, N]) Empty() bool {
	return h.adapter.Len() == 0
}

func (h MinHeap[T, N]) Contains(item T) bool {
	_, ok := h.adapter.item2heapItem[item]
	return ok
}

func NewMinHeap[T comparable, N utils.Number]() MinHeap[T, N] {
	return MinHeap[T, N]{adapter: nil}
}

func NewMinHeapInt[T comparable]() MinHeap[T, int] {
	adapter := minHeapAdapter[T, int]{item2heapItem: make(map[T]*heapItem[T, int])}
	return MinHeap[T, int]{adapter: &adapter}
}

type heapItem[T any, N utils.Number] struct {
	item  T
	index int
	value N
}

type minHeapAdapter[T comparable, N utils.Number] struct {
	heapItems     []*heapItem[T, N]
	item2heapItem map[T]*heapItem[T, N]
}

func (h minHeapAdapter[T, N]) Len() int           { return len(h.heapItems) }
func (h minHeapAdapter[T, N]) Less(i, j int) bool { return h.heapItems[i].value < h.heapItems[j].value }
func (h minHeapAdapter[T, N]) Swap(i, j int) {
	h.heapItems[i], h.heapItems[j] = h.heapItems[j], h.heapItems[i]

	h.heapItems[i].index = i
	h.heapItems[j].index = j
}

func (h *minHeapAdapter[T, N]) Push(item any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	hi := item.(heapItem[T, N])
	hi.index = len(h.heapItems)
	h.heapItems = append(h.heapItems, &hi)

	h.item2heapItem[hi.item] = &hi
}

func (h *minHeapAdapter[T, N]) Pop() any {
	old := h.heapItems
	n := len(old)
	x := old[n-1]
	h.heapItems = old[0 : n-1]

	delete(h.item2heapItem, x.item)

	return x
}
