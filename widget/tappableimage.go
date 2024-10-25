package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// TappableImage is widget which shows an image and runs a function when tapped.
type TappableImage struct {
	widget.BaseWidget
	image *canvas.Image

	// The function that is called when the label is tapped.
	OnTapped func()

	hovered bool
}

var _ fyne.Tappable = (*TappableImage)(nil)
var _ desktop.Hoverable = (*TappableImage)(nil)

// NewTappableImage returns a new instance of a [TappableImage] widget.
func NewTappableImage(res fyne.Resource, tapped func()) *TappableImage {
	ti := &TappableImage{OnTapped: tapped, image: canvas.NewImageFromResource(res)}
	ti.ExtendBaseWidget(ti)
	return ti
}

// SetFillMode sets the fill mode of the image.
func (w *TappableImage) SetFillMode(fillMode canvas.ImageFill) {
	w.image.FillMode = fillMode
}

// SetMinSize sets the minimum size of the image.
func (w *TappableImage) SetMinSize(size fyne.Size) {
	w.image.SetMinSize(size)
}

func (w *TappableImage) Tapped(_ *fyne.PointEvent) {
	if w.OnTapped != nil {
		w.OnTapped()
	}
}

func (w *TappableImage) TappedSecondary(_ *fyne.PointEvent) {
}

// Cursor returns the cursor type of this widget
func (w *TappableImage) Cursor() desktop.Cursor {
	if w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *TappableImage) MouseIn(e *desktop.MouseEvent) {
	w.hovered = true
}

func (w *TappableImage) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *TappableImage) MouseOut() {
	w.hovered = false
}

func (w *TappableImage) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewPadded(w.image))
}
