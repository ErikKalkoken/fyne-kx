package widget

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

// newSizedTappableImage resizes a TappableImage much larger than the theme padding,
// so tests can prove the hover hitbox is limited to the image, not the padded container.
func newSizedTappableImage(t *testing.T, w, h float32) *TappableImage {
	t.Helper()
	test.NewApp()
	img := NewTappableImage(theme.HomeIcon(), nil)
	img.Resize(fyne.NewSize(w, h))
	return img
}

func TestTappableImage_Disable_FadesResourceRegardlessOfType(t *testing.T) {
	// iconQuestionmark32Png is a raster (PNG) resource, unlike the SVG icons
	// used elsewhere in this file. Translucency is applied to the rendered
	// image rather than the resource content, so it must fade a bitmap
	// resource exactly the same way as a vector one.
	test.NewApp()
	img := NewTappableImage(iconQuestionmark32Png, nil)

	img.Disable()
	assert.Equal(t, float64(disabledImageTranslucency), img.image.Translucency)

	img.Enable()
	assert.Equal(t, float64(0), img.image.Translucency)
}

func TestTappableImage_Disabled_CursorNeverShowsPointer(t *testing.T) {
	test.NewApp()
	img := NewTappableImage(theme.HomeIcon(), nil)
	img.hovered = true
	img.Disable()

	assert.Equal(t, desktop.DefaultCursor, img.Cursor())
}

func TestTappableImage_MouseMoved_InsidePaddingStrip_DoesNotHover(t *testing.T) {
	img := newSizedTappableImage(t, 300, 300)

	// The image is inset by theme.Padding() inside the padded container, so a point
	// just inside the widget's top-left corner falls in the padding, not the image.
	inPadding := fyne.NewPos(theme.Padding()/2, theme.Padding()/2)

	img.MouseMoved(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: inPadding}})

	assert.False(t, img.hovered, "a point within the padding strip must not be treated as hovering the image")
}

func TestTappableImage_MouseMoved_NearFarEdge_StillHovers(t *testing.T) {
	img := newSizedTappableImage(t, 300, 300)

	// The image's right/bottom edge sits at (widgetSize - padding), not at widgetSize.
	nearEdge := fyne.NewPos(300-theme.Padding()-1, 300-theme.Padding()-1)

	img.MouseMoved(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: nearEdge}})

	assert.True(t, img.hovered, "a point just inside the image's far edge must still be treated as hovering")
}

func TestTappableImage_MouseMoved_InsideImage_Hovers(t *testing.T) {
	img := newSizedTappableImage(t, 300, 300)

	center := fyne.NewPos(150, 150)

	img.MouseMoved(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: center}})

	assert.True(t, img.hovered, "a point well within the image bounds should set hover")
}
