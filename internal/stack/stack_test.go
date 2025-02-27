package stack_test

import (
	"testing"

	"github.com/ErikKalkoken/fyne-kx/internal/stack"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("can push and pop", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(99)
		v, err := s.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, 99, v)
		}
	})
	t.Run("should return specific error when trying to pop from empty stack", func(t *testing.T) {
		s := stack.New[int]()
		_, err := s.Pop()
		assert.ErrorIs(t, stack.ErrEmpty, err)
	})
	t.Run("should return correct stack size", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(99)
		s.Push(42)
		v := s.Size()
		assert.Equal(t, 2, v)
	})
}
