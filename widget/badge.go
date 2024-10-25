package widget

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	badgeBackgroundColorDefault = theme.ColorNameInputBackground
)

// Badge is a variant of the Fyne label widget that renders a rounded box around the text.
// Badges are commonly used to display counts.
type Badge struct {
	*widget.Label
}

// NewBadge returns a new instance of a [Badge] widget.
func NewBadge(text string) *Badge {
	w := &Badge{Label: widget.NewLabel(text)}
	w.ExtendBaseWidget(w)
	return w
}

func (w *Badge) CreateRenderer() fyne.WidgetRenderer {
	r := w.Label.CreateRenderer()
	b := canvas.NewRectangle(color.Transparent)
	b.CornerRadius = 10
	r2 := &badgeRenderer{labelRenderer: r, background: b, badge: w}
	r2.updateBadge()
	return r2
}

var _ fyne.WidgetRenderer = (*badgeRenderer)(nil)

type badgeRenderer struct {
	labelRenderer fyne.WidgetRenderer
	background    *canvas.Rectangle
	badge         *Badge
}

func (r *badgeRenderer) Destroy() {
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

func (r *badgeRenderer) Refresh() {
	r.updateBadge()
	fmt.Printf("updated badge")
}

func (r *badgeRenderer) Objects() []fyne.CanvasObject {
	objs := []fyne.CanvasObject{r.background, r.labelRenderer.Objects()[0]}
	return objs
}

func (r *badgeRenderer) updateBadge() {
	th := r.badge.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	switch r.badge.Importance {
	default:
		r.background.FillColor = th.Color(theme.ColorNameInputBackground, v)
	}
	r.background.Refresh()
}
