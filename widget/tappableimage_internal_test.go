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
