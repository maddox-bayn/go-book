package Book

type RingQueue[T any] struct {
	buf        []T
	head, tail int
	size       int
}

func NewRingQueue[T any](capacity int) *RingQueue[T] {
	return &RingQueue[T]{buf: make([]T, capacity)}
}

func (q *RingQueue[T]) Enqueue(v T) {
	if q.size == len(q.buf) {
		panic("queue full")
	}
	q.buf[q.tail] = v
	q.tail = (q.tail + 1) % len(q.buf)
	q.size++
}

func (q *RingQueue[T]) Dequeue() T {
	if q.size == 0 {
		panic("empty queue")
	}
	v := q.buf[q.head]
	q.head = (q.head + 1) % len(q.buf)
	q.size--
	return v
}
