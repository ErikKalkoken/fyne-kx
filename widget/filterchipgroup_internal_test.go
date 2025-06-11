package widget

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"

	"github.com/stretchr/testify/assert"
)

func TestFilterChipGroup(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	t.Run("can select option via tap", func(t *testing.T) {
		tapped := make([][]string, 0)
		x := NewFilterChipGroup([]string{"a", "b"}, func(s []string) {
			tapped = append(tapped, s)
		})
		w := test.NewWindow(x)
		defer w.Close()
		w.Resize(fyne.NewSize(150, 50))

		test.Tap(x.chips[1])

		assert.Equal(t, []string{"b"}, x.Selected)
		assert.Equal(t, [][]string{{"b"}}, tapped)
		test.AssertImageMatches(t, "filterchipgroup/tap_select.png", w.Canvas().Capture())
	})
	t.Run("can deselect option via tap", func(t *testing.T) {
		tapped := make([][]string, 0)
		x := NewFilterChipGroup([]string{"a", "b"}, func(s []string) {
			tapped = append(tapped, s)
		})
		x.Selected = []string{"a", "b"}
		w := test.NewWindow(x)
		defer w.Close()
		w.Resize(fyne.NewSize(150, 50))

		test.Tap(x.chips[1])

		assert.Equal(t, []string{"a"}, x.Selected)
		assert.Equal(t, [][]string{{"a"}}, tapped)
		test.AssertImageMatches(t, "filterchipgroup/tap_deselect.png", w.Canvas().Capture())
	})
}
