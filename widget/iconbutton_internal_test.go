package widget

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

type embeddedIconButton struct {
	IconButton
}

func TestIconButton_Super(t *testing.T) {
	t.Run("returns itself when used directly", func(t *testing.T) {
		w := NewIconButton(theme.HomeIcon(), nil)
		assert.Same(t, w, w.super())
	})
	t.Run("returns outer widget when embedded", func(t *testing.T) {
		w := &embeddedIconButton{}
		w.ExtendBaseWidget(w)
		assert.Same(t, w, w.super())
	})
}

func TestIconButton_HoverBackground(t *testing.T) {
	cases := []struct {
		name  string
		setup func(w *IconButton)
		want  bool
	}{
		{"callback", func(w *IconButton) { w.OnTapped = func() {} }, true},
		{"menu with items", func(w *IconButton) {
			w.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("item", nil)})
		}, true},
		{"no callback", func(w *IconButton) {}, false},
		{"empty menu", func(w *IconButton) { w.SetMenuItems(nil) }, false},
		{"disabled while hovered", func(w *IconButton) {
			w.OnTapped = func() {}
			w.MouseIn(&desktop.MouseEvent{})
			w.Disable()
		}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			test.NewTempApp(t)
			w := NewIconButton(theme.HomeIcon(), nil)
			r := test.TempWidgetRenderer(t, w).(*iconButtonRenderer)
			tc.setup(w)
			w.MouseIn(&desktop.MouseEvent{})
			if tc.want {
				assert.NotEqual(t, color.Transparent, r.background.FillColor)
			} else {
				assert.Equal(t, color.Transparent, r.background.FillColor)
			}
		})
	}
	t.Run("clears on mouse out", func(t *testing.T) {
		test.NewTempApp(t)
		w := NewIconButton(theme.HomeIcon(), func() {})
		r := test.TempWidgetRenderer(t, w).(*iconButtonRenderer)
		w.MouseIn(&desktop.MouseEvent{})
		w.MouseOut()
		assert.Equal(t, color.Transparent, r.background.FillColor)
	})
}

func TestIconButton_LayoutKeepsIconSizeWhenStretched(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	w := NewIconButton(theme.HomeIcon(), nil)
	r := test.TempWidgetRenderer(t, w).(*iconButtonRenderer)
	minSize := w.MinSize()

	w.Resize(fyne.NewSize(100, 60))

	th := w.Theme()
	iconSize := th.Size(theme.SizeNameInlineIcon)
	assert.Equal(t, fyne.NewSquareSize(iconSize), r.icon.Size())
	assert.Equal(t, fyne.NewPos((100-iconSize)/2, (60-iconSize)/2), r.icon.Position())
	assert.Equal(t, minSize, r.background.Size())
	assert.Equal(t, fyne.NewPos((100-minSize.Width)/2, (60-minSize.Height)/2), r.background.Position())
}
