package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// disabledImageTranslucency is how much a disabled [TappableImage] fades its
// image. Translucency is applied to the rendered image, not the resource
// content, so it works the same regardless of whether the resource is a
// bitmap or a vector image, unlike [theme.NewDisabledResource] which only
// recolors SVG content.
const disabledImageTranslucency = 0.5

// TappableImage is widget which shows an image and calls a callback when tapped.
type TappableImage struct {
	widget.DisableableWidget

	// The function that is called when the label is tapped.
	OnTapped func()

	image   *canvas.Image
	hovered bool
	menu    *fyne.Menu
	pos     fyne.Position // current mouse position
}

var _ fyne.Tappable = (*TappableImage)(nil)
var _ desktop.Hoverable = (*TappableImage)(nil)
var _ fyne.Disableable = (*TappableImage)(nil)

// NewTappableImageWithMenu returns a new instance of a [TappableImage] widget with a context menu.
func NewTappableImageWithMenu(res fyne.Resource, menu *fyne.Menu) *TappableImage {
	w := newTappableImage(res, nil)
	if menu == nil {
		fyne.LogError("TappableImage misconfigured: missing menu", nil)
		return w
	}
	w.menu = menu
	w.OnTapped = func() {
		if len(w.menu.Items) == 0 {
			return
		}
		c := fyne.CurrentApp().Driver().CanvasForObject(w)
		m := widget.NewPopUpMenu(w.menu, c)
		m.ShowAtPosition(w.pos)
	}
	return w
}

// NewTappableImage returns a new instance of a [TappableImage] widget.
func NewTappableImage(res fyne.Resource, tapped func()) *TappableImage {
	return newTappableImage(res, tapped)
}

func newTappableImage(res fyne.Resource, tapped func()) *TappableImage {
	w := &TappableImage{OnTapped: tapped, image: canvas.NewImageFromResource(res)}
	w.ExtendBaseWidget(w)
	return w
}

// Refresh triggers a redraw of the image, applying the current disabled state.
func (w *TappableImage) Refresh() {
	if w.Disabled() {
		w.image.Translucency = disabledImageTranslucency
	} else {
		w.image.Translucency = 0
	}
	w.DisableableWidget.Refresh()
}

// SetFillMode sets the fill mode of the image.
func (w *TappableImage) SetFillMode(fillMode canvas.ImageFill) {
	w.image.FillMode = fillMode
}

// SetMinSize sets the minimum size of the image.
func (w *TappableImage) SetMinSize(size fyne.Size) {
	w.image.SetMinSize(size)
}

// SetResource sets the resource of the image.
func (w *TappableImage) SetResource(r fyne.Resource) {
	w.image.Resource = r
	w.image.Refresh()
}

// SetMenuItems replaces the menu items.
// Does nothing when the widget has not bee created with [NewTappableImageWithMenu].
func (w *TappableImage) SetMenuItems(menuItems []*fyne.MenuItem) {
	if w.menu == nil {
		return
	}
	w.menu.Items = menuItems
	w.menu.Refresh()
}

func (w *TappableImage) Tapped(pe *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	w.pos = pe.AbsolutePosition
	if w.OnTapped != nil {
		w.OnTapped()
	}
}

func (w *TappableImage) TappedSecondary(_ *fyne.PointEvent) {
}

// Cursor returns the cursor type of this widget
func (w *TappableImage) Cursor() desktop.Cursor {
	if !w.Disabled() && w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *TappableImage) MouseIn(me *desktop.MouseEvent) {
	w.MouseMoved(me)
}

func (w *TappableImage) MouseMoved(me *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.pos = me.AbsolutePosition
	pos := w.image.Position()
	s := w.image.Size()
	w.hovered = s.IsZero() ||
		(me.Position.X >= pos.X && me.Position.X <= pos.X+s.Width &&
			me.Position.Y >= pos.Y && me.Position.Y <= pos.Y+s.Height)
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *TappableImage) MouseOut() {
	w.hovered = false
}

func (w *TappableImage) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewPadded(w.image))
}
