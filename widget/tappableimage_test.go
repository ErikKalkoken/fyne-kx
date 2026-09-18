package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"

	"github.com/ErikKalkoken/fyne-kx/widget"
)

func TestTappableImage_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	image.SetFillMode(canvas.ImageFillContain)
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	test.AssertImageMatches(t, "tappableimage/default.png", w.Canvas().Capture())
}

func TestTappableImage_CanSetResource(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	image.SetFillMode(canvas.ImageFillContain)
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	image.SetResource(theme.ComputerIcon())

	test.AssertImageMatches(t, "tappableimage/set_resource.png", w.Canvas().Capture())
}

func TestTappableImage_CanSetResourceFieldDirectly(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	image.SetFillMode(canvas.ImageFillContain)
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	// Assigning the exported Resource field directly, Fyne-widget style,
	// must render identically to calling SetResource.
	image.Resource = theme.ComputerIcon()
	image.Refresh()

	test.AssertImageMatches(t, "tappableimage/set_resource.png", w.Canvas().Capture())
}

func TestTappableImage_CanSetFillModeFieldDirectly(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	test.AssertImageMatches(t, "tappableimage/default.png", w.Canvas().Capture())
}

func TestTappableImage_CanTap(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	image := widget.NewTappableImage(theme.HomeIcon(), func() {
		tapped = true
	})
	image.SetFillMode(canvas.ImageFillContain)
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	test.Tap(image)
	assert.True(t, tapped)
}

func TestTappableImage_IgnoreTapWhenNoCallback(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	image.SetFillMode(canvas.ImageFillContain)
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	test.Tap(image)
}

func TestTappableImage_CanDisable(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	image.SetFillMode(canvas.ImageFillContain)
	image.SetMinSize(fyne.NewSquareSize(50))
	w := test.NewWindow(image)
	defer w.Close()

	image.Disable()

	assert.True(t, image.Disabled())
	test.AssertImageMatches(t, "tappableimage/disabled.png", w.Canvas().Capture())
}

func TestTappableImage_DisabledIgnoresTap(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	image := widget.NewTappableImage(theme.HomeIcon(), func() {
		tapped = true
	})
	w := test.NewWindow(image)
	defer w.Close()
	image.Disable()

	test.Tap(image)

	assert.False(t, tapped, "a disabled TappableImage must not fire its tap callback")
}

func TestTappableImage_WithMenu_NilMenuLeavesTapNoOp(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), nil)

	assert.NotNil(t, image)
	assert.Nil(t, image.OnTapped, "a misconfigured menu must not wire up a tap handler")
}

func TestTappableImage_WithMenu_TapDoesNothingWhenMenuEmpty(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), fyne.NewMenu(""))
	w := test.NewWindow(image)
	defer w.Close()

	assert.NotPanics(t, func() { test.Tap(image) })
}

func TestTappableImage_SetMenuItemsDoesNothingWithoutMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImage(theme.HomeIcon(), nil)
	w := test.NewWindow(image)
	defer w.Close()

	assert.NotPanics(t, func() {
		image.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("new", nil)})
	})
}

func TestTappableImage_SetMenuItemsReplacesItems(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	menu := fyne.NewMenu("", fyne.NewMenuItem("old", nil))
	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), menu)
	w := test.NewWindow(image)
	defer w.Close()

	newItems := []*fyne.MenuItem{fyne.NewMenuItem("new", nil)}
	image.SetMenuItems(newItems)

	assert.Equal(t, newItems, menu.Items)
}
