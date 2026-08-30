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
	w.value = v
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
}

// columnWidths returns the width reserved for the label and the minimum
// width for the slider. The label width is sized to fit its value at both
// the slider's minimum and maximum, so the slider doesn't shift
// horizontally as the label's text changes.
func (r *sliderRenderer) columnWidths() (labelW, sliderW float32) {
	minW1 := labelWidth(ftoa(r.slider.Max + r.slider.Step))
	minW2 := labelWidth(ftoa(r.slider.Min - r.slider.Step))
	labelW = minW1
	sliderW = max(minW1, minW2, r.slider.MinSize().Width)
	return labelW, sliderW
}

func labelWidth(s string) float32 {
	return widget.NewLabel(s).MinSize().Width
}

func (r *sliderRenderer) Layout(size fyne.Size) {
	padding := theme.Padding()
	labelW, sliderMinW := r.columnWidths()

	r.label.Resize(fyne.NewSize(labelW, r.label.MinSize().Height))
	r.label.Move(fyne.NewPos(0, 0))

	sliderW := fyne.Max(size.Width-labelW-2*padding, sliderMinW)
	r.slider.Resize(fyne.NewSize(sliderW, r.slider.MinSize().Height))
	r.slider.Move(fyne.NewPos(labelW+padding, 0))
}

func (r *sliderRenderer) MinSize() fyne.Size {
	padding := theme.Padding()
	labelW, sliderMinW := r.columnWidths()
	h := fyne.Max(r.label.MinSize().Height, r.slider.MinSize().Height)
	return fyne.NewSize(labelW+sliderMinW+2*padding, h)
}

func (r *sliderRenderer) Refresh() {
	r.slider.Step = r.widget.step
	if r.slider.Value != r.widget.value {
		r.slider.Value = r.widget.value
		r.slider.Refresh()
	}
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
