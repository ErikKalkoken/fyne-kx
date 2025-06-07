package widget_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestSlider_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	slider := kxwidget.NewSlider(0, 10)
	w := test.NewWindow(slider)
	defer w.Close()

	assert.EqualValues(t, 0, slider.Value())
	test.AssertImageMatches(t, "slider/default.png", w.Canvas().Capture())
}

func TestSlider_CanSetValue(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	slider := kxwidget.NewSlider(0, 10)
	w := test.NewWindow(slider)
	defer w.Close()

	slider.SetValue(7)

	assert.EqualValues(t, 7, slider.Value())
	test.AssertImageMatches(t, "slider/set_value.png", w.Canvas().Capture())
}
