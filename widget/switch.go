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

	// On is the state of the switch.
	// Set it and call Refresh to change it without animation and without calling OnChanged.
	On bool

	// OnChanged is called when the state changes.
	OnChanged func(on bool)

	animateNext bool // whether to animate on a refresh
	focused     bool
	hovered     bool

	// graphicPosition and graphicSize are the cached hit-box of the switch graphic
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
// It animates the thumb and calls OnChanged.
func (w *Switch) SetOn(on bool) {
	if on == w.On {
		return
	}
	w.On = on
	// set before OnChanged, which may refresh the switch itself
	prev := w.animateNext // restore, not clear: a nested SetOn from OnChanged must keep the outer flag
	w.animateNext = true
	defer func() { w.animateNext = prev }()
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
	if !w.Disabled() && w.hovered {
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

// graphicContains reports whether pos falls within the switch graphic's hit-box.
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
	r.progress = r.targetProgress() // a renderer created by SetOn's refresh must not animate
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
	anim       *fyne.Animation // running animation, which controls the thumb position; nil when none
	animTarget float32         // progress the running animation heads for
	focus      *canvas.Circle
	orig       fyne.Position
	progress   float32 // decoupled from On so a reversed toggle can animate from mid-flight
	shadow     *canvas.Circle
	thumb      *canvas.Circle
	track      *canvas.Rectangle
	widget     *Switch
}

func (r *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.track, r.focus, r.shadow, r.thumb}
}

func (r *switchRenderer) Destroy() {
	r.stopAnimation()
}

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

	// cache the graphic's hit-box for Tapped/MouseIn/MouseMoved
	r.widget.graphicPosition = r.orig
	r.widget.graphicSize = fyne.NewSize(switchWidth, switchHeight)
}

func (r *switchRenderer) updateThumbPosition() {
	focusOffset := (switchFocusHeight - switchHeight) / float32(2)
	const delta = 1

	offX := r.orig.X
	onX := r.orig.X + switchWidth - switchHeight
	x := offX + (onX-offX)*r.progress

	r.thumb.Position1 = fyne.NewPos(x, r.orig.Y)
	r.thumb.Position2 = r.thumb.Position1.AddXY(switchHeight, switchHeight)

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
	target := r.targetProgress()
	switch {
	case r.anim != nil && r.animTarget == target:
		// the running animation already heads for the current state
	case r.widget.animateNext && r.progress != target:
		r.animateThumb()
	default:
		r.stopAnimation() // On changed without SetOn mid-animation, so jump
		r.progress = target
	}
	r.updateThumbPosition()
}

func (r *switchRenderer) targetProgress() float32 {
	if r.widget.On {
		return 1
	}
	return 0
}

func (r *switchRenderer) stopAnimation() {
	if r.anim != nil {
		r.anim.Stop()
		r.anim = nil
	}
}

// animateThumb animates the thumb from its current position to the state of the switch.
// Stops any animation in flight first, so retapping mid-toggle reverses smoothly
// from wherever the thumb is instead of jumping.
func (r *switchRenderer) animateThumb() {
	r.stopAnimation()
	to := r.targetProgress()
	if !fyne.CurrentApp().Settings().ShowAnimations() {
		r.progress = to
		return
	}

	from := r.progress
	var a *fyne.Animation
	a = fyne.NewAnimation(canvas.DurationShort, func(done float32) {
		if r.anim != a {
			return // replaced by a newer animation
		}
		r.progress = from + (to-from)*done
		if done == 1 {
			r.anim = nil
		}
		r.updateThumbPosition()
		canvas.Refresh(r.thumb)
		canvas.Refresh(r.shadow)
		canvas.Refresh(r.focus)
	})
	a.Curve = fyne.AnimationEaseInOut
	r.anim = a // before Start: the test driver ticks to done inside it
	r.animTarget = to
	a.Start()
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
