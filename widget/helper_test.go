package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
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
		assert.ElementsMatch(t, want, got)
	})
}
