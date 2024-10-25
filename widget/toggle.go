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

// toggle theme params
const (
	toggleBaseUnitSize    = theme.SizeNameText
	toggleSizeFocusBorder = theme.SizeNamePadding
	toggleSizePinBorder   = theme.SizeNameInputBorder

	toggleColorBackgroundOff    = theme.ColorNameInputBorder
	toggleColorBackgroundOn     = theme.ColorNamePrimary
	toggleColorPinDisabledDark  = theme.ColorNameBackground
	toggleColorPinDisabledLight = theme.ColorNameForeground
	toggleColorPinEnabledDark   = theme.ColorNameForeground
	toggleColorPinEnabledLight  = theme.ColorNameBackground
	toggleColorPinFocused       = theme.ColorNameFocus

	toggleScale = 1.75
)

// Toggle is a widget implementing a digital switch with two mutually exclusive states: on/off.
type Toggle struct {
	widget.DisableableWidget
	OnChanged func(on bool)

	focused bool
	hovered bool

	mu sync.RWMutex // own property lock
	On bool
}

var _ fyne.Widget = (*Toggle)(nil)
var _ fyne.Tappable = (*Toggle)(nil)
var _ fyne.Focusable = (*Toggle)(nil)
var _ desktop.Hoverable = (*Toggle)(nil)
var _ fyne.Disableable = (*Toggle)(nil)

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

// FocusGained is called when the Check has been given focus.
func (w *Toggle) FocusGained() {
	if w.Disabled() {
		return
	}
	w.focused = true
	w.Refresh()
}

// FocusLost is called when the Check has had focus removed.
func (w *Toggle) FocusLost() {
	w.focused = false
	w.Refresh()
}

// TypedRune receives text input events when the Check is focused.
func (w *Toggle) TypedRune(r rune) {
	if w.Disabled() {
		return
	}
	if r == ' ' {
		w.SetState(!w.On)
	}
}

// TypedKey receives key input events when the Check is focused.
func (w *Toggle) TypedKey(key *fyne.KeyEvent) {}

// Tapped is called when a pointer tapped event is captured and triggers any change handler
func (w *Toggle) Tapped(_ *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
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
	if w.Disabled() {
		return
	}
	w.hovered = true
	w.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
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
	th := w.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	w.ExtendBaseWidget(w)
	bg := th.Color(toggleColorBackgroundOff, v)
	w.mu.RLock()
	defer w.mu.RUnlock()
	r := &toogleRenderer{
		bgLeft:      canvas.NewCircle(bg),
		bgMiddle:    canvas.NewRectangle(bg),
		bgRight:     canvas.NewCircle(bg),
		pinLeft:     canvas.NewCircle(color.Transparent),
		pinRight:    canvas.NewCircle(color.Transparent),
		shadowLeft:  canvas.NewCircle(color.Transparent),
		shadowRight: canvas.NewCircle(color.Transparent),
		toggle:      w,
	}
	r.updateToggle()
	return r
}

var _ fyne.WidgetRenderer = (*toogleRenderer)(nil)

// toogleRenderer represents the renderer for the Toggle widget.
type toogleRenderer struct {
	bgLeft      *canvas.Circle
	bgMiddle    *canvas.Rectangle
	bgRight     *canvas.Circle
	pinLeft     *canvas.Circle
	pinRight    *canvas.Circle
	shadowLeft  *canvas.Circle
	shadowRight *canvas.Circle
	toggle      *Toggle
}

func (r *toogleRenderer) themeBase() (float32, fyne.Theme) {
	th := r.toggle.Theme()
	return th.Size(toggleBaseUnitSize) * toggleScale, th
}

func (r *toogleRenderer) Destroy() {
}

// MinSize returns the minimum size of the widget that is rendered by this renderer.
func (r *toogleRenderer) MinSize() (size fyne.Size) {
	u, th := r.themeBase()
	innerPadding := th.Size(theme.SizeNameInnerPadding) * 2
	size = fyne.NewSize(2*u+innerPadding, u+innerPadding)
	return
}

