package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// FilterChip represents a basic filter chip widget.
// It shows a label and has two states: on or off.
type FilterChip struct {
	chip

	// OnChanged is called when the state changed
	OnChanged func(on bool)
	On        bool
	Text      string
}

// NewFilterChip returns a new [FilterChip] object.
func NewFilterChip(text string, changed func(on bool)) *FilterChip {
	w := &FilterChip{
		OnChanged: changed,
		Text:      text,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *FilterChip) Refresh() {
	w.on = w.On
	w.text = w.Text
	if w.On {
		w.leadingIcon = theme.ConfirmIcon()
	} else {
		w.leadingIcon = nil
	}
	w.chip.Refresh()
}

// SetState sets the state.
func (w *FilterChip) SetState(v bool) {
	if w.On == v {
		return
	}
	w.On = v
	if w.OnChanged != nil {
		w.OnChanged(v)
	}
	w.Refresh()
}

// SetText sets the label text.
func (w *FilterChip) SetText(text string) {
	w.Text = text
	w.Refresh()
}

func (w *FilterChip) CreateRenderer() fyne.WidgetRenderer {
	w.on = w.On
	w.text = w.Text
	w.onTapped = func() {
		w.SetState(!w.On)
	}
	return w.chip.CreateRenderer()
}
