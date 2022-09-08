package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackFloat64(t *testing.T) {
	expectation := []float64{1.11, 2.22, 3, 4.99, 5.00, 1e3}

	s := NewStackWith[float64](expectation)
	actual := make([]float64, 0)
	for !s.Empty() {
		actual = append([]float64{s.Pop()}, actual...)
	}

	assert.Equal(t, expectation, actual)
}

func TestNewStackWith(t *testing.T) {
	s := NewStackWith[int](nil)
	s.Push(5)

	assert.NotEmpty(t, s.Pop())
	assert.Empty(t, s.Pop())
}

func TestNewStack(t *testing.T) {
	s := NewStack[string]()

	assert.Empty(t, s.Pop())
}
