package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"

	"github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestSpinner_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewSpinner()
	w := test.NewWindow(a)
	defer w.Close()

	test.AssertImageMatches(t, "spinner/created.png", w.Canvas().Capture())
}

func TestSpinner_CanStart(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewSpinner()
	w := test.NewWindow(a)
	defer w.Close()

	a.Start()

	test.AssertImageMatches(t, "spinner/started.png", w.Canvas().Capture())
}

func TestSpinner_Start_TriggersRepaint(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewSpinner()
	w := test.NewWindow(a)
	defer w.Close()

	before := w.Canvas().Capture()
	a.Start()
	after := w.Canvas().Capture()

	assert.NotEqual(t, before, after)
}

func TestSpinner_Stop_NoFurtherChange(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewSpinner()
	w := test.NewWindow(a)
	defer w.Close()

	a.Start()
	a.Stop()

	before := w.Canvas().Capture()
	after := w.Canvas().Capture()

	assert.Equal(t, before, after)
}

func TestSpinner_WithColorName(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewSpinner()
	a.SetColorName(theme.ColorNamePrimary)
	w := test.NewWindow(a)
	defer w.Close()

	a.Start()

	test.AssertImageMatches(t, "spinner/started_primary.png", w.Canvas().Capture())
}

func TestSpinner_Resized(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewSpinner()
	w := test.NewWindow(a)
	defer w.Close()
	w.Resize(fyne.NewSize(80, 80))

	a.Start()

	test.AssertImageMatches(t, "spinner/started_large.png", w.Canvas().Capture())
}
