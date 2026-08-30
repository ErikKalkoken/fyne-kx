package widget

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

// newSizedSwitch resizes a Switch much larger than its graphic, so tests can prove the hitbox is limited to the graphic.
func newSizedSwitch(t *testing.T, w, h float32) *Switch {
	t.Helper()
	test.NewApp()
	sw := NewSwitch(nil)
	sw.Resize(fyne.NewSize(w, h))
	return sw
}

func TestSwitch_Layout_CachesGraphicBounds(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	assert.Equal(t, fyne.NewSize(switchWidth, switchHeight), sw.graphicSize,
		"Layout should cache the graphic's fixed size")
	assert.InDelta(t, float64(300)/2-float64(switchHeight)/2, float64(sw.graphicPosition.Y), 0.5,
		"Layout should cache the graphic's vertically centered position")
}

func TestSwitch_Layout_RecachesOnResize(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)
	firstY := sw.graphicPosition.Y

	sw.Resize(fyne.NewSize(300, 600))

	assert.NotEqual(t, firstY, sw.graphicPosition.Y,
		"resizing should recompute and recache the graphic bounds via Layout")
}

func TestSwitch_GraphicContains_InsideGraphic(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	// The center of the graphic should always be within the hitbox.
	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)

	assert.True(t, sw.graphicContains(center), "center of switch graphic should be within hitbox")
}

func TestSwitch_GraphicContains_OutsideGraphic_ContainerLarger(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	farAway := fyne.NewPos(250, 250)

	assert.False(t, sw.graphicContains(farAway), "points outside the graphic but inside the container must not be within the hitbox")
}

func TestSwitch_Tapped_OutsideGraphic_DoesNotToggle(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)
	assert.False(t, sw.On)

	sw.Tapped(&fyne.PointEvent{Position: fyne.NewPos(280, 280)})

	assert.False(t, sw.On, "tapping outside the graphic hitbox should not toggle the switch")
}

func TestSwitch_Tapped_InsideGraphic_Toggles(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)
	assert.False(t, sw.On)

	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)
	sw.Tapped(&fyne.PointEvent{Position: center})

	assert.True(t, sw.On, "tapping inside the graphic hitbox should toggle the switch")
}

func TestSwitch_Tapped_WhenDisabled_NeverToggles(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)
	sw.Disable()

	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)
	sw.Tapped(&fyne.PointEvent{Position: center})

	assert.False(t, sw.On, "a disabled switch should never toggle regardless of tap position")
}

func TestSwitch_MouseIn_OutsideGraphic_DoesNotHover(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	sw.MouseIn(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: fyne.NewPos(280, 280)}})

	assert.False(t, sw.hovered, "mouse entering the container far from the graphic should not set hover")
}

func TestSwitch_MouseIn_InsideGraphic_SetsHover(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)
	sw.MouseIn(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: center}})

	assert.True(t, sw.hovered, "mouse entering within the graphic bounds should set hover")
}

func TestSwitch_MouseMoved_TransitionsHoverInAndOut(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)
	outside := fyne.NewPos(280, 280)

	sw.MouseMoved(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: center}})
	assert.True(t, sw.hovered, "moving into the graphic should set hover")

	sw.MouseMoved(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: outside}})
	assert.False(t, sw.hovered, "moving out of the graphic (but still within the container) should clear hover")
}

func TestSwitch_MouseOut_ClearsHover(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)
	sw.MouseIn(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: center}})
	assert.True(t, sw.hovered)

	sw.MouseOut()
	assert.False(t, sw.hovered, "MouseOut should always clear hover")
}

func TestSwitch_Cursor_ReflectsHoverState(t *testing.T) {
	sw := newSizedSwitch(t, 300, 300)

	assert.Equal(t, desktop.DefaultCursor, sw.Cursor())

	center := fyne.NewPos(float32(switchWidth)/2+2, float32(300)/2)
	sw.MouseIn(&desktop.MouseEvent{PointEvent: fyne.PointEvent{Position: center}})
	assert.Equal(t, desktop.PointerCursor, sw.Cursor())
}
