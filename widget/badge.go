package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Badge is a variant of Fyne label widget that renders as badge.
// Badges are common UI elements to display counts.
type Badge struct {
	*widget.Label
}

// NewBadge returns a new instance of a [Badge] widget.
func NewBadge(text string) *Badge {
	w := &Badge{Label: widget.NewLabel(text)}
	w.ExtendBaseWidget(w)
	w.SetText(text)
	return w
}

func (w *Badge) CreateRenderer() fyne.WidgetRenderer {
	r := w.Label.CreateRenderer()
	b := canvas.NewRectangle(theme.Color(theme.ColorNameInputBackground))
	b.CornerRadius = 10
	return &badgeRenderer{WidgetRenderer: r, background: b}
}

type badgeRenderer struct {
	fyne.WidgetRenderer

	background *canvas.Rectangle
}

func (r *badgeRenderer) Layout(size fyne.Size) {
	topLeft := fyne.NewPos(0, 0)
	objs := r.Objects()
	bg := objs[0]
	label := objs[1]

	s := label.MinSize()
	label.Resize(s)
	label.Move(topLeft)
	bg.Resize(s)
	bg.Move(topLeft)
}

func (r *badgeRenderer) MinSize() fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range r.Objects() {
		minSize = minSize.Max(child.MinSize())
	}
	return minSize
}

func (r *badgeRenderer) Objects() []fyne.CanvasObject {
	objs := []fyne.CanvasObject{r.background, r.WidgetRenderer.Objects()[0]}
	return objs
}
