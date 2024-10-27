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

// - [ ] Review disabled state
// - [ ] Review focused state
// - Animation?

// Switch is a widget implementing a digital switch with two mutually exclusive states: on/off.
// WIP - DO NOT USE
type Switch struct {
	widget.DisableableWidget
	OnChanged func(on bool)

	focused bool
	hovered bool

	mu sync.RWMutex // own property lock
	On bool
}

var _ fyne.Widget = (*Switch)(nil)
var _ fyne.Tappable = (*Switch)(nil)
var _ fyne.Focusable = (*Switch)(nil)
var _ desktop.Hoverable = (*Switch)(nil)
var _ fyne.Disableable = (*Switch)(nil)

// NewSwitch returns a new [Switch] instance.
func NewSwitch(changed func(on bool)) *Switch {
	w := &Switch{
		OnChanged: changed,
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetState sets the state for a switch.
func (w *Switch) SetState(on bool) {
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
func (w *Switch) FocusGained() {
	if w.Disabled() {
		return
	}
	w.focused = true
	w.Refresh()
}

// FocusLost is called when the Check has had focus removed.
func (w *Switch) FocusLost() {
	w.focused = false
	w.Refresh()
}

// TypedRune receives text input events when the Check is focused.
func (w *Switch) TypedRune(r rune) {
	if w.Disabled() {
		return
	}
	if r == ' ' {
		w.SetState(!w.On)
	}
}

// TypedKey receives key input events when the Check is focused.
func (w *Switch) TypedKey(key *fyne.KeyEvent) {}

// Tapped is called when a pointer tapped event is captured and triggers any change handler
func (w *Switch) Tapped(_ *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	w.SetState(!w.On)
}

func (w *Switch) TappedSecondary(_ *fyne.PointEvent) {
}

// Cursor returns the cursor type of this widget
func (w *Switch) Cursor() desktop.Cursor {
	if w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MinSize returns the size that this widget should not shrink below
func (w *Switch) MinSize() fyne.Size {
	w.ExtendBaseWidget(w)
	return w.BaseWidget.MinSize()
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *Switch) MouseIn(e *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.hovered = true
	w.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (w *Switch) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *Switch) MouseOut() {
	w.hovered = false
	w.Refresh()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer.
func (w *Switch) CreateRenderer() fyne.WidgetRenderer {
	w.ExtendBaseWidget(w)
	w.mu.RLock()
	defer w.mu.RUnlock()
	track := canvas.NewRectangle(color.Transparent)
	track.CornerRadius = 15
	r := &switchRenderer{
		track:       track,
		handleLeft:  canvas.NewCircle(color.Transparent),
		handleRight: canvas.NewCircle(color.Transparent),
		focusLeft:   canvas.NewCircle(color.Transparent),
		focusRight:  canvas.NewCircle(color.Transparent),
		widget:      w,
	}
	r.updateSwitch()
	return r
}

var _ fyne.WidgetRenderer = (*switchRenderer)(nil)

// switch measurements
const (
	switchTotalWidth             = 52
	switchTotalHeight            = 32
	switchHandleHeightSelected   = 24
	switchHandleHeightUnselected = 16
	switchOutlineWidth           = 2
)

// switch theme params
const (
	switchBaseUnitSize    = theme.SizeNameText
	switchSizeFocusBorder = theme.SizeNamePadding
	switchSizePinBorder   = theme.SizeNameInputBorder

	switchOffTrackFill    = theme.ColorNameInputBackground
	switchOffTrackOutline = theme.ColorNameScrollBar

	switchOnTrackFill           = theme.ColorNamePrimary
	switchColorPinEnabledDark   = theme.ColorNameForeground
	switchColorPinEnabledLight  = theme.ColorNameBackground
	switchColorPinFocused       = theme.ColorNameFocus
	switchColorPinDisabledDark  = theme.ColorNameDisabledButton
	switchColorPinDisabledLight = theme.ColorNamePlaceHolder
)

// switchRenderer represents the renderer for the Switch widget.
type switchRenderer struct {
	track       *canvas.Rectangle
	handleLeft  *canvas.Circle
	handleRight *canvas.Circle
	focusLeft   *canvas.Circle
	focusRight  *canvas.Circle
	widget      *Switch
}

func (r *switchRenderer) Destroy() {
}

// MinSize returns the minimum size of the widget that is rendered by this renderer.
func (r *switchRenderer) MinSize() (size fyne.Size) {
	th := r.widget.Theme()
	innerPadding := th.Size(theme.SizeNameInnerPadding)
	size = fyne.NewSize(size.Width+2*innerPadding, switchTotalHeight+2*innerPadding)
	return
}

// Layout lays out the objects of this widget.
func (r *switchRenderer) Layout(size fyne.Size) {
	th := r.widget.Theme()
	innerPadding := th.Size(theme.SizeNameInnerPadding)
	orig := fyne.NewPos(innerPadding, size.Height/2-switchTotalHeight/2) // center vertically
	r.track.Move(orig)
	r.track.Resize(fyne.NewSize(switchTotalWidth, switchTotalHeight))

	// fmt.Printf("%+v\n", r.bgLeft.Position())
	// fmt.Printf("%+v\n", r.bgRight.Position())
	// fmt.Printf("%+v\n\n", r.bgMiddle.Position())

	var d1 float32 = (switchTotalHeight - switchHandleHeightUnselected) / 2
	r.handleLeft.Position1 = orig.AddXY(d1, d1)
	r.handleLeft.Position2 = orig.AddXY(switchTotalHeight-d1, switchTotalHeight-d1)

	var d2 float32 = (switchTotalHeight - switchHandleHeightSelected) / 2
	r.handleRight.Position1 = orig.AddXY(switchTotalWidth-switchTotalHeight+d2, d2)
	r.handleRight.Position2 = orig.AddXY(switchTotalWidth-d2, switchTotalHeight-d2)

	// border2 := th.Size(switchSizeFocusBorder)
	// r.shadowLeft.Position1 = orig.AddXY(0-border2, 0-border2)
	// r.shadowLeft.Position2 = orig.AddXY(u+border2, u+border2)

	// r.shadowRight.Position1 = orig.AddXY(u-border2, 0-border2)
	// r.shadowRight.Position2 = orig.AddXY(2*u+border2, u+border2)

	// fmt.Printf("bgLeft: %+v - %+v\n", r.bgLeft.Position1, r.bgLeft.Position2)
	// fmt.Printf("bgRight: %+v - %+v\n", r.bgRight.Position1, r.bgRight.Position2)
	// fmt.Printf("pin: %+v - %+v\n", r.pin.Position1, r.pin.Position2)
	// fmt.Println()
}

// updateSwitch updates the rendered switch based on it's current state.
func (r *switchRenderer) updateSwitch() {
	th := r.widget.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	var pinColor color.Color
	if r.widget.Disabled() {
		if v == theme.VariantLight {
			pinColor = th.Color(switchColorPinDisabledLight, v)
		} else {
			pinColor = th.Color(switchColorPinDisabledDark, v)
		}
	} else {
		if v == theme.VariantLight {
			pinColor = th.Color(switchColorPinEnabledLight, v)
		} else {
			pinColor = th.Color(switchColorPinEnabledDark, v)
		}
	}
	var focusColor color.Color
	if r.widget.focused {
		focusColor = th.Color(switchColorPinFocused, v)
	} else {
		focusColor = color.Transparent
	}

	if r.widget.On {
		r.track.FillColor = th.Color(switchOnTrackFill, v)
		r.track.StrokeWidth = 0
		r.handleLeft.FillColor = color.Transparent
		r.focusLeft.FillColor = color.Transparent
		r.handleRight.FillColor = pinColor
		r.focusRight.FillColor = focusColor
	} else {
		r.handleLeft.FillColor = th.Color(switchOffTrackOutline, v)
		r.focusLeft.FillColor = focusColor
		r.handleRight.FillColor = color.Transparent
		r.focusRight.FillColor = color.Transparent
		r.track.FillColor = th.Color(switchOffTrackFill, v)
		r.track.StrokeColor = th.Color(switchOffTrackOutline, v)
		r.track.StrokeWidth = switchOutlineWidth
	}
	r.handleLeft.Refresh()
	r.focusLeft.Refresh()
	r.handleRight.Refresh()
	r.focusRight.Refresh()
	r.track.Refresh()
}

// Refresh is called if the widget has updated and needs to be redrawn.
func (r *switchRenderer) Refresh() {
	func() {
		r.widget.mu.RLock()
		defer r.widget.mu.RUnlock()
		r.updateSwitch()
	}()
}

// Objects returns the objects that should be rendered.
func (r *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.track, r.handleLeft, r.handleRight, r.focusLeft, r.focusRight}
}
