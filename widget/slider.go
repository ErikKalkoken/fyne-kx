package widget

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/ErikKalkoken/fyne-kx/layout"
)

// Slider is a variant of the Fyne Slider widget that also displays the current value.
type Slider struct {
	widget.BaseWidget

	OnChangeEnded func(float64)

	label  *widget.Label
	layout fyne.Layout
	slider *widget.Slider
}

// NewSlider returns a new instance of a [Slider] widget.
func NewSlider(min, max float64) *Slider {
	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignTrailing
	w := &Slider{
		label:  label,
		slider: widget.NewSlider(min, max),
	}
	w.ExtendBaseWidget(w)
	w.updateLayout()
	w.slider.OnChangeEnded = func(v float64) {
		if w.OnChangeEnded == nil {
			return
		}
		w.OnChangeEnded(v)
	}
	updateLabel := func(f float64) {
		w.label.SetText(ftoa(f))
	}
	w.slider.OnChanged = func(f float64) {
		updateLabel(f)
	}
	updateLabel(w.slider.Value)
	return w
}

// SetStep sets a custom step for a slider.
func (w *Slider) SetStep(step float64) {
	w.slider.Step = step
	w.updateLayout()
}

func (w *Slider) updateLayout() {
	x1 := widget.NewLabel(ftoa(w.slider.Max + w.slider.Step))
	minW1 := x1.MinSize().Width
	x2 := widget.NewLabel(ftoa(w.slider.Min - w.slider.Step))
	minW2 := x2.MinSize().Width
	w.layout = layout.NewColumns(minW1, fyne.Max(fyne.Max(minW1, minW2), w.slider.MinSize().Width))
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

// ftoa returns a string representation of a float without any unnecessary zeros.
func ftoa(f float64) string {
	return stripTrailingZeros(strconv.FormatFloat(f, 'f', 6, 64))
}

func stripTrailingZeros(s string) string {
	if !strings.ContainsRune(s, '.') {
		return s
	}
	offset := len(s) - 1
	for offset > 0 {
		if s[offset] == '.' {
			offset--
			break
		}
		if s[offset] != '0' {
			break
		}
		offset--
	}
	return s[:offset+1]
}
