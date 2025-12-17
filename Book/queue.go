package Book

type Queue[T any] struct {
	items []T
	head  int
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() T {
	if q.head >= len(q.items) {
		panic("dequeue from empty queue")
	}
	v := q.items[q.head]
	q.head++

	// Optional cleanup to avoid memory leak
	if q.head > 64 && q.head*2 >= len(q.items) {
		q.items = append([]T(nil), q.items[q.head:]...)
		q.head = 0
	}

	return v
}

func (q *Queue[T]) Empty() bool {
	return q.head >= len(q.items)
}
