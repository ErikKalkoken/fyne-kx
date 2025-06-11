package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceDeduplicate(t *testing.T) {
	t.Run("can remove duplicate elements", func(t *testing.T) {
		s := []string{"b", "a", "b"}
		got := sliceDeduplicate(s)
		want := []string{"b", "a"}
		assert.Equal(t, want, got)
	})
	t.Run("can process slices with no duplicates", func(t *testing.T) {
		s := []string{"b", "a"}
		got := sliceDeduplicate(s)
		want := []string{"b", "a"}
		assert.Equal(t, want, got)
	})
	t.Run("can process empty slice", func(t *testing.T) {
		s := []string{}
		got := sliceDeduplicate(s)
		want := []string{}
		assert.Equal(t, want, got)
	})
}

func TestSliceDeleteFunc(t *testing.T) {
	t.Run("can delete matching elements", func(t *testing.T) {
		s := []string{"b", "x", "a"}
		got := sliceDeleteFunc(s, func(x string) bool {
			return x == "x"
		})
		want := []string{"b", "a"}
		assert.Equal(t, want, got)
	})
	t.Run("can delete empty elements", func(t *testing.T) {
		s := []string{"b", ""}
		got := sliceDeleteFunc(s, func(x string) bool {
			return x == ""
		})
		want := []string{"b"}
		assert.Equal(t, want, got)
	})
}

func TestSliceClone(t *testing.T) {
	t.Run("normal slice", func(t *testing.T) {
		s := []string{"b", "x", "a"}
		got := sliceClone(s)
		want := []string{"b", "x", "a"}
		assert.Equal(t, want, got)
	})
	t.Run("empty slice", func(t *testing.T) {
		s := []string{}
		got := sliceClone(s)
		want := []string{}
		assert.Equal(t, want, got)
	})
}

func TestSliceContains(t *testing.T) {
	t.Run("return true when element found", func(t *testing.T) {
		s := []string{"b", "x", "a"}
		assert.True(t, sliceContains(s, "x"))
	})
	t.Run("return false when element not found", func(t *testing.T) {
		s := []string{"b", "x", "a"}
		assert.False(t, sliceContains(s, "y"))
	})
	t.Run("can find zero elements", func(t *testing.T) {
		s := []string{"b", "", "a"}
		assert.True(t, sliceContains(s, ""))
	})
}
