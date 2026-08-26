package xslices_test

import (
	"testing"

	"github.com/ErikKalkoken/fyne-kx/internal/xslices"
	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
	t.Run("can remove duplicate elements", func(t *testing.T) {
		s := []string{"b", "a", "b"}
		got := xslices.Deduplicate(s)
		want := []string{"b", "a"}
		assert.Equal(t, want, got)
	})
	t.Run("can process slices with no duplicates", func(t *testing.T) {
		s := []string{"b", "a"}
		got := xslices.Deduplicate(s)
		want := []string{"b", "a"}
		assert.Equal(t, want, got)
	})
	t.Run("can process empty slice", func(t *testing.T) {
		s := []string{}
		got := xslices.Deduplicate(s)
		want := []string{}
		assert.ElementsMatch(t, want, got)
	})
}
