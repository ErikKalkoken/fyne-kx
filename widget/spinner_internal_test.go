package widget

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

func newTestSpinner(t *testing.T) (*Spinner, *spinnerRenderer) {
	t.Helper()
	test.NewApp()
	a := NewSpinner()
	return a, a.CreateRenderer().(*spinnerRenderer)
}

func TestSpinner_SetColorName_UpdatesField(t *testing.T) {
	a, _ := newTestSpinner(t)

	a.SetColorName(theme.ColorNameError)

	assert.Equal(t, theme.ColorNameError, a.ColorName)
}

func TestSpinnerRenderer_Animate_AtCycleStart(t *testing.T) {
	_, r := newTestSpinner(t)

	r.animate(0)

	assert.Equal(t, float32(0), r.arc.StartAngle)
	assert.Equal(t, spinnerMinSweep, r.arc.EndAngle)
}

func TestSpinnerRenderer_Animate_TailTracksOnlySpinDuringFirstHalf(t *testing.T) {
	_, r := newTestSpinner(t)

	r.animate(0.1)
	tailBefore, headBefore := r.arc.StartAngle, r.arc.EndAngle
	r.animate(0.4)

	wantTailDelta := float64((0.4 - 0.1) * spinnerRotationPerCycle)
	assert.InDelta(t, wantTailDelta, float64(r.arc.StartAngle-tailBefore), 0.01,
		"tail should advance by exactly the background spin during the first half")
	assert.Greater(t, r.arc.EndAngle-headBefore, r.arc.StartAngle-tailBefore, "head should advance faster than the tail")
}

func TestSpinnerRenderer_Animate_HeadTracksOnlySpinDuringSecondHalf(t *testing.T) {
	_, r := newTestSpinner(t)

	r.animate(0.6)
	tailBefore, headBefore := r.arc.StartAngle, r.arc.EndAngle
	r.animate(0.9)

	wantHeadDelta := float64((0.9 - 0.6) * spinnerRotationPerCycle)
	assert.InDelta(t, wantHeadDelta, float64(r.arc.EndAngle-headBefore), 0.01,
		"head should advance by exactly the background spin during the second half")
	assert.Greater(t, r.arc.StartAngle-tailBefore, r.arc.EndAngle-headBefore, "tail should advance faster than the head")
}

func TestSpinnerRenderer_Animate_SweepIsPeriodic(t *testing.T) {
	_, r := newTestSpinner(t)

	r.animate(0)
	gapAtStart := r.arc.EndAngle - r.arc.StartAngle
	r.animate(1)
	gapAtEnd := r.arc.EndAngle - r.arc.StartAngle

	assert.InDelta(t, gapAtStart, gapAtEnd, 0.01, "sweep length should be equal at done==0 and done==1")
}

func TestSpinnerRenderer_Animate_TailNeverOvertakesHead(t *testing.T) {
	_, r := newTestSpinner(t)

	for i := 0; i <= 100; i++ {
		r.animate(float32(i) / 100)
		assert.GreaterOrEqual(t, r.arc.EndAngle, r.arc.StartAngle)
	}
}

func TestSpinnerRenderer_Animate_ContinuousAcrossCycleWrap(t *testing.T) {
	_, r := newTestSpinner(t)

	// A tiny epsilon on each side of the seam: some drift is expected since
	// real "done" time still passes between the two samples (the spin term
	// keeps advancing at its normal rate), but for an epsilon this small
	// that drift is under a degree - nowhere near the ~260 degree pop the
	// old (unpatched) formula produced at the seam.
	const eps = float32(0.001)
	r.animate(1 - eps)
	tailBefore, headBefore := r.arc.StartAngle, r.arc.EndAngle
	r.animate(eps) // simulate the animation wrapping into its next cycle

	assert.InDelta(t, float64(tailBefore), float64(r.arc.StartAngle), 2, "tail should not jump across the cycle seam")
	assert.InDelta(t, float64(headBefore), float64(r.arc.EndAngle), 2, "head should not jump across the cycle seam")
	assert.Equal(t, float32(1), r.completedCycles)
}

func TestSpinnerRenderer_StartStop_ToggleWasStarted(t *testing.T) {
	a, r := newTestSpinner(t)
	a.started = true

	r.start()
	assert.True(t, r.wasStarted)

	r.stop()
	assert.False(t, r.wasStarted)
}

func TestSpinnerRenderer_Start_ResetsCycleState(t *testing.T) {
	_, r := newTestSpinner(t)
	r.animate(0.5)
	r.completedCycles = 3

	r.start()

	// start() resets completedCycles to 0 before re-arming the animation;
	// lastDone isn't asserted here since starting also ticks the animation
	// once (at least under the test driver), which legitimately advances it.
	assert.Equal(t, float32(0), r.completedCycles, "start should reset the cycle counter")
}

func TestSpinnerRenderer_DrawStaticArc(t *testing.T) {
	_, r := newTestSpinner(t)

	r.drawStaticArc()

	assert.Equal(t, float32(0), r.arc.StartAngle)
	assert.Equal(t, spinnerMaxSweep, r.arc.EndAngle)
}

func TestSpinnerRenderer_HideArc(t *testing.T) {
	_, r := newTestSpinner(t)

	r.hideArc()

	assert.Equal(t, float32(0), r.arc.StartAngle)
	assert.Equal(t, float32(0), r.arc.EndAngle)
}

func TestEaseInOut_IsMonotonicAndBounded(t *testing.T) {
	assert.Equal(t, float32(0), easeInOut(0))
	assert.Equal(t, float32(1), easeInOut(1))
	assert.Equal(t, float32(0.5), easeInOut(0.5))

	prev := float32(-1)
	for i := 0; i <= 100; i++ {
		v := easeInOut(float32(i) / 100)
		assert.GreaterOrEqual(t, v, prev)
		prev = v
	}
}
