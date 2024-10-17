package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/constraints"

	"github.com/ErikKalkoken/fyne-kx/layout"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

// Slider is a widget that can slide between two fixed values
// and always displays the current value.
// The widget can be instantiated for any numeric type.
type Slider[T Numeric] struct {
	widget.BaseWidget

	OnChangeEnded func(T)

	data   binding.Float
	label  *widget.Label
	layout fyne.Layout
	slider *widget.Slider
}

// NewSlider returns a new instance of a [Slider] widget.
func NewSlider[T Numeric](min, max, step T) *Slider[T] {
	x := widget.NewLabel(fmt.Sprintf("%v", max+step))
	minW := x.MinSize().Width
	d := binding.NewFloat()
	w := &Slider[T]{
		label:  widget.NewLabelWithData(binding.FloatToStringWithFormat(d, "%v")),
		slider: widget.NewSliderWithData(float64(min), float64(max), d),
		data:   d,
		layout: layout.NewColumnsLayout(minW, 2*minW),
	}
	w.label.Alignment = fyne.TextAlignTrailing
	w.slider.OnChangeEnded = func(v float64) {
		if w.OnChangeEnded == nil {
			return
		}
		w.OnChangeEnded(T(v))
	}
	w.slider.Step = float64(step)
	w.ExtendBaseWidget(w)
	return w
}

// Value returns the current value of a slider.
func (w *Slider[T]) Value() T {
	return T(w.slider.Value)
}

// SetValue set the value of a slider.
func (w *Slider[T]) SetValue(v T) {
	w.slider.SetValue(float64(v))
}

func (w *Slider[T]) CreateRenderer() fyne.WidgetRenderer {
	c := container.New(w.layout, w.label, w.slider)
	return widget.NewSimpleRenderer(c)
}
