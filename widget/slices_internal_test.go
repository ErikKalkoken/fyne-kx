package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicateSlice(t *testing.T) {
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
