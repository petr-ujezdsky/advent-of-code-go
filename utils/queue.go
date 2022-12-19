package utils

import (
	"fmt"
	"strings"
)

// see https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/queue/circular-queue.go

type Queue[T any] struct {
	list                []T
	front, rear, length int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Peek() T {
	if q.Empty() {
		panic("Queue is empty")
	}
	return q.list[q.front]
}

func (q *Queue[T]) PeekRear() T {
	if q.Empty() {
		panic("Queue is empty")
	}
	return q.list[q.rear]
}

func (q *Queue[T]) PeekAll() []T {
	list := make([]T, q.length)

	j := 0
	i := q.front
	for {
		list[j] = q.list[i]
		i = (i + 1) % len(q.list)
		j++
		if i == q.rear {
			break
		}
	}

	return list
}

func (q *Queue[T]) Push(elm T) {
	// check size end enlarge if needed
	if q.length == len(q.list) {
		list := make([]T, Max(1, 2*len(q.list)))

		j := 0
		if len(q.list) > 0 {
			i := q.front
			for {
				list[j] = q.list[i]
				i = (i + 1) % len(q.list)
				j++
				if i == q.rear {
					break
				}
			}
		}

		q.list = list
		q.front = 0
		q.rear = j
	}

	q.list[q.rear] = elm
	q.rear = (q.rear + 1) % len(q.list)
	q.length++
}

func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("Queue is empty")
	}

	elm := q.list[q.front]
	q.front = (q.front + 1) % len(q.list)
	q.length--

	return elm
}

func (q *Queue[T]) Empty() bool {
	return q.Length() == 0
}

func (q *Queue[T]) Length() int {
	return q.length
}

func (q *Queue[T]) Clone() Queue[T] {
	return Queue[T]{
		list:   ShallowCopy(q.list),
		front:  q.front,
		rear:   q.rear,
		length: q.length,
	}
}

func (q *Queue[T]) String() string {
	sb := &strings.Builder{}

	sb.WriteString("[front <- ")

	if len(q.list) > 0 {
		i := q.front
		for {
			sb.WriteString(fmt.Sprint(q.list[i]))
			i = (i + 1) % len(q.list)
			if i == q.rear {
				break
			}
			sb.WriteString(" <- ")
		}
	}

	sb.WriteString(" <- rear]")

	return sb.String()
}
