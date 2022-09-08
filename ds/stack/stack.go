package stack

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Stack[T any] struct {
	*arraystack.Stack
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Stack: arraystack.New(),
	}
}

func NewStackWith[T any](v []T) *Stack[T] {
	s := arraystack.New()
	for _, e := range v {
		s.Push(e)
	}
	return &Stack[T]{
		Stack: s,
	}
}

func (s *Stack[T]) Push(v T) {
	s.Stack.Push(v)
}

func (s *Stack[T]) Pop() (v T) {
	val, ok := s.Stack.Pop()
	if !ok {
		return
	}
	v, ok = val.(T)
	if !ok {
		return
	}
	return v
}
