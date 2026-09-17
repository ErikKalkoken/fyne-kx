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
	spinnerAnimationDuration = time.Second
	spinnerMinSweep          = float32(20)
	spinnerMaxSweep          = float32(280)
	spinnerCutoutRatio       = float32(0.70) // defines the ring width
	spinnerRotationPerCycle  = float32(360)
	// distance each edge travels per inchworm half
	spinnerGrowRange = spinnerMaxSweep - spinnerMinSweep
	// completedCycles wraps at this value in animate() to keep it bounded
	spinnerCycleWrap = float32(18)
)

var _ fyne.Widget = (*Spinner)(nil)

// A Spinner indicates that an operation is in progress with unknown duration.
// It shows an animated ring, similar to a Material Design circular progress indicator.
type Spinner struct {
	widget.BaseWidget

	// ColorName is the named theme color used to draw the ring. It defaults
	// to [theme.ColorNameForeground] when left empty.
	ColorName fyne.ThemeColorName

	started bool
}

// NewSpinner returns a new [Spinner].
func NewSpinner() *Spinner {
	a := &Spinner{}
	a.ExtendBaseWidget(a)
	return a
}

// SetColorName sets the named theme color used to draw the ring.
func (a *Spinner) SetColorName(name fyne.ThemeColorName) {
	a.ColorName = name
	a.Refresh()
}

func (a *Spinner) MinSize() fyne.Size {
	a.ExtendBaseWidget(a)
	return a.BaseWidget.MinSize()
}

// Start the activity indicator animation.
func (a *Spinner) Start() {
	if a.started {
		return // already started
	}

	a.started = true

	a.Refresh()
}

// Stop the activity indicator animation.
func (a *Spinner) Stop() {
	if !a.started {
		return // already stopped
	}

	a.started = false

	a.Refresh()
}

func (a *Spinner) CreateRenderer() fyne.WidgetRenderer {
	arc := canvas.NewArc(0, 0, spinnerCutoutRatio, color.Transparent)
	r := &spinnerRenderer{arc: arc, parent: a}
	r.anim = &fyne.Animation{
		Duration:    spinnerAnimationDuration,
		RepeatCount: fyne.AnimationRepeatForever,
		Curve:       fyne.AnimationLinear,
		Tick:        r.animate,
	}
	r.updateColor()

	if a.started {
		r.start()
	}

	return r
}

var _ fyne.WidgetRenderer = (*spinnerRenderer)(nil)

type spinnerRenderer struct {
	anim   *fyne.Animation
	arc    *canvas.Arc
	parent *Spinner

	wasStarted bool

	// completedCycles counts full cycles seen so far, for the continuous
	// spin and cycleAdvance in animate. Wrapped by spinnerCycleWrap.
	completedCycles float32
	// lastDone is done from the previous tick, to detect a new cycle starting.
	lastDone float32
}

func (r *spinnerRenderer) Destroy() {
	r.parent.started = false
	r.stop()
}

func (r *spinnerRenderer) Layout(size fyne.Size) {
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

func (r *spinnerRenderer) MinSize() fyne.Size {
	return fyne.NewSquareSize(r.parent.Theme().Size(theme.SizeNameInlineIcon))
}

func (r *spinnerRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.arc}
}

func (r *spinnerRenderer) Refresh() {
	if r.parent.started {
		if !r.wasStarted {
			r.start()
		}
	} else if r.wasStarted {
		r.stop()
	}

	r.updateColor()
}

// animate uses an inchworm motion so the tail can never overtake the head.
func (r *spinnerRenderer) animate(done float32) {
	if done < r.lastDone {
		r.completedCycles++
		if r.completedCycles >= spinnerCycleWrap {
			r.completedCycles -= spinnerCycleWrap
		}
	}
	r.lastDone = done

	cycleAdvance := r.completedCycles * spinnerGrowRange
	spin := (r.completedCycles + done) * spinnerRotationPerCycle
	base := cycleAdvance + spin

	var tailFrac, headFrac float32
	if done < 0.5 {
		headFrac = easeInOut(done * 2)
	} else {
		tailFrac = easeInOut((done - 0.5) * 2)
		headFrac = 1
	}

	r.arc.StartAngle = base + tailFrac*spinnerGrowRange
	r.arc.EndAngle = base + spinnerMinSweep + headFrac*spinnerGrowRange
	r.arc.Refresh()
}

// easeInOut is a smoothstep curve (t in 0..1): slow start, fast middle, slow
// end, and strictly monotonic so the edge it drives never reverses.
func easeInOut(t float32) float32 {
	return t * t * (3 - 2*t)
}

func (r *spinnerRenderer) start() {
	r.wasStarted = true
	r.completedCycles = 0
	r.lastDone = 0
	if !fyne.CurrentApp().Settings().ShowAnimations() {
		r.drawStaticArc()
		return
	}
	r.anim.Start()
}

func (r *spinnerRenderer) stop() {
	r.wasStarted = false
	r.anim.Stop()
	if !fyne.CurrentApp().Settings().ShowAnimations() {
		r.hideArc()
	}
}

// drawStaticArc places a fixed partial ring when animations are disabled, so
// the widget still indicates something is happening without any motion.
func (r *spinnerRenderer) drawStaticArc() {
	r.arc.StartAngle = 0
	r.arc.EndAngle = spinnerMaxSweep
	r.arc.Refresh()
}

func (r *spinnerRenderer) hideArc() {
	r.arc.StartAngle = 0
	r.arc.EndAngle = 0
	r.arc.Refresh()
}

func (r *spinnerRenderer) updateColor() {
	name := r.parent.ColorName
	if name == "" {
		name = theme.ColorNameForeground
	}
	v := fyne.CurrentApp().Settings().ThemeVariant()
	r.arc.FillColor = r.parent.Theme().Color(name, v)
	r.arc.Refresh()
}
