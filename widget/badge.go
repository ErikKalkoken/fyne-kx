package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Badge is a variant of the Fyne label widget that renders a rounded box around the text.
// Badges are commonly used to display counts.
type Badge struct {
	widget.BaseWidget

	Importance widget.Importance // Importance of the badge
	Text       string            // Text of the badge

	label      *widget.Label
	background *canvas.Rectangle
}

// NewBadge returns a new instance of a [Badge] widget.
func NewBadge(text string) *Badge {
	bg := canvas.NewRectangle(color.Transparent)
	bg.CornerRadius = 10
	w := &Badge{
		background: bg,
		label:      widget.NewLabel(text),
		Text:       text,
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetText sets the text of the badge.
func (w *Badge) SetText(text string) {
	w.Text = text
	w.Refresh()
}

func (w *Badge) Refresh() {
	w.label.Text = w.Text
	w.updateBadge()
	w.label.Refresh()
}

func (w *Badge) updateBadge() {
	th := w.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	switch w.Importance {
	case widget.DangerImportance:
		w.background.FillColor = th.Color(theme.ColorNameError, v)
	case widget.HighImportance:
		w.background.FillColor = th.Color(theme.ColorNamePrimary, v)
	case widget.LowImportance:
		w.background.FillColor = th.Color(theme.ColorNameDisabled, v)
	case widget.SuccessImportance:
		w.background.FillColor = th.Color(theme.ColorNameSuccess, v)
	case widget.WarningImportance:
		w.background.FillColor = th.Color(theme.ColorNameWarning, v)
	default:
		w.background.FillColor = th.Color(theme.ColorNameInputBackground, v)
	}
	p := th.Size(theme.SizeNameInnerPadding)
	s := w.label.MinSize().SubtractWidthHeight(p/2, p)
	w.background.SetMinSize(s)
	w.background.Refresh()
}

func (w *Badge) CreateRenderer() fyne.WidgetRenderer {
	p := w.Theme().Size(theme.SizeNameInnerPadding)
	w.updateBadge()
	return widget.NewSimpleRenderer(container.NewStack(
		container.New(layout.NewCustomPaddedLayout(p/2, p/2, p, p), w.background),
		container.NewCenter(w.label),
	))
}
