// Package testutil provides helpers for tests.
package testutil

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/require"
)

// TapMenuItem taps a menu item identified by label.
func TapMenuItem(t *testing.T, w fyne.Window, label string) {
	t.Helper()
	overlay := w.Canvas().Overlays().Top()
	require.NotNil(t, overlay, "expected a menu overlay to be showing")

	item := findObjectByText(overlay, label)
	require.NotNil(t, item, "menu item %q not found", label)

	pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(item)
	center := pos.Add(fyne.NewPos(item.Size().Width/2, item.Size().Height/2))
	test.TapCanvas(w.Canvas(), center)
}

// findObjectByText recursively searches obj and its children (including
// widget renderer output) for a text-bearing object matching want.
func findObjectByText(obj fyne.CanvasObject, want string) fyne.CanvasObject {
	switch v := obj.(type) {
	case *widget.Label:
		if v.Text == want {
			return v
		}
	case *canvas.Text:
		if v.Text == want {
			return v
		}
	}
	if c, ok := obj.(*fyne.Container); ok {
		for _, child := range c.Objects {
			if found := findObjectByText(child, want); found != nil {
				return found
			}
		}
	}
	if w, ok := obj.(fyne.Widget); ok {
		for _, child := range test.WidgetRenderer(w).Objects() {
			if found := findObjectByText(child, want); found != nil {
				return found
			}
		}
	}
	return nil
}
