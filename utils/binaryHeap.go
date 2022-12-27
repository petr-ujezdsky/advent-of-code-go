package utils

// BinaryHeap is implementation of binary min-heap
// see https://maupanelo.com/posts/how-to-write-a-binary-heap-in-golang/
type BinaryHeap[T any] struct {
	heap []T
	less func(i, j T) bool
}

func NewBinaryHeap[T any](less func(i, j T) bool) BinaryHeap[T] {
	return BinaryHeap[T]{heap: nil, less: less}
}

func (bh *BinaryHeap[T]) Push(item T) {
	bh.heap = append(bh.heap, item)
	bh.bubbleUp(len(bh.heap) - 1)
}

func (bh *BinaryHeap[T]) Pop() T {
	popped := bh.heap[0]
	heapSize := len(bh.heap)

	if heapSize > 1 {
		bh.heap[0] = bh.heap[len(bh.heap)-1]
	}

	bh.heap = bh.heap[:len(bh.heap)-1]
	bh.bubbleDown(0)
	return popped
}

func (bh *BinaryHeap[T]) Empty() bool {
	return len(bh.heap) == 0
}

func (bh *BinaryHeap[T]) bubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2

		if bh.less(bh.heap[parentIndex], bh.heap[index]) {
			return
		}

		bh.heap[parentIndex], bh.heap[index] = bh.heap[index], bh.heap[parentIndex]
		index = parentIndex
	}
}

func (bh *BinaryHeap[T]) bubbleDown(index int) {
	for 2*index+1 < len(bh.heap) {
		minChildIndex := bh.minChildIndex(index)

		// this is actually >= and intended was only >, but it should be OK
		if !bh.less(bh.heap[minChildIndex], bh.heap[index]) {
			return
		}

		bh.heap[minChildIndex], bh.heap[index] = bh.heap[index], bh.heap[minChildIndex]
		index = minChildIndex
	}
}

func (bh *BinaryHeap[T]) minChildIndex(index int) int {
	if 2*index+2 >= len(bh.heap) {
		return 2*index + 1
	}

	if bh.less(bh.heap[2*index+2], bh.heap[2*index+1]) {
		return 2*index + 2
	}

	return 2*index + 1
}

// Integer cost heap simplification

type BinaryHeapInt[T any] struct {
	binaryHeap BinaryHeap[heapItemInt[T]]
}

type heapItemInt[T any] struct {
	value T
	cost  int
}

func NewBinaryHeapInt[T any]() BinaryHeapInt[T] {
	return BinaryHeapInt[T]{binaryHeap: NewBinaryHeap[heapItemInt[T]](func(i, j heapItemInt[T]) bool { return i.cost < j.cost })}
}

func (bh *BinaryHeapInt[T]) Push(item T, cost int) {
	heapItem := heapItemInt[T]{value: item, cost: cost}
	bh.binaryHeap.Push(heapItem)
}

func (bh *BinaryHeapInt[T]) Pop() T {
	return bh.binaryHeap.Pop().value
}

func (bh *BinaryHeapInt[T]) Empty() bool {
	return bh.binaryHeap.Empty()
}
