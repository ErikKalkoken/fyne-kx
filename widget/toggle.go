package widget

import (
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO: Add disabled feature

// Toggle is a widget implementing a digital switch with two mutually exclusive states: on/off.
type Toggle struct {
	widget.BaseWidget
	OnChanged func(on bool)

	mu      sync.RWMutex // own property lock
	On      bool
	hovered bool
}

var _ fyne.Tappable = (*Toggle)(nil)
var _ desktop.Hoverable = (*Toggle)(nil)

var u float32 = theme.TextSize() // used for sizing the UI: 4u x 2u

// NewToggle returns a new [Toggle] instance.
func NewToggle(changed func(on bool)) *Toggle {
	w := &Toggle{
		OnChanged: changed,
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetState sets the state for a toggle.
func (w *Toggle) SetState(on bool) {
	func() {
		w.mu.Lock()
		defer w.mu.Unlock()
		if on == w.On {
			return
		}
		w.On = on
	}()
	if w.OnChanged != nil {
		w.OnChanged(on)
	}
	w.Refresh()
}

// Tapped is called when a pointer tapped event is captured and triggers any change handler
func (w *Toggle) Tapped(_ *fyne.PointEvent) {
	w.SetState(!w.On)
}

func (w *Toggle) TappedSecondary(_ *fyne.PointEvent) {
}

// Cursor returns the cursor type of this widget
func (w *Toggle) Cursor() desktop.Cursor {
	if w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MinSize returns the size that this widget should not shrink below
func (w *Toggle) MinSize() fyne.Size {
	w.ExtendBaseWidget(w)
	return w.BaseWidget.MinSize()
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *Toggle) MouseIn(e *desktop.MouseEvent) {
	w.hovered = true
	w.Refresh()
}

func (w *Toggle) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *Toggle) MouseOut() {
	w.hovered = false
	w.Refresh()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer.
func (w *Toggle) CreateRenderer() fyne.WidgetRenderer {
	w.ExtendBaseWidget(w)
	bg := theme.Color(theme.ColorNameDisabled)
	fg := theme.Color(theme.ColorNameForeground)
	w.mu.RLock()
	defer w.mu.RUnlock()
	r := &toogleRenderer{
		bgLeft:   canvas.NewCircle(bg),
		bgMiddle: canvas.NewRectangle(bg),
		bgRight:  canvas.NewCircle(bg),
		pin:      canvas.NewCircle(fg),
		toggle:   w,
	}
	r.updateState()
	return r
}

type toogleRenderer struct {
	bgLeft   *canvas.Circle
	bgMiddle *canvas.Rectangle
	bgRight  *canvas.Circle
	pin      *canvas.Circle
	toggle   *Toggle
}

func (r *toogleRenderer) Destroy() {
}

func (r *toogleRenderer) MinSize() (size fyne.Size) {
	size = fyne.Size{Width: 3.5 * u, Height: 2.0 * u}
	return
}

func (r *toogleRenderer) Layout(size fyne.Size) {
	r.bgLeft.Position1 = fyne.NewPos(0, 0)
	r.bgLeft.Position2 = fyne.NewPos(2*u, 2*u)
	r.bgRight.Position1 = fyne.NewPos(1.5*u, 0)
	r.bgRight.Position2 = fyne.NewPos(3.5*u, 2*u)
	r.bgMiddle.Move(fyne.NewPos(1*u, 0))
	r.bgMiddle.Resize(fyne.NewSize(1.5*u, 2*u))
	r.updateState()
}

// updateState updates the rendered toggle based on it's current state.
func (r *toogleRenderer) updateState() {
	border := theme.SelectionRadiusSize() / 2
	var x float32
	if r.toggle.On {
		x = 1.5 * u
	}
	r.pin.Position1 = fyne.NewPos(border+x, border)
	r.pin.Position2 = fyne.NewPos(2*u-1.5*border+x, 2*u-1.5*border)
	r.pin.Refresh()

	var bg color.Color
	if r.toggle.On {
		bg = theme.Color(theme.ColorNamePrimary)
	} else {
		bg = theme.Color(theme.ColorNameDisabled)
	}
	r.bgLeft.FillColor = bg
	r.bgLeft.Refresh()
	r.bgRight.FillColor = bg
	r.bgRight.Refresh()
	r.bgMiddle.FillColor = bg
	r.bgMiddle.Refresh()
}

func (r *toogleRenderer) Refresh() {
	func() {
		r.toggle.mu.RLock()
		defer r.toggle.mu.RUnlock()
		r.updateState()

	}()
	canvas.Refresh(r.toggle)
}

func (r *toogleRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bgLeft, r.bgRight, r.bgMiddle, r.pin}
}
