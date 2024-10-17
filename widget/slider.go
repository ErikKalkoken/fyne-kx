package widget

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/ErikKalkoken/fyne-kwidget/layout"
)

// Slider is a widget that can slide between two fixed values
// and always shows the current value.
type Slider struct {
	widget.BaseWidget

	OnChangeEnded func(int)

	data   binding.Float
	label  *widget.Label
	layout fyne.Layout
	slider *widget.Slider
}

// NewSlider returns a new instance of a [Slider] widget.
func NewSlider(min, max, start int) *Slider {
	temp := widget.NewLabel(strconv.Itoa(max))
	minW := temp.MinSize().Width
	d := binding.NewFloat()
	w := &Slider{
		label:  widget.NewLabelWithData(binding.FloatToStringWithFormat(d, "%.0f")),
		slider: widget.NewSliderWithData(float64(min), float64(max), d),
		data:   d,
		layout: layout.NewColumnsLayout(minW, 2*minW),
	}
	w.label.Alignment = fyne.TextAlignTrailing
	w.slider.OnChangeEnded = func(v float64) {
		if w.OnChangeEnded == nil {
			return
		}
		w.OnChangeEnded(int(v))
	}
	w.slider.SetValue(float64(start))
	w.ExtendBaseWidget(w)
	return w
}

// Value returns the current value of a slider.
func (w *Slider) Value() int {
	return int(w.slider.Value)
}

// SetValue set the value of a slider.
func (w *Slider) SetValue(v int) {
	w.slider.SetValue(float64(v))
}

func (w *Slider) CreateRenderer() fyne.WidgetRenderer {
	c := container.New(w.layout, w.label, w.slider)
	return widget.NewSimpleRenderer(c)
}
