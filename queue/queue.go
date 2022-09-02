package queue

type Queue[T any] struct {
	q []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]T, 0),
	}
}

func NewQueueWith[T any](v []T) *Queue[T] {
	return &Queue[T]{
		q: v,
	}
}

func (q *Queue[T]) Enqueue(n T) {
	q.q = append(q.q, n)
}

func (q *Queue[T]) Dequeue() (v T) {
	if q.Empty() {
		return
	}

	f := make([]T, 1)
	copy(f, q.q)
	v = f[0]

	nq := make([]T, len(q.q)-1)
	copy(nq, q.q[1:])
	q.q = nil
	q.q = nq
	return
}

func (q *Queue[T]) Empty() bool {
	return len(q.q) == 0
}
