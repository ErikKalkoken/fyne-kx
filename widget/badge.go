package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Badge is a variant of Fyne label widget that renders as badge.
// Badges are common UI elements to display counts.
type Badge struct {
	widget.BaseWidget
	label *widget.Label
}

// NewBadge returns a new instance of a [Badge] widget.
func NewBadge(text string) *Badge {
	l := widget.NewLabel(fmt.Sprintf(" %s ", text))
	l.Alignment = fyne.TextAlignCenter
	w := &Badge{label: l}
	w.ExtendBaseWidget(w)
	return w
}

// SetText sets the text of a badge.
func (w *Badge) SetText(text string) {
	w.label.SetText(fmt.Sprintf(" %s ", text))
}

func (w *Badge) CreateRenderer() fyne.WidgetRenderer {
	r := canvas.NewRectangle(theme.Color(theme.ColorNameInputBackground))
	r.CornerRadius = 10
	c := container.NewStack(container.NewPadded(r), w.label)
	return widget.NewSimpleRenderer(c)
}
