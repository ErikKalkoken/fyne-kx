package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// disabledImageTranslucency is how much a disabled [TappableImage] fades its image.
const disabledImageTranslucency = 0.5

// TappableImage is widget which shows an image and runs a callback when tapped.
//
// It can be extended by embedding it by value and calling ExtendBaseWidget.
// Overrides of ExtendBaseWidget must call the TappableImage version.
// When combined with another hoverable type (e.g. a tooltip extension),
// MouseIn, MouseOut and MouseMoved must be forwarded to both.
type TappableImage struct {
	widget.DisableableWidget

	// The function that is called when the image is tapped. Takes precedence over the menu.
	OnTapped func()

	// Resource is the resource shown by this image.
	Resource fyne.Resource

	// FillMode is the fill mode of the image.
	FillMode canvas.ImageFill

	// ScaleMode sets the scaling filter used to scale the image.
	ScaleMode canvas.ImageScale

	// CornerRadius specifies a radius to apply to round the corners of the image.
	CornerRadius float32

	// Translucency sets a base translucency value > 0.0 to fade the image.
	// While disabled, the disabled-state fade is composited on top of this
	// value rather than replacing it.
	Translucency float64

	image   *canvas.Image // image is created lazily in CreateRenderer
	minSize fyne.Size
	hovered bool
	menu    *fyne.Menu
	pos     fyne.Position // current mouse position
	self    fyne.Widget   // outermost widget, differs from w when embedded
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
	return w
}

// NewTappableImage returns a new instance of a [TappableImage] widget.
func NewTappableImage(res fyne.Resource, tapped func()) *TappableImage {
	return newTappableImage(res, tapped)
}

func newTappableImage(res fyne.Resource, tapped func()) *TappableImage {
	w := &TappableImage{OnTapped: tapped, Resource: res}
	w.ExtendBaseWidget(w)
	return w
}

// ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.
func (w *TappableImage) ExtendBaseWidget(wid fyne.Widget) {
	w.self = wid
	w.DisableableWidget.ExtendBaseWidget(wid)
}

// super returns the outermost widget, which Fyne knows as part of the canvas tree.
func (w *TappableImage) super() fyne.CanvasObject {
	if w.self == nil {
		return w
	}
	return w.self
}

// SetFillMode sets the fill mode of the image.
func (w *TappableImage) SetFillMode(fillMode canvas.ImageFill) {
	w.FillMode = fillMode
	w.Refresh()
}

// SetMinSize sets the minimum size of the image.
func (w *TappableImage) SetMinSize(size fyne.Size) {
	w.minSize = size
	if w.image != nil {
		w.image.SetMinSize(size)
		w.image.Refresh()
	}
}

// SetResource sets the resource of the image.
func (w *TappableImage) SetResource(r fyne.Resource) {
	w.Resource = r
	w.Refresh()
}

func (w *TappableImage) effectiveTranslucency() float64 {
	if !w.Disabled() {
		return w.Translucency
	}
	return w.Translucency + disabledImageTranslucency*(1-w.Translucency)
}

// SetMenuItems replaces the menu items and adds a menu if needed.
// Clears OnTapped so taps show the menu.
func (w *TappableImage) SetMenuItems(menuItems []*fyne.MenuItem) {
	if w.menu == nil {
		w.menu = fyne.NewMenu("")
	}
	w.menu.Items = menuItems
	w.OnTapped = nil
	w.menu.Refresh()
}

func (w *TappableImage) Tapped(pe *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	w.pos = pe.AbsolutePosition
	if w.OnTapped != nil {
		w.OnTapped()
		return
	}
	if w.menu != nil {
		w.showMenu()
	}
}

func (w *TappableImage) showMenu() {
	if len(w.menu.Items) == 0 {
		return
	}
	m := widget.NewPopUpMenu(w.menu, fyne.CurrentApp().Driver().CanvasForObject(w.super()))
	if m == nil {
		return // not on a canvas
	}
	m.ShowAtPosition(w.pos)
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
	if w.Disabled() || w.image == nil {
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
	if w.image == nil {
		w.image = canvas.NewImageFromResource(w.Resource)
		w.image.FillMode = w.FillMode
		w.image.ScaleMode = w.ScaleMode
		w.image.CornerRadius = w.CornerRadius
		w.image.SetMinSize(w.minSize)
		w.image.Translucency = w.effectiveTranslucency()
	}
	return newTappableImageRenderer(w)
}

type tappableImageRenderer struct {
	widget *TappableImage
}

var _ fyne.WidgetRenderer = (*tappableImageRenderer)(nil)

func newTappableImageRenderer(w *TappableImage) *tappableImageRenderer {
	return &tappableImageRenderer{widget: w}
}

func (r *tappableImageRenderer) Destroy() {
}

func (r *tappableImageRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.widget.image}
}

func (r *tappableImageRenderer) Layout(size fyne.Size) {
	pad := theme.Padding()
	r.widget.image.Move(fyne.NewPos(pad, pad))
	r.widget.image.Resize(fyne.NewSize(size.Width-2*pad, size.Height-2*pad))
}

func (r *tappableImageRenderer) MinSize() fyne.Size {
	pad := theme.Padding()
	imgMin := r.widget.image.MinSize()
	return fyne.NewSize(imgMin.Width+2*pad, imgMin.Height+2*pad)
}

func (r *tappableImageRenderer) Refresh() {
	w := r.widget
	w.image.Resource = w.Resource
	w.image.FillMode = w.FillMode
	w.image.ScaleMode = w.ScaleMode
	w.image.CornerRadius = w.CornerRadius
	w.image.Translucency = w.effectiveTranslucency()
	w.image.Refresh()
	canvas.Refresh(w.super())
}
