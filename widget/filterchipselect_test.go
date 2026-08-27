package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestFilterChipSelect_CanCreateEnabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Alpha", "Bravo"}, nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchipselect/enabled_off.png", w.Canvas().Capture())
}

func TestFilterChipSelect_CanCreateEnabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Alpha", "Bravo"}, nil)
	c.Selected = "Alpha"
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchipselect/enabled_on.png", w.Canvas().Capture())
}

func TestFilterChipSelect_CanCreateDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Alpha", "Bravo"}, nil)
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchipselect/disabled_off.png", w.Canvas().Capture())
}

func TestFilterChipSelect_CanCreateDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Alpha", "Bravo"}, nil)
	c.Selected = "Alpha"
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchipselect/disabled_on.png", w.Canvas().Capture())
}
func TestFilterChipSelect_NewFilterChipSelect(t *testing.T) {
	test.NewTempApp(t)
	t.Run("options are deduplicated", func(t *testing.T) {
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"b", "a", "b"}, nil)
		assert.Equal(t, []string{"b", "a"}, x.Options)
	})
}

func TestFilterChipSelected_NewFilterChipSelectWithSearch(t *testing.T) {
	a := test.NewTempApp(t)
	w := a.NewWindow("Dummy")
	t.Run("options are deduplicated", func(t *testing.T) {
		x := kxwidget.NewFilterChipSelectWithSearch("placeholder", []string{"b", "a", "b", "a"}, nil, w)
		assert.Equal(t, []string{"b", "a"}, x.Options)
	})
}

func TestFilterChipSelect_SetSelected(t *testing.T) {
	test.NewTempApp(t)
	t.Run("can select an option", func(t *testing.T) {
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b"}, nil)
		x.SetSelected("a")
		assert.Equal(t, "a", x.Selected)
		assert.Equal(t, []string{"a", "b"}, x.Options)
	})
	t.Run("selecting invalid option is ignored", func(t *testing.T) {
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b"}, nil)
		x.SetSelected("a")
		x.SetSelected("x")
		assert.Equal(t, "a", x.Selected)
	})
	t.Run("can clear selection", func(t *testing.T) {
		// given
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b"}, nil)
		x.SetSelected("a")
		// when
		x.SetSelected("")
		// then
		assert.Equal(t, "", x.Selected)
	})
	t.Run("can not clear selection when no placeholder", func(t *testing.T) {
		// given
		x := kxwidget.NewFilterChipSelect("", []string{"a", "b"}, nil)
		x.SetSelected("a")
		// when
		x.SetSelected("")
		// then
		assert.Equal(t, "a", x.Selected)
	})
	t.Run("selecting an option triggers callback when changed", func(t *testing.T) {
		var isCalled bool
		var v string
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b"}, func(selected string) {
			isCalled = true
			v = selected
		})
		x.SetSelected("a")
		assert.True(t, isCalled)
		assert.Equal(t, "a", v)
	})
	t.Run("selecting an option does not trigger callback when not changed", func(t *testing.T) {
		// given
		var isCalled bool
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b"}, func(selected string) {
			isCalled = true
		})
		x.SetSelected("a")
		isCalled = false
		// when
		x.SetSelected("a")
		// then
		assert.False(t, isCalled)
	})
	t.Run("options are deduplicated, but not sorted when disabled", func(t *testing.T) {
		x := kxwidget.NewFilterChipSelect("", []string{"b", "a", "b"}, nil)
		assert.Equal(t, []string{"b", "a"}, x.Options)
	})
}

func TestFilterChipSelect_ClearSelected(t *testing.T) {
	test.NewTempApp(t)
	t.Run("can clear selection", func(t *testing.T) {
		// given
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b"}, nil)
		x.SetSelected("a")
		// when
		x.ClearSelected()
		// then
		assert.Equal(t, "", x.Selected)
	})
	t.Run("clearing selection triggers callback", func(t *testing.T) {
		// given
		var isCalled bool
		var v string
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"a", "b", "c"}, func(selected string) {
			isCalled = true
			v = selected
		})
		x.SetSelected("a")
		isCalled = false
		v = "xx"
		// when
		x.ClearSelected()
		// then
		assert.True(t, isCalled)
		assert.Equal(t, "", v)
	})
}

