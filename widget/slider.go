package widget

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Slider is a variant of the Fyne Slider widget that also displays the current value.
type Slider struct {
	widget.BaseWidget

	OnChangeEnded func(float64)

	min, max, step, value float64
}

// NewSlider returns a new instance of a [Slider] widget.
func NewSlider(min, max float64) *Slider {
	w := &Slider{
		min:   min,
		max:   max,
		value: min,
		step:  1, // matches widget.Slider's own default
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetStep sets a custom step for a slider.
func (w *Slider) SetStep(step float64) {
	w.step = step
	w.Refresh()
}

// Value returns the current value of a slider.
func (w *Slider) Value() float64 {
	return w.value
}

// SetValue set the value of a slider.
func (w *Slider) SetValue(v float64) {
	s := widget.NewSlider(w.min, w.max)
	s.Step = w.step
	s.SetValue(v)
	w.value = s.Value
	w.Refresh()
}

func (w *Slider) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(ftoa(w.value))
	label.Alignment = fyne.TextAlignTrailing

	slider := widget.NewSlider(w.min, w.max)
	slider.Step = w.step
	slider.Value = w.value

	r := &sliderRenderer{
		widget:  w,
		label:   label,
		slider:  slider,
		objects: []fyne.CanvasObject{label, slider},
	}
	r.updateColumnWidths()

	slider.OnChanged = func(v float64) {
		w.value = v
		label.SetText(ftoa(v))
	}
	slider.OnChangeEnded = func(v float64) {
		w.value = v
		if w.OnChangeEnded != nil {
			w.OnChangeEnded(v)
		}
	}

	return r
}

// sliderRenderer renders a [Slider]. It owns the child widgets and arranges
// them in two columns: the value label at a fixed width, and the slider
// filling the remaining space.
type sliderRenderer struct {
	widget  *Slider
	label   *widget.Label
	slider  *widget.Slider
	objects []fyne.CanvasObject

	// labelW and sliderMinW cache the result of updateColumnWidths, since
	// they only change when the slider's min, max or step change.
	labelW, sliderMinW float32
}

// updateColumnWidths recomputes labelW, the width reserved for the label,
// and sliderMinW, the minimum width for the slider. The label width is
// sized to fit its value at both the slider's minimum and maximum, so the
// slider doesn't shift horizontally as the label's text changes.
//
// This is only called on construction and from Refresh (which runs on
// SetStep and on theme changes), not from Layout/MinSize, since those are
// called far more often but min/max/step don't change between calls.
func (r *sliderRenderer) updateColumnWidths() {
	minW1 := labelWidth(ftoa(r.slider.Max + r.slider.Step))
	minW2 := labelWidth(ftoa(r.slider.Min - r.slider.Step))
	r.labelW = minW1
	r.sliderMinW = max(minW1, minW2, r.slider.MinSize().Width)
}

func labelWidth(s string) float32 {
	return widget.NewLabel(s).MinSize().Width
}

func (r *sliderRenderer) Layout(size fyne.Size) {
	padding := theme.Padding()

	r.label.Resize(fyne.NewSize(r.labelW, r.label.MinSize().Height))
	r.label.Move(fyne.NewPos(0, 0))

	sliderW := fyne.Max(size.Width-r.labelW-2*padding, r.sliderMinW)
	r.slider.Resize(fyne.NewSize(sliderW, r.slider.MinSize().Height))
	r.slider.Move(fyne.NewPos(r.labelW+padding, 0))
}

func (r *sliderRenderer) MinSize() fyne.Size {
	padding := theme.Padding()
	h := fyne.Max(r.label.MinSize().Height, r.slider.MinSize().Height)
	return fyne.NewSize(r.labelW+r.sliderMinW+2*padding, h)
}

func (r *sliderRenderer) Refresh() {
	r.slider.Step = r.widget.step
	if r.slider.Value != r.widget.value {
		r.slider.Value = r.widget.value
		r.slider.Refresh()
	}
	r.updateColumnWidths()
	r.label.SetText(ftoa(r.widget.value))
	r.label.Refresh()
	r.slider.Refresh()
}

func (r *sliderRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *sliderRenderer) Destroy() {}

// ftoa returns a string representation of a float without any unnecessary zeros.
func ftoa(f float64) string {
	s := strconv.FormatFloat(f, 'f', 6, 64)
	s = strings.TrimRight(s, "0")
	return strings.TrimSuffix(s, ".")
}
