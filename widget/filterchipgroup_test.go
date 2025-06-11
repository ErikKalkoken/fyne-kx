package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestFilterChipGroup_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	t.Run("options are deduplicated and cleaned", func(t *testing.T) {
		x := kxwidget.NewFilterChipGroup([]string{"a", "c", "c", "", "d"}, nil)
		want := []string{"a", "c", "d"}
		assert.Equal(t, want, x.Options)
	})
	t.Run("can set initial selection", func(t *testing.T) {
		x := kxwidget.NewFilterChipGroup([]string{"a", "b"}, nil)
		x.Selected = []string{"a"}
		w := test.NewWindow(x)
		defer w.Close()
		w.Resize(fyne.NewSize(150, 50))

		assert.Equal(t, []string{"a"}, x.Selected)
		test.AssertImageMatches(t, "filterchipgroup/initial.png", w.Canvas().Capture())
	})
}

func TestFilterChipGroup_SetSelection(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	t.Run("can select option", func(t *testing.T) {
		tapped := make([][]string, 0)
		x := kxwidget.NewFilterChipGroup([]string{"a", "b", "c"}, func(s []string) {
			tapped = append(tapped, s)
		})
		w := test.NewWindow(x)
		defer w.Close()
		w.Resize(fyne.NewSize(250, 50))

		x.SetSelected([]string{"b", "a"})

		assert.Equal(t, []string{"a", "b"}, x.Selected)
		assert.Equal(t, [][]string{{"a", "b"}}, tapped)
		test.AssertImageMatches(t, "filterchipgroup/setselected_select.png", w.Canvas().Capture())
	})
	t.Run("can deselect option", func(t *testing.T) {
		x := kxwidget.NewFilterChipGroup([]string{"a", "b"}, nil)
		x.Selected = []string{"a"}
		w := test.NewWindow(x)
		defer w.Close()
		w.Resize(fyne.NewSize(150, 50))

		x.SetSelected([]string{})

		assert.Equal(t, []string{}, x.Selected)
		test.AssertImageMatches(t, "filterchipgroup/setselected_deselect.png", w.Canvas().Capture())
	})
	t.Run("can ignore invalid elements in selection", func(t *testing.T) {
		x := kxwidget.NewFilterChipGroup([]string{"a", "b"}, nil)
		w := test.NewWindow(x)
		defer w.Close()
		w.Resize(fyne.NewSize(150, 50))

		x.SetSelected([]string{"b", "c", ""})
		assert.Equal(t, []string{"b"}, x.Selected)
		test.AssertImageMatches(t, "filterchipgroup/setselected_ignore.png", w.Canvas().Capture())
	})
}
