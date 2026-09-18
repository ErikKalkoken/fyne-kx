package widget_test

import (
	"testing"

	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	"github.com/ErikKalkoken/fyne-kx/widget"
)

func TestTappableLabel_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	icon := widget.NewTappableLabel("Test", nil)
	w := test.NewWindow(icon)
	defer w.Close()

	test.AssertImageMatches(t, "tappablelabel/default.png", w.Canvas().Capture())
}

func TestTappableLabel_CanTap(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	icon := widget.NewTappableLabel("Test", func() {
		tapped = true
	})
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
	assert.True(t, tapped)
}

func TestTappableLabel_IgnoreTapWhenNoCallback(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := widget.NewTappableLabel("Test", nil)
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
}

func TestTappableLabel_CanDisable(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	label := widget.NewTappableLabel("Test", nil)
	w := test.NewWindow(label)
	defer w.Close()

	label.Disable()

	assert.True(t, label.Disabled())
	test.AssertImageMatches(t, "tappablelabel/disabled.png", w.Canvas().Capture())
}

func TestTappableLabel_DisabledIgnoresTap(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	label := widget.NewTappableLabel("Test", func() {
		tapped = true
	})
	w := test.NewWindow(label)
	defer w.Close()
	label.Disable()

	test.Tap(label)

	assert.False(t, tapped, "a disabled TappableLabel must not fire its tap callback")
}

func TestTappableLabel_MouseInOutTogglesPointerCursor(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	label := widget.NewTappableLabel("Test", nil)
	w := test.NewWindow(label)
	defer w.Close()

	assert.Equal(t, desktop.DefaultCursor, label.Cursor())

	label.MouseIn(&desktop.MouseEvent{})
	assert.Equal(t, desktop.PointerCursor, label.Cursor())

	label.MouseOut()
	assert.Equal(t, desktop.DefaultCursor, label.Cursor())
}
