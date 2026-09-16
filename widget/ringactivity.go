package widget

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	ringActivityAnimationDuration = time.Second
	ringActivityMinSweep          = float32(20)
	ringActivityMaxSweep          = float32(280)
	// ring width = (1-CutoutRatio) * radius, 30%
	ringActivityCutoutRatio = float32(0.70)
	// background spin shared by both edges, so neither is ever fully still
	ringActivityRotationPerCycle = float32(360)
	// distance each edge travels per inchworm half (see animate)
	ringActivityGrowRange = ringActivityMaxSweep - ringActivityMinSweep
	// completedCycles wraps at this value in animate() to keep it bounded
	// over long runs. 18 cycles * 620 (grow range + rotation) = 11160 = 31
	// full turns, so wrapping has no visual effect.
	ringActivityCycleWrap = float32(18)
)

var _ fyne.Widget = (*RingActivity)(nil)

// RingActivity is used to indicate that something is happening that should be
// waited for, or is in the background (depending on usage).
// It shows an animated ring, similar to a Material Design loading indicator.
type RingActivity struct {
	widget.BaseWidget

	// ColorName is the named theme color used to draw the ring. It defaults
	// to [theme.ColorNameForeground] when left empty.
	ColorName fyne.ThemeColorName

	started bool
}

// NewRingActivity returns a widget for indicating activity with an animated ring.
func NewRingActivity() *RingActivity {
	a := &RingActivity{}
	a.ExtendBaseWidget(a)
	return a
}

// SetColorName sets the named theme color used to draw the ring.
func (a *RingActivity) SetColorName(name fyne.ThemeColorName) {
	a.ColorName = name
	a.Refresh()
}

func (a *RingActivity) MinSize() fyne.Size {
	a.ExtendBaseWidget(a)
	return a.BaseWidget.MinSize()
}

// Start the activity indicator animation.
func (a *RingActivity) Start() {
	if a.started {
		return // already started
	}

	a.started = true

	a.Refresh()
}

// Stop the activity indicator animation.
func (a *RingActivity) Stop() {
	if !a.started {
		return // already stopped
	}

	a.started = false

	a.Refresh()
}

func (a *RingActivity) CreateRenderer() fyne.WidgetRenderer {
	arc := canvas.NewArc(0, 0, ringActivityCutoutRatio, color.Transparent)
	r := &ringActivityRenderer{arc: arc, parent: a}
	r.anim = &fyne.Animation{
		Duration:    ringActivityAnimationDuration,
		RepeatCount: fyne.AnimationRepeatForever,
		// animate() does its own per-phase easing below
		Curve: fyne.AnimationLinear,
		Tick:  r.animate,
	}
	r.updateColor()

	if a.started {
		r.start()
	}

	return r
}

var _ fyne.WidgetRenderer = (*ringActivityRenderer)(nil)

type ringActivityRenderer struct {
	anim   *fyne.Animation
	arc    *canvas.Arc
	parent *RingActivity

	wasStarted bool

	// completedCycles counts full cycles seen so far, for the continuous
	// spin and cycleAdvance in animate. Wrapped by ringActivityCycleWrap.
	completedCycles float32
	// lastDone is done from the previous tick, to detect a new cycle starting.
	lastDone float32
}

func (r *ringActivityRenderer) Destroy() {
	r.parent.started = false
	r.stop()
}

func (r *ringActivityRenderer) Layout(size fyne.Size) {
	rad := fyne.Min(size.Width, size.Height) / 2
	mid := fyne.NewPos(size.Width/2, size.Height/2)

	// Despite canvas.Arc's doc, both rendering backends treat its position
	// as the top-left of the bounding box (like canvas.Circle), not the
	// center - so offset by the radius, same as Activity does for its dots.
	r.arc.Resize(fyne.NewSquareSize(rad * 2))
	r.arc.Move(mid.Subtract(fyne.NewSquareOffsetPos(rad)))

	if r.parent.started && !fyne.CurrentApp().Settings().ShowAnimations() {
		r.drawStaticArc()
	}
}

func (r *ringActivityRenderer) MinSize() fyne.Size {
	return fyne.NewSquareSize(r.parent.Theme().Size(theme.SizeNameInlineIcon))
}

func (r *ringActivityRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.arc}
}

func (r *ringActivityRenderer) Refresh() {
	if r.parent.started {
		if !r.wasStarted {
			r.start()
		}
	} else if r.wasStarted {
		r.stop()
	}

	r.updateColor()
}

// animate moves the ring like Material Design's loading indicator:
// an inchworm motion (head sweeps ahead in the first half of each
// cycle, tail catches up in the second, each by ringActivityGrowRange) on
// top of a shared base made of cycleAdvance (compensates the inchworm's
// per-cycle reset, keeping position continuous across the seam) and spin (a
// small continuous rotation so neither edge is ever fully still).
// Since the inchworm term is shared and monotonic per half, the tail can never
// overtake the head. done goes from 0 to 1 once per animation cycle.
func (r *ringActivityRenderer) animate(done float32) {
	if done < r.lastDone {
		r.completedCycles++
		if r.completedCycles >= ringActivityCycleWrap {
			r.completedCycles -= ringActivityCycleWrap
		}
	}
	r.lastDone = done

	cycleAdvance := r.completedCycles * ringActivityGrowRange
	spin := (r.completedCycles + done) * ringActivityRotationPerCycle
	base := cycleAdvance + spin

	var tailFrac, headFrac float32
	if done < 0.5 {
		headFrac = easeInOut(done * 2)
	} else {
		tailFrac = easeInOut((done - 0.5) * 2)
		headFrac = 1
	}

	r.arc.StartAngle = base + tailFrac*ringActivityGrowRange
	r.arc.EndAngle = base + ringActivityMinSweep + headFrac*ringActivityGrowRange
	r.arc.Refresh()
}

// easeInOut is a smoothstep curve (t in 0..1): slow start, fast middle, slow
// end, and strictly monotonic so the edge it drives never reverses.
func easeInOut(t float32) float32 {
	return t * t * (3 - 2*t)
}

func (r *ringActivityRenderer) start() {
	r.wasStarted = true
	r.completedCycles = 0
	r.lastDone = 0
	if !fyne.CurrentApp().Settings().ShowAnimations() {
		r.drawStaticArc()
		return
	}
	r.anim.Start()
}

func (r *ringActivityRenderer) stop() {
	r.wasStarted = false
	r.anim.Stop()
	if !fyne.CurrentApp().Settings().ShowAnimations() {
		r.hideArc()
	}
}

// drawStaticArc places a fixed partial ring when animations are disabled, so
// the widget still indicates something is happening without any motion.
func (r *ringActivityRenderer) drawStaticArc() {
	r.arc.StartAngle = 0
	r.arc.EndAngle = ringActivityMaxSweep
	r.arc.Refresh()
}

func (r *ringActivityRenderer) hideArc() {
	r.arc.StartAngle = 0
	r.arc.EndAngle = 0
	r.arc.Refresh()
}

func (r *ringActivityRenderer) updateColor() {
	name := r.parent.ColorName
	if name == "" {
		name = theme.ColorNameForeground
	}
	v := fyne.CurrentApp().Settings().ThemeVariant()
	r.arc.FillColor = r.parent.Theme().Color(name, v)
	r.arc.Refresh()
}
