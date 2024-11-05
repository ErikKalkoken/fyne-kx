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

// - Add animation?

// Switch is a widget representing a digital switch with two mutually exclusive states: on/off.
type Switch struct {
	widget.DisableableWidget
	On bool

	OnChanged func(on bool)

	focused bool
	hovered bool
	minSize fyne.Size    // cached for hover/top pos calcs
	mu      sync.RWMutex // property lock
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

// State return the state of a switch.
func (w *Switch) State() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.On
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
	th := w.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	track := canvas.NewRectangle(color.Transparent)
	track.CornerRadius = 7
	shadowColor := th.Color(theme.ColorNameShadow, v)
	r := &switchRenderer{
		track:  track,
		thumb:  canvas.NewCircle(color.Transparent),
		focus:  canvas.NewCircle(color.Transparent),
		shadow: canvas.NewCircle(shadowColor),
		widget: w,
	}
	r.Refresh()
	return r
}

var _ fyne.WidgetRenderer = (*switchRenderer)(nil)

// switch measurements
const (
	switchWidth       = 36
	switchInnerHeight = 14
	switchHeight      = 20
	switchFocusHeight = 30
)

// switchRenderer represents the renderer for the Switch widget.
type switchRenderer struct {
	focus  *canvas.Circle
	orig   fyne.Position
	shadow *canvas.Circle
	thumb  *canvas.Circle
	track  *canvas.Rectangle
	widget *Switch
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
	r.orig = fyne.NewPos(innerPadding, size.Height/2-switchHeight/2) // center vertically
	r.track.Move(r.orig.AddXY(0, (switchHeight-switchInnerHeight)/2))
	r.track.Resize(fyne.NewSize(switchWidth, switchInnerHeight))
	r.Refresh()
}

// refreshSwitch refreshes the switch's drawing for it's current state.
// Should be called with a read lock.
func (r *switchRenderer) refreshSwitch() {
	th := r.widget.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	// focus colors and state
	var focusColor color.Color
	if r.widget.focused {
		focusColor = th.Color(theme.ColorNameFocus, v)
		r.focus.Show()
	} else if r.widget.hovered {
		focusColor = th.Color(theme.ColorNameHover, v)
		r.focus.Show()
	} else {
		r.focus.Hide()
	}

	// theme dependent parameters
	var colorModifierMode modifiedColorMode
	var disabledModifier, trackColorModifier float32
	isDark := fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark
	if isDark {
		colorModifierMode = modeDarker
		trackColorModifier = 0.5
		disabledModifier = 0.75
	} else {
		colorModifierMode = modeBrighter
		trackColorModifier = 0.5
		disabledModifier = 0.2
	}

	// colors of all objects and position of thumb/shadow/focus
	focusOffset := (switchFocusHeight - switchHeight) / float32(2)
	const delta = 1
	isDisabled := r.widget.Disabled()
	if r.widget.On {
		r.thumb.Position1 = r.orig.AddXY(switchWidth-switchHeight, 0)
		r.thumb.Position2 = r.thumb.Position1.AddXY(switchHeight, switchHeight)
		thumbOnColor := th.Color(theme.ColorNamePrimary, v)
		if isDisabled {
			c := newModifiedColor(thumbOnColor, colorModifierMode, disabledModifier)
			r.thumb.FillColor = c
			r.track.FillColor = newModifiedColor(c, colorModifierMode, trackColorModifier)
		} else {
			r.thumb.FillColor = thumbOnColor
			r.track.FillColor = newModifiedColor(thumbOnColor, colorModifierMode, trackColorModifier)
			r.focus.FillColor = focusColor
		}
	} else {
		r.thumb.Position1 = r.orig
		r.thumb.Position2 = r.thumb.Position1.AddXY(switchHeight, switchHeight)
		if isDisabled {
			r.thumb.FillColor = th.Color(theme.ColorNameDisabled, v)
			r.track.FillColor = th.Color(theme.ColorNameDisabledButton, v)
		} else {
			if isDark {
				r.thumb.FillColor = th.Color(theme.ColorNameForeground, v)
			} else {
				r.thumb.FillColor = th.Color(theme.ColorNameButton, v)
			}
			r.track.FillColor = th.Color(theme.ColorNameInputBorder, v)
			r.focus.FillColor = focusColor
		}
	}
	r.shadow.Position1 = r.thumb.Position1.AddXY(-delta, delta)
	r.shadow.Position2 = r.thumb.Position2.AddXY(-delta, delta)
	r.focus.Position1 = r.thumb.Position1.AddXY(-focusOffset, -focusOffset)
	r.focus.Position2 = r.focus.Position1.AddXY(switchFocusHeight, switchFocusHeight)

	r.track.Refresh()
	r.focus.Refresh()
	r.shadow.Refresh()
	r.thumb.Refresh()
}

// Refresh is called if the widget has updated and needs to be redrawn.
func (r *switchRenderer) Refresh() {
	r.widget.mu.RLock()
	defer r.widget.mu.RUnlock()
	r.refreshSwitch()
}

// Objects returns the objects that should be rendered.
func (r *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.track, r.focus, r.shadow, r.thumb}
}