func TestFilterChipSelected_SetOptions(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	t.Run("options are sanitized before shown as drop down", func(t *testing.T) {
		c := kxwidget.NewFilterChipSelect("placeholder", []string{}, nil)
		c.Options = []string{"b", "a", "b", "a", ""}
		w := test.NewWindow(container.NewCenter(c))
		defer w.Close()
		w.Resize(fyne.NewSize(200, 200))

		test.Tap(c)

		test.AssertImageMatches(t, "filterchipselect/santized_options_dropdown.png", w.Canvas().Capture())
	})

	t.Run("options are sanitized before shown in search modal", func(t *testing.T) {
		w := test.NewWindow(nil)
		defer w.Close()
		w.Resize(fyne.NewSize(200, 200))
		c := kxwidget.NewFilterChipSelectWithSearch("placeholder", []string{}, nil, w)
		c.Options = []string{"b", "a", "b", "a", ""}
		w.SetContent(container.NewCenter(c))

		test.Tap(c)

		test.AssertImageMatches(t, "filterchipselect/santized_options_modal.png", w.Canvas().Capture())
	})

	t.Run("selection is not cleared when no longer valid", func(t *testing.T) {
		// given
		x := kxwidget.NewFilterChipSelect("placeholder", []string{"c"}, nil)
		x.SetSelected("c")
		// when
		x.SetOptions([]string{"a"})
		// then
		assert.Equal(t, "c", x.Selected)
	})
}

func TestFilterChipSelect_CanShowDropDownSorted(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Bravo", "Alpha"}, nil)
	w := test.NewWindow(container.NewCenter(c))
	defer w.Close()
	w.Resize(fyne.NewSize(100, 200))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/dropdown_sorted.png", w.Canvas().Capture())
}

func TestFilterChipSelect_CanShowDropDownUnSorted(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Bravo", "Alpha"}, nil)
	c.SortDisabled = true
	w := test.NewWindow(container.NewCenter(c))
	defer w.Close()
	w.Resize(fyne.NewSize(100, 200))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/dropdown_unsorted.png", w.Canvas().Capture())
}

func TestFilterChipSelect_DropDownContainsSelectedWhenOtherOptions(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{"Bravo", "Alpha"}, nil)
	c.Selected = "Charlie"
	w := test.NewWindow(container.NewCenter(c))
	defer w.Close()
	w.Resize(fyne.NewSize(200, 500))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/dropdown_roque_option.png", w.Canvas().Capture())
}

func TestFilterChipSelect_DropDownContainsSelectedWhenNoOptions(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChipSelect("Test", []string{}, nil)
	c.Selected = "Charlie"
	w := test.NewWindow(container.NewCenter(c))
	defer w.Close()
	w.Resize(fyne.NewSize(200, 300))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/dropdown_roque_option_2.png", w.Canvas().Capture())
}

func TestFilterChipSelect_CanShowSearchBoxSorted(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	w := test.NewWindow(nil)
	defer w.Close()
	w.Resize(fyne.NewSize(500, 500))
	c := kxwidget.NewFilterChipSelectWithSearch("Test", []string{"Bravo", "Alpha"}, nil, w)
	w.SetContent(container.NewCenter(c))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/search_sorted.png", w.Canvas().Capture())
}

func TestFilterChipSelect_CanShowSearchBoxUnsorted(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	w := test.NewWindow(nil)
	defer w.Close()
	w.Resize(fyne.NewSize(500, 500))
	c := kxwidget.NewFilterChipSelectWithSearch("Test", []string{"Bravo", "Alpha"}, nil, w)
	c.SortDisabled = true
	w.SetContent(container.NewCenter(c))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/search_unsorted.png", w.Canvas().Capture())
}

func TestFilterChipSelect_SearchBoxContainsSelectedWhenOtherOptions(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	w := test.NewWindow(nil)
	defer w.Close()
	w.Resize(fyne.NewSize(500, 500))
	c := kxwidget.NewFilterChipSelectWithSearch("Test", []string{"Bravo", "Alpha"}, nil, w)
	c.Selected = "Charlie"
	w.SetContent(container.NewCenter(c))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/search_roque_option.png", w.Canvas().Capture())
}

func TestFilterChipSelect_SearchBoxContainsSelectedWhenNoOptions(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	w := test.NewWindow(nil)
	defer w.Close()
	w.Resize(fyne.NewSize(500, 500))
	c := kxwidget.NewFilterChipSelectWithSearch("Test", []string{}, nil, w)
	c.Selected = "Charlie"
	w.SetContent(container.NewCenter(c))

	test.Tap(c)

	test.AssertImageMatches(t, "filterchipselect/search_roque_option_2.png", w.Canvas().Capture())
}
