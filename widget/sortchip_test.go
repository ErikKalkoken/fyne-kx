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

func TestSortChip_CanCreateWithDefaults(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, nil)
	w := test.NewWindow(c)
	defer w.Close()

	assert.Equal(t, "Alpha", c.DefaultColumn)
	assert.Equal(t, kxwidget.SortOrderAscending, c.DefaultOrder)
	assert.Equal(t, "Alpha", c.Column)
	assert.Equal(t, kxwidget.SortOrderAscending, c.Order)
	test.AssertImageMatches(t, "sortchip/create_default.png", w.Canvas().Capture())
}

func TestSortChip_CanCreateWithConfiguredDefaults(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, nil)
	c.DefaultColumn = "Bravo"
	c.DefaultOrder = kxwidget.SortOrderDescending
	w := test.NewWindow(c)
	defer w.Close()

	assert.Equal(t, "Bravo", c.DefaultColumn)
	assert.Equal(t, kxwidget.SortOrderDescending, c.DefaultOrder)
	assert.Equal(t, "Bravo", c.Column)
	assert.Equal(t, kxwidget.SortOrderDescending, c.Order)
	test.AssertImageMatches(t, "sortchip/create_configured.png", w.Canvas().Capture())
}

func TestSortChip_CanCreateEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, nil)
	c.Column = "Bravo"
	c.Order = kxwidget.SortOrderDescending
	w := test.NewWindow(c)
	defer w.Close()

	assert.Equal(t, "Alpha", c.DefaultColumn)
	assert.Equal(t, kxwidget.SortOrderAscending, c.DefaultOrder)
	assert.Equal(t, "Bravo", c.Column)
	assert.Equal(t, kxwidget.SortOrderDescending, c.Order)
	test.AssertImageMatches(t, "sortchip/create_enabled.png", w.Canvas().Capture())
}

func TestSortChip_CanCreateEmpty(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{}, nil)
	w := test.NewWindow(c)
	defer w.Close()

	assert.Equal(t, "", c.DefaultColumn)
	assert.Equal(t, kxwidget.SortOrderNone, c.DefaultOrder)
	assert.Equal(t, "", c.Column)
	assert.Equal(t, kxwidget.SortOrderNone, c.Order)
	test.AssertImageMatches(t, "sortchip/create_empty.png", w.Canvas().Capture())
}

func TestSortChip_ShouldResetInvalidConfigOnCreation(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, nil)
	c.DefaultColumn = "invalid"
	c.DefaultOrder = kxwidget.SortOrderNone
	c.Column = "invalid"
	c.Order = kxwidget.SortOrderNone
	w := test.NewWindow(c)
	defer w.Close()

	assert.Equal(t, "Alpha", c.DefaultColumn)
	assert.Equal(t, kxwidget.SortOrderAscending, c.DefaultOrder)
	assert.Equal(t, "Alpha", c.Column)
	assert.Equal(t, kxwidget.SortOrderAscending, c.Order)
}

func TestSortChip_ShouldResetInvalidConfigOnRefresh(t *testing.T) {
	// given
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, nil)
	w := test.NewWindow(c)
	defer w.Close()

	// when
	c.DefaultColumn = "invalid"
	c.DefaultOrder = kxwidget.SortOrderNone
	c.Column = "invalid"
	c.Order = kxwidget.SortOrderNone
	c.Refresh()

	// then
	assert.Equal(t, "Alpha", c.DefaultColumn)
	assert.Equal(t, kxwidget.SortOrderAscending, c.DefaultOrder)
	assert.Equal(t, "Alpha", c.Column)
	assert.Equal(t, kxwidget.SortOrderAscending, c.Order)
}

func TestSortChip_DropDown(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	t.Run("columns are sanitized before shown as drop down", func(t *testing.T) {
		c := kxwidget.NewSortChip([]string{"b", "a", "a", "b", ""}, nil)
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
	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, func(col string, order kxwidget.SortOrder) {
		called = true
	})
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

	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, func(col string, order kxwidget.SortOrder) {
		callCount++
		gotColumn = col
		gotOrder = order
	})
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

	c := kxwidget.NewSortChip([]string{"Alpha", "Bravo", "Charlie"}, func(col string, order kxwidget.SortOrder) {
		callCount++
		gotColumn = col
		gotOrder = order
	})
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
