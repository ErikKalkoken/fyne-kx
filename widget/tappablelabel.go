package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// TappableLabel is a variant of the Fyne Label which runs a function when tapped.
type TappableLabel struct {
	*widget.Label

	// The function that is called when the label is tapped.
	OnTapped func()

	hovered bool
}

var _ fyne.Tappable = (*TappableLabel)(nil)
var _ desktop.Hoverable = (*TappableLabel)(nil)

// NewTappableLabel returns a new TappableLabel instance.
func NewTappableLabel(text string, tapped func()) *TappableLabel {
	w := &TappableLabel{OnTapped: tapped, Label: widget.NewLabel(text)}
	w.ExtendBaseWidget(w)
	return w
}

func (w *TappableLabel) Tapped(_ *fyne.PointEvent) {
	if w.OnTapped != nil {
		w.OnTapped()
	}
}

// Cursor returns the cursor type of this widget
func (w *TappableLabel) Cursor() desktop.Cursor {
	if w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *TappableLabel) MouseIn(e *desktop.MouseEvent) {
	w.hovered = true
}

func (w *TappableLabel) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *TappableLabel) MouseOut() {
	w.hovered = false
}
