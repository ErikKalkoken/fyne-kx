package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/ErikKalkoken/fyne-kx/layout"
)

// SliderWithValue is a variation of the Slider widget that also displays the current value.
type SliderWithValue struct {
	widget.BaseWidget

	OnChangeEnded func(float64)

	data   binding.Float
	label  *widget.Label
	layout fyne.Layout
	slider *widget.Slider
}

// NewSliderWithValue returns a new instance of a [SliderWithValue] widget.
func NewSliderWithValue(min, max float64) *SliderWithValue {
	d := binding.NewFloat()
	w := &SliderWithValue{
		label:  widget.NewLabelWithData(binding.FloatToStringWithFormat(d, "%v")),
		slider: widget.NewSliderWithData(min, max, d),
		data:   d,
	}
	w.updateLayout()
	w.label.Alignment = fyne.TextAlignTrailing
	w.slider.OnChangeEnded = func(v float64) {
		if w.OnChangeEnded == nil {
			return
		}
		w.OnChangeEnded(v)
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetStep sets a custom step for a slider.
func (w *SliderWithValue) SetStep(step float64) {
	w.slider.Step = step
	w.updateLayout()
}

func (w *SliderWithValue) updateLayout() {
	x := widget.NewLabel(fmt.Sprintf("%v", w.slider.Max+w.slider.Step))
	minW := x.MinSize().Width
	w.layout = layout.NewColumns(minW, minW)
}

// Value returns the current value of a slider.
func (w *SliderWithValue) Value() float64 {
	return w.slider.Value
}

// SetValue set the value of a slider.
func (w *SliderWithValue) SetValue(v float64) {
	w.slider.SetValue(float64(v))
}

func (w *SliderWithValue) CreateRenderer() fyne.WidgetRenderer {
	c := container.New(w.layout, w.label, w.slider)
	return widget.NewSimpleRenderer(c)
}
