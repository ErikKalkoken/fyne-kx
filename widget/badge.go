package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Badge is a variant of the Fyne label widget that renders a rounded box around the text.
// Badges are commonly used to display counts.
type Badge struct {
	widget.BaseWidget

	Importance widget.Importance
	Text       string
	Alignment  fyne.TextAlign // Alignment of the badge text and background
	SizeName   fyne.ThemeSizeName
}

// NewBadge returns a new instance of a [Badge] widget.
func NewBadge(text string) *Badge {
	w := &Badge{
		Text: text,
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetText sets the text of the badge.
func (w *Badge) SetText(text string) {
	w.Text = text
	w.Refresh()
}

func (w *Badge) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(color.Transparent)

	innerBg := canvas.NewRectangle(color.Transparent)
	innerBg.CornerRadius = theme.Size(theme.SizeNameDialogRadius)

	label := widget.NewLabel(w.Text)

	r := &badgeRenderer{
		badge:           w,
		background:      bg,
		innerBackground: innerBg,
		label:           label,
	}
	r.Refresh()
	return r
}

type badgeRenderer struct {
	badge           *Badge
	background      *canvas.Rectangle
	innerBackground *canvas.Rectangle
	label           *widget.Label
}

func (r *badgeRenderer) Destroy() {}

func (r *badgeRenderer) Layout(size fyne.Size) {
	labelSize := r.label.MinSize()
	bgSize := labelSize
	padding := theme.Padding()

	// Calculate horizontal positioning based on Alignment
	var xPos float32
	switch r.badge.Alignment {
	case fyne.TextAlignTrailing:
		xPos = size.Width - bgSize.Width
	case fyne.TextAlignCenter:
		xPos = (size.Width - bgSize.Width) / 2
	case fyne.TextAlignLeading:
		fallthrough
	default:
		xPos = 0
	}

	// Vertical centering within available height
	yPosBg := (size.Height - bgSize.Height) / 2
	yPosLabel := (size.Height - labelSize.Height) / 2

	// Position and resize the outer background
	r.background.Move(fyne.NewPos(xPos, yPosBg))
	r.background.Resize(bgSize)

	// Position and resize the inner background (inset by padding)
	innerPos := fyne.NewPos(xPos+padding, yPosBg+padding)
	innerSize := fyne.NewSize(bgSize.Width-(padding*2), bgSize.Height-(padding*2))

	// Guard against negative dimensions when label is very small
	if innerSize.Width < 0 {
		innerSize.Width = 0
	}
	if innerSize.Height < 0 {
		innerSize.Height = 0
	}

	r.innerBackground.Move(innerPos)
	r.innerBackground.Resize(innerSize)

	// Position the label centered relative to the background
	labelXPos := xPos + (bgSize.Width-labelSize.Width)/2
	r.label.Move(fyne.NewPos(labelXPos, yPosLabel))
	r.label.Resize(labelSize)
}

func (r *badgeRenderer) MinSize() fyne.Size {
	return r.label.MinSize()
}

func (r *badgeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.innerBackground, r.label}
}

func (r *badgeRenderer) Refresh() {
	r.label.Text = r.badge.Text
	r.label.Alignment = r.badge.Alignment
	r.label.SizeName = r.badge.SizeName
	r.label.Refresh()

	th := r.badge.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	// Outer background is explicitly transparent
	r.background.FillColor = color.Transparent

	// Inner background fill color mapped from Importance
	switch r.badge.Importance {
	case widget.DangerImportance:
		r.innerBackground.FillColor = th.Color(theme.ColorNameError, v)
	case widget.HighImportance:
		r.innerBackground.FillColor = th.Color(theme.ColorNamePrimary, v)
	case widget.LowImportance:
		r.innerBackground.FillColor = th.Color(theme.ColorNameDisabled, v)
	case widget.SuccessImportance:
		r.innerBackground.FillColor = th.Color(theme.ColorNameSuccess, v)
	case widget.WarningImportance:
		r.innerBackground.FillColor = th.Color(theme.ColorNameWarning, v)
	default:
		r.innerBackground.FillColor = th.Color(theme.ColorNameInputBackground, v)
	}

	r.background.Refresh()
	r.innerBackground.Refresh()
}
