package widget_test

import (
	"testing"

	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"

	"github.com/ErikKalkoken/fyne-kx/widget"
)

func TestTappableIcon_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	icon := widget.NewTappableIcon(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	test.AssertImageMatches(t, "tappableicon/default.png", w.Canvas().Capture())
}
func TestTappableIcon_CanTap(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	icon := widget.NewTappableIcon(theme.HomeIcon(), func() {
		tapped = true
	})
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
	assert.True(t, tapped)
}

func TestTappableIcon_IgnoreTapWhenNoCallback(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := widget.NewTappableIcon(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
}

func TestTappableIcon_CanDisable(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := widget.NewTappableIcon(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	icon.Disable()

	assert.True(t, icon.Disabled())
	test.AssertImageMatches(t, "tappableicon/disabled.png", w.Canvas().Capture())
}

func TestTappableIcon_DisabledIgnoresTap(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	icon := widget.NewTappableIcon(theme.HomeIcon(), func() {
		tapped = true
	})
	w := test.NewWindow(icon)
	defer w.Close()
	icon.Disable()

	test.Tap(icon)

	assert.False(t, tapped, "a disabled TappableIcon must not fire its tap callback")
}

func TestTappableIcon_MouseInOutTogglesPointerCursor(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := widget.NewTappableIcon(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	assert.Equal(t, desktop.DefaultCursor, icon.Cursor())

	icon.MouseIn(&desktop.MouseEvent{})
	assert.Equal(t, desktop.PointerCursor, icon.Cursor())

	icon.MouseOut()
	assert.Equal(t, desktop.DefaultCursor, icon.Cursor())
}
