package stack_test

import (
	"sync"
	"testing"

	"github.com/ErikKalkoken/fyne-kx/internal/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack_Pop(t *testing.T) {
	t.Run("should return items in LIFO order", func(t *testing.T) {
		st := stack.New[int]()
		st.Push(99)
		st.Push(42)

		v, ok := st.Pop()
		require.True(t, ok)
		assert.Equal(t, 42, v)
		v, ok = st.Pop()
		require.True(t, ok)
		assert.Equal(t, 99, v)
	})

	t.Run("should return false and zero value when trying to pop from empty stack", func(t *testing.T) {
		st := stack.New[int]()
		v, ok := st.Pop()
		assert.False(t, ok)
		assert.Zero(t, v)
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("should return stack size when not empty", func(t *testing.T) {
		st := stack.New[int]()
		st.Push(99)
		st.Push(42)
		assert.Equal(t, 2, st.Size())
	})

	t.Run("should return stack size when empty", func(t *testing.T) {
		st := stack.New[int]()
		assert.Equal(t, 0, st.Size())
	})
}

func TestStack_LifecycleAndMemoryReset(t *testing.T) {
	t.Run("should accurately reflect Size during push/pop lifecycle", func(t *testing.T) {
		st := stack.New[string]()
		require.Equal(t, 0, st.Size())

		st.Push("first")
		st.Push("second")
		require.Equal(t, 2, st.Size())

		v1, ok := st.Pop()
		require.True(t, ok)
		assert.Equal(t, "second", v1)
		require.Equal(t, 1, st.Size())

		v2, ok := st.Pop()
		require.True(t, ok)
		assert.Equal(t, "first", v2)
		require.Equal(t, 0, st.Size())

		// Extra pop should fail
		_, ok = st.Pop()
		require.False(t, ok)
	})
}

func TestStack_Concurrent(t *testing.T) {
	t.Run("should handle concurrent pushes and pops safely", func(t *testing.T) {
		const goroutines = 100
		const itemsPerGoroutine = 100

		st := stack.New[int]()
		var wg sync.WaitGroup

		// Concurrently push items
		wg.Add(goroutines)
		for i := range goroutines {
			go func(base int) {
				defer wg.Done()
				for j := range itemsPerGoroutine {
					st.Push(base + j)
				}
			}(i * itemsPerGoroutine)
		}

		// Concurrently call Size() while pushing
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range goroutines {
				_ = st.Size()
			}
		}()

		wg.Wait()

		// Verify total items pushed
		require.Equal(t, goroutines*itemsPerGoroutine, st.Size())

		// Concurrently pop all items
		popped := make(chan int, goroutines*itemsPerGoroutine)
		wg.Add(goroutines)

		for range goroutines {
			go func() {
				defer wg.Done()
				for range itemsPerGoroutine {
					if v, ok := st.Pop(); ok {
						popped <- v
					}
				}
			}()
		}

		wg.Wait()
		close(popped)

		// Verify final stack state
		assert.Equal(t, 0, st.Size())
		assert.Equal(t, goroutines*itemsPerGoroutine, len(popped))

		// Ensure popping on empty stack yields ErrEmpty
		_, ok := st.Pop()
		require.False(t, ok)
	})
}
