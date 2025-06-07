package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/ErikKalkoken/fyne-kx/layout"
)

// Slider is a variant of the Fyne Slider widget that also displays the current value.
type Slider struct {
	widget.BaseWidget

	OnChangeEnded func(float64)

	data   binding.Float
	label  *widget.Label
	layout fyne.Layout
	slider *widget.Slider
}

// NewSlider returns a new instance of a [Slider] widget.
func NewSlider(min, max float64) *Slider {
	d := binding.NewFloat()
	label := widget.NewLabelWithData(binding.FloatToStringWithFormat(d, "%v"))
	label.Alignment = fyne.TextAlignTrailing
	w := &Slider{
		label:  label,
		slider: widget.NewSliderWithData(min, max, d),
		data:   d,
	}
	w.ExtendBaseWidget(w)
	w.updateLayout()
	w.slider.OnChangeEnded = func(v float64) {
		if w.OnChangeEnded == nil {
			return
		}
		w.OnChangeEnded(v)
	}
	return w
}

// SetStep sets a custom step for a slider.
func (w *Slider) SetStep(step float64) {
	w.slider.Step = step
	w.updateLayout()
}

func (w *Slider) updateLayout() {
	x := widget.NewLabel(fmt.Sprint(w.slider.Max + w.slider.Step))
	minW := x.MinSize().Width
	w.layout = layout.NewColumns(minW, fyne.Max(minW, w.slider.MinSize().Width))
}

// Value returns the current value of a slider.
func (w *Slider) Value() float64 {
	return w.slider.Value
}

// SetValue set the value of a slider.
func (w *Slider) SetValue(v float64) {
	w.slider.SetValue(float64(v))
}

func (w *Slider) CreateRenderer() fyne.WidgetRenderer {
	c := container.New(w.layout, w.label, w.slider)
	return widget.NewSimpleRenderer(c)
}
