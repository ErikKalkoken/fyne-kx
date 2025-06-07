package widget_test

import (
	"testing"

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
