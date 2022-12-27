package utils

// MinHeap is implementation of binary min-heap
// see https://maupanelo.com/posts/how-to-write-a-binary-heap-in-golang/
type MinHeap[T any] struct {
	heap []T
	less func(i, j T) bool
}

func NewMinHeap[T any](less func(i, j T) bool) MinHeap[T] {
	return MinHeap[T]{heap: nil, less: less}
}

func (bh *MinHeap[T]) Push(item T) {
	bh.heap = append(bh.heap, item)
	bh.bubbleUp(len(bh.heap) - 1)
}

func (bh *MinHeap[T]) Pop() T {
	popped := bh.heap[0]
	heapSize := len(bh.heap)

	if heapSize > 1 {
		bh.heap[0] = bh.heap[len(bh.heap)-1]
	}

	bh.heap = bh.heap[:len(bh.heap)-1]
	bh.bubbleDown(0)
	return popped
}

func (bh *MinHeap[T]) Empty() bool {
	return len(bh.heap) == 0
}

func (bh *MinHeap[T]) bubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2

		if bh.less(bh.heap[parentIndex], bh.heap[index]) {
			return
		}

		bh.heap[parentIndex], bh.heap[index] = bh.heap[index], bh.heap[parentIndex]
		index = parentIndex
	}
}

func (bh *MinHeap[T]) bubbleDown(index int) {
	for 2*index+1 < len(bh.heap) {
		minChildIndex := bh.minChildIndex(index)

		if bh.less(bh.heap[index], bh.heap[minChildIndex]) {
			return
		}

		bh.heap[minChildIndex], bh.heap[index] = bh.heap[index], bh.heap[minChildIndex]
		index = minChildIndex
	}
}

func (bh *MinHeap[T]) minChildIndex(index int) int {
	if 2*index+2 >= len(bh.heap) {
		return 2*index + 1
	}

	if bh.less(bh.heap[2*index+2], bh.heap[2*index+1]) {
		return 2*index + 2
	}

	return 2*index + 1
}

// Integer cost heap simplification

type MinHeapInt[T any] struct {
	minHeap MinHeap[heapItemInt[T]]
}

type heapItemInt[T any] struct {
	value T
	cost  int
}

func NewMinHeapInt[T any]() MinHeapInt[T] {
	return MinHeapInt[T]{minHeap: NewMinHeap[heapItemInt[T]](func(i, j heapItemInt[T]) bool { return i.cost < j.cost })}
}

func (bh *MinHeapInt[T]) Push(item T, cost int) {
	heapItem := heapItemInt[T]{value: item, cost: cost}
	bh.minHeap.Push(heapItem)
}

func (bh *MinHeapInt[T]) PopWithCost() (T, int) {
	item := bh.minHeap.Pop()
	return item.value, item.cost
}

func (bh *MinHeapInt[T]) Pop() T {
	return bh.minHeap.Pop().value
}

func (bh *MinHeapInt[T]) Empty() bool {
	return bh.minHeap.Empty()
}
