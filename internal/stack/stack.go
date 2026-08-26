// Package stack provides a basic stack.
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
	return &Stack[T]{}
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
	idx := len(st.s) - 1
	v = st.s[idx]
	clear(st.s[idx:]) // zeros out the element for GC
	st.s = st.s[:idx]
	return v, nil
}

// Size returns the number of items in the stack.
func (st *Stack[T]) Size() int {
	st.mu.Lock()
	defer st.mu.Unlock()
	return len(st.s)
}
