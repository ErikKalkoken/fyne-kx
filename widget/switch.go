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
	minSize fyne.Size // cached for hover/top pos calcs

	mu sync.RWMutex // own property lock
	on bool
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
		if on == w.on {
			return
		}
		w.on = on
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
		w.SetState(!w.on)
	}
}

// TypedKey receives key input events when the Check is focused.
func (w *Switch) TypedKey(key *fyne.KeyEvent) {}

// Tapped is called when a pointer tapped event is captured and triggers any change handler
func (w *Switch) Tapped(pe *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	if !w.minSize.IsZero() &&
		(pe.Position.X > w.minSize.Width || pe.Position.Y > w.minSize.Height) {
		// tapped outside
		return
	}
	if !w.focused {
		if !fyne.CurrentDevice().IsMobile() {
			if c := fyne.CurrentApp().Driver().CanvasForObject(w); c != nil {
				c.Focus(w)
			}
		}
	}
	w.SetState(!w.on)
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
	w.minSize = w.BaseWidget.MinSize()
	return w.minSize
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *Switch) MouseIn(me *desktop.MouseEvent) {
	w.MouseMoved(me)
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (w *Switch) MouseMoved(me *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	oldHovered := w.hovered

	w.hovered = w.minSize.IsZero() ||
		(me.Position.X <= w.minSize.Width && me.Position.Y <= w.minSize.Height)

	if oldHovered != w.hovered {
		w.Refresh()
	}
}

func (w *Switch) MouseOut() {
	if w.hovered {
		w.hovered = false
		w.Refresh()
	}
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer.
func (w *Switch) CreateRenderer() fyne.WidgetRenderer {
	w.ExtendBaseWidget(w)
	w.mu.RLock()
	defer w.mu.RUnlock()
	track := canvas.NewRectangle(color.Transparent)
	track.CornerRadius = 7
	r := &switchRenderer{
		track:       track,
		handleLeft:  canvas.NewCircle(color.Transparent),
		handleRight: canvas.NewCircle(color.Transparent),
		focusLeft:   canvas.NewCircle(color.Transparent),
		focusRight:  canvas.NewCircle(color.Transparent),
		widget:      w,
	}
	r.refreshSwitch()
	return r
}

var _ fyne.WidgetRenderer = (*switchRenderer)(nil)

// switch measurements
const (
	switchWidth       = 36
	switchInnerHeight = 14
	switchHeight      = 20
	switchFocusHeight = 35
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
	size = fyne.NewSize(switchWidth+2*innerPadding, switchHeight+2*innerPadding)
	return
}

// Layout lays out the objects of this widget.
func (r *switchRenderer) Layout(size fyne.Size) {
	th := r.widget.Theme()
	innerPadding := th.Size(theme.SizeNameInnerPadding)
	orig := fyne.NewPos(innerPadding, size.Height/2-switchHeight/2) // center vertically
	r.track.Move(orig.AddXY(0, (switchHeight-switchInnerHeight)/2))
	r.track.Resize(fyne.NewSize(switchWidth, switchInnerHeight))

	r.handleLeft.Position1 = orig
	r.handleLeft.Position2 = r.handleLeft.Position1.AddXY(switchHeight, switchHeight)
	r.handleRight.Position1 = orig.AddXY(switchWidth-switchHeight, 0)
	r.handleRight.Position2 = r.handleRight.Position1.AddXY(switchHeight, switchHeight)

	d := (switchFocusHeight - switchHeight) / float32(2)
	r.focusLeft.Position1 = orig.AddXY(0-d, 0-d)
	r.focusLeft.Position2 = r.focusLeft.Position1.AddXY(switchFocusHeight, switchFocusHeight)
	r.focusRight.Position1 = orig.AddXY(switchWidth-switchHeight-d, 0-d)
	r.focusRight.Position2 = r.focusRight.Position1.AddXY(switchFocusHeight, switchFocusHeight)
}

// refreshSwitch refreshes the rendered switch for it's current state.
func (r *switchRenderer) refreshSwitch() {
	th := r.widget.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	// isDisabled := r.widget.Disabled()
	// isDark := fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark

	var focusColor color.Color
	if r.widget.focused {
		focusColor = th.Color(theme.ColorNameFocus, v)
	} else if r.widget.hovered {
		focusColor = th.Color(theme.ColorNameHover, v)
	} else {
		focusColor = color.Transparent
	}

	if r.widget.on {
		primaryColor := th.Color(theme.ColorNamePrimary, v)
		r.track.FillColor = newColorWithReducedIntensity(primaryColor, 0.5)
		r.handleLeft.FillColor = color.Transparent
		r.handleRight.FillColor = primaryColor
		r.focusLeft.FillColor = color.Transparent
		r.focusRight.FillColor = focusColor
	} else {
		r.track.FillColor = th.Color(theme.ColorNameInputBorder, v)
		r.handleLeft.FillColor = th.Color(theme.ColorNameForeground, v)
		r.handleRight.FillColor = color.Transparent
		r.focusLeft.FillColor = focusColor
		r.focusRight.FillColor = color.Transparent
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
		r.refreshSwitch()
	}()
}

// Objects returns the objects that should be rendered.
func (r *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.track, r.focusLeft, r.focusRight, r.handleLeft, r.handleRight}
}

type myColor struct {
	c color.Color
	t float32
}

// newColorWithReducedIntensity returns a new instance of color with modified intensity.
//
// The intensity value is expected between 0 and 1. Larger and smaller numbers will be truncated.
func newColorWithReducedIntensity(c color.Color, intensity float32) color.Color {
	if intensity < 0 {
		intensity = 0
	}
	if intensity > 1 {
		intensity = 1
	}
	return myColor{c: c, t: intensity}
}

func (mc myColor) RGBA() (r, g, b, a uint32) {
	r, g, b, a = mc.c.RGBA()
	r2 := float32(r) / 0xffff * mc.t
	g2 := float32(g) / 0xffff * mc.t
	b2 := float32(b) / 0xffff * mc.t
	r = uint32(r2 * 0xffff)
	g = uint32(g2 * 0xffff)
	b = uint32(b2 * 0xffff)
	return
}
