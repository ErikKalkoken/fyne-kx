package widget_test

import (
	"testing"

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
