package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"

	"github.com/ErikKalkoken/fyne-kx/internal/testutil"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestSortChip_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Alpha", kxwidget.SortOrderAscending, nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_minimal.png", w.Canvas().Capture())
}

func TestSortChip_CanChangeDefaultAfterCreation(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Alpha", kxwidget.SortOrderAscending, nil)
	c.DefaultColumn = "Bravo"
	c.DefaultOrder = kxwidget.SortOrderDescending
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_changed_default.png", w.Canvas().Capture())
}

func TestSortChip_CanCreateEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Alpha", kxwidget.SortOrderAscending, nil)
	c.Column = "Bravo"
	c.Order = kxwidget.SortOrderDescending
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_enabled.png", w.Canvas().Capture())
}

func TestSortChip_ShouldTreatInvalidColumnAsNotSet(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Alpha", kxwidget.SortOrderAscending, nil)
	c.Column = "Invalid"
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_minimal.png", w.Canvas().Capture())
}

func TestSortChip_ShouldRenderInvalidDefaultColumn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Invalid", kxwidget.SortOrderAscending, nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_unset.png", w.Canvas().Capture())
}

func TestSortChip_ShouldRenderInvalidDefaultOrder(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Alpha", 0, nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_no_sort.png", w.Canvas().Capture())
}

func TestSortChip_ShouldRenderAsEmptyWhenNoColumns(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{}, "", kxwidget.SortOrderAscending, nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_empty.png", w.Canvas().Capture())
}

func TestSortChip_DropDown(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	t.Run("columns are sanitized before shown as drop down", func(t *testing.T) {
		c := kxwidget.NewSortChip([]string{"b", "a", "a", "b", ""}, "b", kxwidget.SortOrderAscending, nil)
		w := test.NewWindow(container.NewVBox(c))
		defer w.Close()
		w.Resize(fyne.NewSize(300, 400))

		test.Tap(c)

		test.AssertImageMatches(t, "sortchip/santized_options_dropdown.png", w.Canvas().Capture())
	})
}

func TestSortChip_ResetSilent(t *testing.T) {
	// given
	test.NewTempApp(t)
	var called bool
	c := kxwidget.NewSortChip(
		[]string{"Alpha", "Bravo", "Charlie"},
		"Alpha",
		kxwidget.SortOrderAscending,
		func(col string, order kxwidget.SortOrder) {
			called = true
		},
	)
	c.Column = "Bravo"
	c.Order = kxwidget.SortOrderDescending
	w := test.NewWindow(c)
	defer w.Close()

	// when
	c.ResetSilent()

	// then
	assert.Equal(t, "Alpha", c.Column)
	assert.Equal(t, kxwidget.SortOrderAscending, c.Order)
	assert.False(t, called)
}

func TestSortChip_SelectColumnFromMenu(t *testing.T) {
	// given
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	var callCount int
	var gotColumn string
	var gotOrder kxwidget.SortOrder

	c := kxwidget.NewSortChip([]string{
		"Alpha", "Bravo", "Charlie"},
		"Alpha",
		kxwidget.SortOrderAscending,
		func(col string, order kxwidget.SortOrder) {
			callCount++
			gotColumn = col
			gotOrder = order
		},
	)
	w := test.NewWindow(container.NewVBox(c))
	defer w.Close()
	w.Resize(fyne.NewSize(300, 400))

	// when
	test.Tap(c)
	testutil.TapMenuItem(t, w, "Bravo")

	// then
	assert.Equal(t, 1, callCount)
	assert.Equal(t, "Bravo", gotColumn)
	assert.Equal(t, kxwidget.SortOrderAscending, gotOrder) // unchanged, only column selected
	assert.Equal(t, "Bravo", c.Column)
}

func TestSortChip_SelectOrderFromMenu(t *testing.T) {
	// given
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	var callCount int
	var gotColumn string
	var gotOrder kxwidget.SortOrder

	c := kxwidget.NewSortChip(
		[]string{"Alpha", "Bravo", "Charlie"},
		"Alpha",
		kxwidget.SortOrderAscending,
		func(col string, order kxwidget.SortOrder) {
			callCount++
			gotColumn = col
			gotOrder = order
		},
	)
	w := test.NewWindow(container.NewVBox(c))
	defer w.Close()
	w.Resize(fyne.NewSize(300, 400))

	// when
	test.Tap(c)
	testutil.TapMenuItem(t, w, "Descending")

	// then
	assert.Equal(t, 1, callCount)
	assert.Equal(t, "Alpha", gotColumn) // unchanged, only order selected
	assert.Equal(t, kxwidget.SortOrderDescending, gotOrder)
	assert.Equal(t, kxwidget.SortOrderDescending, c.Order)
}

func TestSortChip_ShouldUseColumnWhenDefaultColumnIsInvalid(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Invalid", kxwidget.SortOrderAscending, nil)
	c.Column = "Bravo" // explicit, valid — should NOT render as unset
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_enabled_bravo.png", w.Canvas().Capture())
}

func TestSortChip_ShouldUseOrderWhenDefaultOrderIsInvalid(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, "Alpha", 0, nil)
	c.Order = kxwidget.SortOrderDescending // explicit, valid — should NOT render as unset
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "sortchip/create_enabled_descending.png", w.Canvas().Capture())
}
