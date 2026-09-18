package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// TappableLabel is a variant of the Fyne Label which runs a function when tapped.
type TappableLabel struct {
	widget.Label

	// The function that is called when the label is tapped.
	OnTapped func()

	hovered  bool
	disabled bool
}

var _ fyne.Tappable = (*TappableLabel)(nil)
var _ desktop.Hoverable = (*TappableLabel)(nil)
var _ fyne.Disableable = (*TappableLabel)(nil)

// NewTappableLabel returns a new TappableLabel instance.
func NewTappableLabel(text string, tapped func()) *TappableLabel {
	w := &TappableLabel{OnTapped: tapped}
	w.ExtendBaseWidget(w)
	w.SetText(text)
	return w
}

// Enable this widget, updating any style or features appropriately.
//
// This resets [TappableLabel.Importance] to [widget.MediumImportance]; any
// other importance set before disabling is not restored.
func (w *TappableLabel) Enable() {
	if !w.disabled {
		return
	}
	w.disabled = false
	w.Importance = widget.MediumImportance
	w.Refresh()
}

// Disable this widget so that it cannot be interacted with.
func (w *TappableLabel) Disable() {
	if w.disabled {
		return
	}
	w.disabled = true
	w.Importance = widget.LowImportance
	w.Refresh()
}

// Disabled returns true if this widget is currently disabled.
func (w *TappableLabel) Disabled() bool {
	return w.disabled
}

func (w *TappableLabel) Tapped(_ *fyne.PointEvent) {
	if w.disabled {
		return
	}
	if w.OnTapped != nil {
		w.OnTapped()
	}
}

// Cursor returns the cursor type of this widget
func (w *TappableLabel) Cursor() desktop.Cursor {
	if !w.disabled && w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *TappableLabel) MouseIn(e *desktop.MouseEvent) {
	if w.disabled {
		return
	}
	w.hovered = true
}

func (w *TappableLabel) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *TappableLabel) MouseOut() {
	w.hovered = false
}
