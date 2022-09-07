package queue

type Queue[T any] struct {
	q []T
	s int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]T, 0),
	}
}

func NewQueueWith[T any](v []T) *Queue[T] {
	if v == nil {
		return NewQueue[T]()
	}

	return &Queue[T]{
		q: v,
		s: len(v),
	}
}

func (q *Queue[T]) Enqueue(n T) {
	q.q = append(q.q, n)
	q.s++
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
	q.s--
	return
}

func (q *Queue[T]) Empty() bool {
	return q.s == 0
}
