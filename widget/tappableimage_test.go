package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
	w := test.NewWindow(container.NewCenter(image))
	defer w.Close()

	test.Tap(image)

	assert.Nil(t, w.Canvas().Overlays().Top())
}

func TestTappableImage_WithMenu_TapDoesNothingWhenMenuEmpty(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), fyne.NewMenu(""))
	w := test.NewWindow(container.NewCenter(image))
	defer w.Close()

	test.Tap(image)

	assert.Nil(t, w.Canvas().Overlays().Top())
}

func TestTappableImage_WithMenu_ShowsMenuWhenTapped(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(image))
	defer w.Close()

	test.Tap(image)

	assert.NotNil(t, w.Canvas().Overlays().Top(), "should show menu")
}

func TestTappableImage_SetMenuItemsAddsMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	image := widget.NewTappableImage(theme.HomeIcon(), func() {
		tapped = true
	})
	w := test.NewWindow(container.NewCenter(image))
	defer w.Close()

	image.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("new", nil)})
	test.Tap(image)

	assert.NotNil(t, w.Canvas().Overlays().Top(), "should show menu")
	assert.False(t, tapped, "should have cleared OnTapped")
}

func TestTappableImage_OnTappedTakesPrecedenceOverMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(image))
	defer w.Close()
	var tapped bool
	image.OnTapped = func() {
		tapped = true
	}

	test.Tap(image)

	assert.True(t, tapped)
	assert.Nil(t, w.Canvas().Overlays().Top(), "should not show menu")
}

func TestTappableImage_ShowMenuAgainWhenOnTappedCleared(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	image := widget.NewTappableImageWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(image))
	defer w.Close()
	image.OnTapped = func() {}
	image.OnTapped = nil

	test.Tap(image)

	assert.NotNil(t, w.Canvas().Overlays().Top(), "should show menu")
}

type embeddedTappableImage struct {
	widget.TappableImage
}

func newEmbeddedTappableImage() *embeddedTappableImage {
	w := &embeddedTappableImage{}
	w.Resource = theme.HomeIcon()
	w.ExtendBaseWidget(w)
	return w
}

func TestTappableImage_Embedded(t *testing.T) {
	t.Run("can tap", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		var tapped bool
		image := newEmbeddedTappableImage()
		image.OnTapped = func() {
			tapped = true
		}
		w := test.NewWindow(image)
		defer w.Close()

		test.Tap(image)

		assert.True(t, tapped)
	})
	t.Run("can show menu", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		image := newEmbeddedTappableImage()
		image.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("item", nil)})
		w := test.NewWindow(container.NewCenter(image))
		defer w.Close()

		test.Tap(image)

		assert.NotNil(t, w.Canvas().Overlays().Top(), "should show menu")
	})
	t.Run("shows no menu when disabled", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		image := newEmbeddedTappableImage()
		image.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("item", nil)})
		image.Disable()
		w := test.NewWindow(container.NewCenter(image))
		defer w.Close()

		test.Tap(image)

		assert.Nil(t, w.Canvas().Overlays().Top())
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
