package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"

	"github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestRingActivity_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewRingActivity()
	w := test.NewWindow(a)
	defer w.Close()

	test.AssertImageMatches(t, "ringactivity/created.png", w.Canvas().Capture())
}

func TestRingActivity_CanStart(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewRingActivity()
	w := test.NewWindow(a)
	defer w.Close()

	a.Start()

	test.AssertImageMatches(t, "ringactivity/started.png", w.Canvas().Capture())
}

func TestRingActivity_Start_TriggersRepaint(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewRingActivity()
	w := test.NewWindow(a)
	defer w.Close()

	before := w.Canvas().Capture()
	a.Start()
	after := w.Canvas().Capture()

	assert.NotEqual(t, before, after)
}

func TestRingActivity_Stop_NoFurtherChange(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewRingActivity()
	w := test.NewWindow(a)
	defer w.Close()

	a.Start()
	a.Stop()

	before := w.Canvas().Capture()
	after := w.Canvas().Capture()

	assert.Equal(t, before, after)
}

func TestRingActivity_WithColorName(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewRingActivity()
	a.SetColorName(theme.ColorNamePrimary)
	w := test.NewWindow(a)
	defer w.Close()

	a.Start()

	test.AssertImageMatches(t, "ringactivity/started_primary.png", w.Canvas().Capture())
}

func TestRingActivity_Resized(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	a := widget.NewRingActivity()
	w := test.NewWindow(a)
	defer w.Close()
	w.Resize(fyne.NewSize(80, 80))

	a.Start()

	test.AssertImageMatches(t, "ringactivity/started_large.png", w.Canvas().Capture())
}
