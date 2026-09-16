package main

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

// TestMakeFunctions is a smoke test which ensures every page constructor in
// the demo app runs without panicking and returns a usable canvas object.
func TestMakeFunctions(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")
	defer w.Close()

	cases := []struct {
		name string
		make func() fyne.CanvasObject
	}{
		{"All", makeAll},
		{"Badge", makeBadge},
		{"Columns", makeColumns},
		{"Dialogs", func() fyne.CanvasObject { return makeDialogs(w) }},
		{"FilterChip", makeFilterChip},
		{"FilterChipGroup", makeFilterChipGroup},
		{"FilterChipSelect", func() fyne.CanvasObject { return makeFilterChipSelect(w) }},
		{"IconButton", makeIconButton},
		{"Modals", func() fyne.CanvasObject { return makeModals(w) }},
		{"LoadingButton", makeLoadingButton},
		{"Slider", makeSlider},
		{"SortChip", makeSortChip},
		{"Switch", makeSwitch},
		{"TappableIcon", makeTappableIcon},
		{"TappableImage", makeTappableImage},
		{"TappableLabel", makeTappableLabel},
		{"ToolbarActionMenu", makeToolbarActionMenu},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			obj := c.make()
			assert.NotNil(t, obj)
		})
	}
}
