// Package queue provides queues.
package stack

import (
	"errors"
	"sync"
)

var ErrEmpty = errors.New("empty stack")

// Stack represents a basic stack which can be used concurrently.
type Stack[T any] struct {
	mu sync.Mutex
	s  []T
}

// New returns a new [Stack].
func New[T any]() *Stack[T] {
	st := &Stack[T]{s: make([]T, 0)}
	return st
}

// Push adds an item on the stack.
func (st *Stack[T]) Push(v T) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.s = append(st.s, v)
}

// Pop returns the item from the top of the stack or an error when the stack is empty.
func (st *Stack[T]) Pop() (T, error) {
	var v T
	st.mu.Lock()
	defer st.mu.Unlock()
	if len(st.s) == 0 {
		return v, ErrEmpty
	}
	v = st.s[len(st.s)-1]
	st.s = st.s[:len(st.s)-1]
	return v, nil
}

// Size returns the number of items in the stack.
func (st *Stack[T]) Size() int {
	st.mu.Lock()
	defer st.mu.Unlock()
	return len(st.s)
}