// Layout lays out the objects of this widget.
func (r *toogleRenderer) Layout(size fyne.Size) {
	u, th := r.themeBase()
	innerPadding := th.Size(theme.SizeNameInnerPadding)
	orig := fyne.NewPos(innerPadding, (size.Height-u)/2) // center vertically
	r.bgLeft.Position1 = orig
	r.bgLeft.Position2 = orig.AddXY(u, u)
	r.bgRight.Position1 = orig.AddXY(u, 0)
	r.bgRight.Position2 = orig.AddXY(2*u, u)
	r.bgMiddle.Move(orig.AddXY(0.5*u, 0))
	r.bgMiddle.Resize(fyne.NewSquareSize(u))

	border1 := th.Size(toggleSizePinBorder)
	r.pinLeft.Position1 = orig.AddXY(border1, border1)
	r.pinLeft.Position2 = orig.AddXY(u-2*border1, u-2*border1)

	r.pinRight.Position1 = orig.AddXY(border1+u, border1)
	r.pinRight.Position2 = orig.AddXY(2*u-2*border1, u-2*border1)

	border2 := th.Size(toggleSizeFocusBorder)
	r.shadowLeft.Position1 = orig.AddXY(0-border2, 0-border2)
	r.shadowLeft.Position2 = orig.AddXY(u+border2, u+border2)

	r.shadowRight.Position1 = orig.AddXY(u-border2, 0-border2)
	r.shadowRight.Position2 = orig.AddXY(2*u+border2, u+border2)

	// fmt.Printf("bgLeft: %+v - %+v\n", r.bgLeft.Position1, r.bgLeft.Position2)
	// fmt.Printf("bgRight: %+v - %+v\n", r.bgRight.Position1, r.bgRight.Position2)
	// fmt.Printf("pin: %+v - %+v\n", r.pin.Position1, r.pin.Position2)
	// fmt.Println()
}

// updateToggle updates the rendered toggle based on it's current state.
func (r *toogleRenderer) updateToggle() {
	_, th := r.themeBase()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	var pinColor color.Color
	if r.toggle.Disabled() {
		if v == theme.VariantLight {
			pinColor = th.Color(toggleColorPinDisabledLight, v)
		} else {
			pinColor = th.Color(toggleColorPinDisabledDark, v)
		}
	} else {
		if v == theme.VariantLight {
			pinColor = th.Color(toggleColorPinEnabledLight, v)
		} else {
			pinColor = th.Color(toggleColorPinEnabledDark, v)
		}
	}
	var focusColor color.Color
	if r.toggle.focused {
		focusColor = th.Color(toggleColorPinFocused, v)
	} else {
		focusColor = color.Transparent
	}

	if r.toggle.On {
		r.pinLeft.FillColor = color.Transparent
		r.shadowLeft.FillColor = color.Transparent
		r.pinRight.FillColor = pinColor
		r.shadowRight.FillColor = focusColor
	} else {
		r.pinLeft.FillColor = pinColor
		r.shadowLeft.FillColor = focusColor
		r.pinRight.FillColor = color.Transparent
		r.shadowRight.FillColor = color.Transparent
	}
	r.pinLeft.Refresh()
	r.shadowLeft.Refresh()
	r.pinRight.Refresh()
	r.shadowRight.Refresh()

	var bg color.Color
	if r.toggle.On {
		bg = th.Color(toggleColorBackgroundOn, v)
	} else {
		bg = th.Color(toggleColorBackgroundOff, v)
	}
	r.bgLeft.FillColor = bg
	r.bgLeft.Refresh()
	r.bgRight.FillColor = bg
	r.bgRight.Refresh()
	r.bgMiddle.FillColor = bg
	r.bgMiddle.Refresh()
}

// Refresh is called if the widget has updated and needs to be redrawn.
func (r *toogleRenderer) Refresh() {
	func() {
		r.toggle.mu.RLock()
		defer r.toggle.mu.RUnlock()
		r.updateToggle()
	}()
}

// Objects returns the objects that should be rendered.
func (r *toogleRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bgLeft, r.bgRight, r.bgMiddle, r.pinLeft, r.pinRight, r.shadowLeft, r.shadowRight}
}
