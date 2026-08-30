package layout

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

// rects returns n canvas rectangles to use as layout children.
func rects(n int) []fyne.CanvasObject {
	objs := make([]fyne.CanvasObject, n)
	for i := range objs {
		objs[i] = canvas.NewRectangle(nil)
	}
	return objs
}

func TestNewColumns_PanicsWithNoWidths(t *testing.T) {
	assert.Panics(t, func() {
		NewColumns()
	})
}

func TestMinSize(t *testing.T) {
	padding := theme.Padding()
	l := NewColumns(100, 50, 30)

	size := l.MinSize(rects(3))

	assert.Equal(t, 100+50+30+2*padding, size.Width, "padding should only appear between columns")
}

func TestMinSize_ReusesLastWidth(t *testing.T) {
	padding := theme.Padding()
	l := NewColumns(100)

	size := l.MinSize(rects(3))

	assert.Equal(t, 100*3+2*padding, size.Width)
}

func TestLayout_PositionsColumnsSequentially(t *testing.T) {
	padding := theme.Padding()
	l := NewColumns(100, 50)
	objs := rects(2)

	l.Layout(objs, fyne.NewSize(-1, 10)) // negative width = no stretch

	assert.Equal(t, fyne.NewPos(0, 0), objs[0].Position())
	assert.Equal(t, float32(100), objs[0].Size().Width)

	assert.Equal(t, fyne.NewPos(100+padding, 0), objs[1].Position())
	assert.Equal(t, float32(50), objs[1].Size().Width)
}

func TestLayout_ReusesLastWidthForExtraColumns(t *testing.T) {
	padding := theme.Padding()
	l := NewColumns(100)
	objs := rects(3)

	l.Layout(objs, fyne.NewSize(-1, 10))

	assert.Equal(t, fyne.NewPos(0, 0), objs[0].Position())
	assert.Equal(t, fyne.NewPos(100+padding, 0), objs[1].Position())
	assert.Equal(t, fyne.NewPos(2*(100+padding), 0), objs[2].Position())
}

func TestLayout_LastColumnStretchesToFillContainer(t *testing.T) {
	padding := theme.Padding()
	l := NewColumns(100, 50)
	objs := rects(2)

	l.Layout(objs, fyne.NewSize(300, 10))

	assert.Equal(t, float32(100), objs[0].Size().Width)
	assert.Equal(t, float32(300)-(100+padding)-padding, objs[1].Size().Width)
}

func TestLayout_LastColumnNeverShrinksBelowConfiguredWidth(t *testing.T) {
	l := NewColumns(100, 50)
	objs := rects(2)

	l.Layout(objs, fyne.NewSize(120, 10)) // too small to fit both

	assert.Equal(t, float32(50), objs[1].Size().Width)
}

func TestLayout_TrailingStretchAppliesWhenWidthReused(t *testing.T) {
	padding := theme.Padding()
	l := NewColumns(100)
	objs := rects(3)

	l.Layout(objs, fyne.NewSize(500, 10))

	assert.Equal(t, float32(100), objs[0].Size().Width)
	assert.Equal(t, float32(100), objs[1].Size().Width)
	assert.Equal(t, float32(500)-2*(100+padding)-padding, objs[2].Size().Width)
}
