package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Switch is a widget representing a digital switch with two mutually exclusive states: on/off.
type Switch struct {
	widget.DisableableWidget

	On        bool
	OnChanged func(on bool)

	focused bool
	hovered bool

	// graphicPosition and graphicSize are the cached hitbox of the switch graphic
	graphicPosition fyne.Position
	graphicSize     fyne.Size
}

var _ desktop.Hoverable = (*Switch)(nil)
var _ fyne.Disableable = (*Switch)(nil)
var _ fyne.Focusable = (*Switch)(nil)
var _ fyne.Tappable = (*Switch)(nil)
var _ fyne.Widget = (*Switch)(nil)

// NewSwitch returns a new [Switch] instance.
func NewSwitch(changed func(on bool)) *Switch {
	w := &Switch{
		OnChanged: changed,
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetOn sets the state for a switch.
func (w *Switch) SetOn(on bool) {
	if on == w.On {
		return
	}
	w.On = on
	if w.OnChanged != nil {
		w.OnChanged(on)
	}
	w.Refresh()
}

func (w *Switch) FocusGained() {
	if w.Disabled() {
		return
	}
	w.focused = true
	w.Refresh()
}

func (w *Switch) FocusLost() {
	w.focused = false
	w.Refresh()
}

func (w *Switch) TypedRune(r rune) {
	if w.Disabled() {
		return
	}
	if r == ' ' {
		w.SetOn(!w.On)
	}
}

func (w *Switch) TypedKey(key *fyne.KeyEvent) {}

func (w *Switch) Tapped(pe *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	if !w.graphicContains(pe.Position) {
		return
	}
	w.SetOn(!w.On)
}

func (w *Switch) TappedSecondary(_ *fyne.PointEvent) {}

func (w *Switch) Cursor() desktop.Cursor {
	if w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

func (w *Switch) MouseIn(me *desktop.MouseEvent) {
	w.MouseMoved(me)
}

func (w *Switch) MouseMoved(me *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.updateHover(&me.Position)
}

func (w *Switch) MouseOut() {
	w.updateHover(nil)
}

// updateHover sets hover based on pos; nil (from MouseOut) always clears it.
func (w *Switch) updateHover(pos *fyne.Position) {
	oldHovered := w.hovered
	w.hovered = pos != nil && w.graphicContains(*pos)
	if oldHovered != w.hovered {
		w.Refresh()
	}
}

// graphicContains reports whether pos falls within the switch graphic's hitbox.
func (w *Switch) graphicContains(pos fyne.Position) bool {
	orig := w.graphicPosition
	size := w.graphicSize
	if size.IsZero() {
		size = w.Size() // renderer hasn't laid out yet; fall back to full widget
	}

	return pos.X >= orig.X && pos.X <= orig.X+size.Width &&
		pos.Y >= orig.Y && pos.Y <= orig.Y+size.Height
}

func (w *Switch) CreateRenderer() fyne.WidgetRenderer {
	w.ExtendBaseWidget(w)
	track := canvas.NewRectangle(color.Transparent)
	track.CornerRadius = 7
	r := &switchRenderer{
		focus:  canvas.NewCircle(color.Transparent),
		shadow: canvas.NewCircle(color.Transparent),
		thumb:  canvas.NewCircle(color.Transparent),
		track:  track,
		widget: w,
	}
	r.updateState()
	return r
}

var _ fyne.WidgetRenderer = (*switchRenderer)(nil)

// switch measurements
const (
	switchFocusHeight = 30
	switchHeight      = 20
	switchInnerHeight = 14
	switchWidth       = 36
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

func (r *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.track, r.focus, r.shadow, r.thumb}
}

func (r *switchRenderer) Destroy() {}

func (r *switchRenderer) MinSize() fyne.Size {
	th := r.widget.Theme()
	p := th.Size(theme.SizeNameInnerPadding)
	return fyne.NewSize(switchWidth+2*p, switchHeight+2*p)
}

func (r *switchRenderer) Layout(size fyne.Size) {
	th := r.widget.Theme()
	p := th.Size(theme.SizeNameInnerPadding)
	r.orig = fyne.NewPos(p, size.Height/2-switchHeight/2) // center vertically
	r.track.Move(r.orig.AddXY(0, (switchHeight-switchInnerHeight)/2))
	r.track.Resize(fyne.NewSize(switchWidth, switchInnerHeight))
	r.updateThumbPosition()

	// cache the graphic's hitbox for Tapped/MouseIn/MouseMoved
	r.widget.graphicPosition = r.orig
	r.widget.graphicSize = fyne.NewSize(switchWidth, switchHeight)
}

func (r *switchRenderer) updateThumbPosition() {
	focusOffset := (switchFocusHeight - switchHeight) / float32(2)
	const delta = 1
	if r.widget.On {
		r.thumb.Position1 = r.orig.AddXY(switchWidth-switchHeight, 0)
		r.thumb.Position2 = r.thumb.Position1.AddXY(switchHeight, switchHeight)
	} else {
		r.thumb.Position1 = r.orig
		r.thumb.Position2 = r.thumb.Position1.AddXY(switchHeight, switchHeight)
	}
	r.shadow.Position1 = r.thumb.Position1.AddXY(-delta, delta)
	r.shadow.Position2 = r.thumb.Position2.AddXY(-delta, delta)
	r.focus.Position1 = r.thumb.Position1.AddXY(-focusOffset, -focusOffset)
	r.focus.Position2 = r.focus.Position1.AddXY(switchFocusHeight, switchFocusHeight)
}

func (r *switchRenderer) Refresh() {
	r.updateState()

	r.track.Refresh()
	r.focus.Refresh()
	r.shadow.Refresh()
	r.thumb.Refresh()
}

func (r *switchRenderer) updateState() {
	r.updateColors()
	r.updateThumbPosition()
}

func (r *switchRenderer) updateColors() {
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
	isDark := v == theme.VariantDark
	if isDark {
		colorModifierMode = modeDarker
		trackColorModifier = 0.5
		disabledModifier = 0.75
	} else {
		colorModifierMode = modeBrighter
		trackColorModifier = 0.5
		disabledModifier = 0.2
	}

	// colors of thumb/shadow/focus
	if r.widget.On {
		thumbOnColor := th.Color(theme.ColorNamePrimary, v)
		if r.widget.Disabled() {
			c := newModifiedColor(thumbOnColor, colorModifierMode, disabledModifier)
			r.thumb.FillColor = c
			r.track.FillColor = newModifiedColor(c, colorModifierMode, trackColorModifier)
		} else {
			r.thumb.FillColor = thumbOnColor
			r.track.FillColor = newModifiedColor(thumbOnColor, colorModifierMode, trackColorModifier)
			r.focus.FillColor = focusColor
		}
	} else {
		if r.widget.Disabled() {
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

	r.shadow.FillColor = th.Color(theme.ColorNameShadow, v)
}
