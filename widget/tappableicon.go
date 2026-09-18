package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TappableIcon is an icon widget, which runs a function when tapped.
type TappableIcon struct {
	widget.Icon

	// The function that is called when the icon is tapped.
	OnTapped func()

	hovered          bool
	disabled         bool
	resource         fyne.Resource
	resourceDisabled fyne.Resource
}

var _ fyne.Tappable = (*TappableIcon)(nil)
var _ desktop.Hoverable = (*TappableIcon)(nil)
var _ fyne.Disableable = (*TappableIcon)(nil)

// NewTappableIcon returns a new instance of a [TappableIcon] widget.
func NewTappableIcon(res fyne.Resource, tapped func()) *TappableIcon {
	w := &TappableIcon{OnTapped: tapped}
	w.ExtendBaseWidget(w)
	w.SetResource(res)
	return w
}

// SetResource updates the resource rendered in this icon widget.
func (w *TappableIcon) SetResource(res fyne.Resource) {
	w.resource = res
	if isResourceSVG(res) {
		w.resourceDisabled = theme.NewDisabledResource(res)
	} else {
		w.resourceDisabled = res
	}
	w.Refresh()
}

// Refresh triggers a redraw of the icon, applying the current disabled state.
func (w *TappableIcon) Refresh() {
	if w.disabled {
		w.Icon.Resource = w.resourceDisabled
	} else {
		w.Icon.Resource = w.resource
	}
	w.Icon.Refresh()
}

// Enable this widget, updating any style or features appropriately.
func (w *TappableIcon) Enable() {
	if !w.disabled {
		return
	}
	w.disabled = false
	w.Refresh()
}

// Disable this widget so that it cannot be interacted with.
func (w *TappableIcon) Disable() {
	if w.disabled {
		return
	}
	w.disabled = true
	w.Refresh()
}

// Disabled returns true if this widget is currently disabled.
func (w *TappableIcon) Disabled() bool {
	return w.disabled
}

func (w *TappableIcon) Tapped(_ *fyne.PointEvent) {
	if w.disabled {
		return
	}
	if w.OnTapped != nil {
		w.OnTapped()
	}
}

func (w *TappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}

// Cursor returns the cursor type of this widget
func (w *TappableIcon) Cursor() desktop.Cursor {
	if !w.disabled && w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *TappableIcon) MouseIn(e *desktop.MouseEvent) {
	if w.disabled {
		return
	}
	w.hovered = true
}

func (w *TappableIcon) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *TappableIcon) MouseOut() {
	w.hovered = false
}
