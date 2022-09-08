package queue

import (
	"github.com/emirpasic/gods/queues/arrayqueue"
)

type Queue[T any] struct {
	*arrayqueue.Queue
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Queue: arrayqueue.New(),
	}
}

func NewQueueWith[T any](v []T) *Queue[T] {
	q := arrayqueue.New()
	for _, e := range v {
		q.Enqueue(e)
	}
	return &Queue[T]{
		Queue: q,
	}
}

func (q *Queue[T]) Enqueue(n T) {
	q.Queue.Enqueue(n)
}

func (q *Queue[T]) Dequeue() (v T) {
	val, ok := q.Queue.Dequeue()
	if !ok {
		return
	}
	v, ok = val.(T)
	if !ok {
		return
	}
	return v
}
