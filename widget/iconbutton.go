package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// IconButton is a widget which help people take minor actions with one tap.
// It shows a hover background when tapping would do something.
//
// It can be extended by embedding it by value and calling ExtendBaseWidget.
// Overrides of ExtendBaseWidget and Refresh must call the IconButton versions.
// When combined with another hoverable type (e.g. a tooltip extension),
// MouseIn, MouseOut and MouseMoved must be forwarded to both.
type IconButton struct {
	widget.DisableableWidget

	// This callback runs when the icon is tapped. Takes precedence over the menu.
	OnTapped func()

	hovered          bool
	menu             *fyne.Menu
	resource         fyne.Resource
	resourceDisabled fyne.Resource
	self             fyne.Widget // outermost widget, differs from w when embedded
}

var _ fyne.Tappable = (*IconButton)(nil)
var _ desktop.Hoverable = (*IconButton)(nil)

// NewIconButton returns a new instance of an [IconButton].
func NewIconButton(icon fyne.Resource, tapped func()) *IconButton {
	w := &IconButton{
		OnTapped: tapped,
	}
	w.ExtendBaseWidget(w)
	w.setIconResource(icon)
	return w
}

// NewIconButtonWithMenu returns an [IconButton] with a context menu.
func NewIconButtonWithMenu(icon fyne.Resource, menu *fyne.Menu) *IconButton {
	w := NewIconButton(icon, nil)
	if menu == nil {
		fyne.LogError("IconButton misconfigured: missing menu", nil)
		return w
	}
	w.menu = menu
	return w
}

// ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.
func (w *IconButton) ExtendBaseWidget(wid fyne.Widget) {
	w.self = wid
	w.DisableableWidget.ExtendBaseWidget(wid)
}

// super returns the outermost widget, which Fyne knows as part of the canvas tree.
func (w *IconButton) super() fyne.CanvasObject {
	if w.self == nil {
		return w
	}
	return w.self
}

// SetIcon replaces the current icon.
func (w *IconButton) SetIcon(icon fyne.Resource) {
	w.setIconResource(icon)
	w.Refresh()
}

func (w *IconButton) setIconResource(icon fyne.Resource) {
	w.resource = icon
	if isResourceSVG(icon) {
		w.resourceDisabled = theme.NewDisabledResource(icon)
	} else {
		w.resourceDisabled = icon
	}
}

// SetMenuItems replaces the menu items and adds a menu if needed.
// Clears OnTapped so taps show the menu.
func (w *IconButton) SetMenuItems(menuItems []*fyne.MenuItem) {
	if w.menu == nil {
		w.menu = fyne.NewMenu("")
	}
	w.menu.Items = menuItems
	w.OnTapped = nil
	w.Refresh()
}

func (w *IconButton) Refresh() {
	if w.menu != nil {
		w.menu.Refresh()
	}
	w.BaseWidget.Refresh()
}

func (w *IconButton) Tapped(_ *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	if w.OnTapped != nil {
		w.OnTapped()
		return
	}
	if w.menu != nil {
		w.showMenu()
	}
}

func (w *IconButton) showMenu() {
	if len(w.menu.Items) == 0 {
		return
	}
	o := w.super()
	m := widget.NewPopUpMenu(w.menu, fyne.CurrentApp().Driver().CanvasForObject(o))
	if m == nil {
		return // not on a canvas
	}
	m.ShowAtRelativePosition(
		fyne.NewPos(
			-m.Size().Width+w.Size().Width,
			w.Size().Height,
		),
		o,
	)
}

func (w *IconButton) TappedSecondary(_ *fyne.PointEvent) {
}

func (w *IconButton) isTappable() bool {
	return w.OnTapped != nil || (w.menu != nil && len(w.menu.Items) > 0)
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *IconButton) MouseIn(_ *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.setHovered(true)
}

func (w *IconButton) MouseMoved(_ *desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *IconButton) MouseOut() {
	w.setHovered(false)
}

func (w *IconButton) setHovered(hovered bool) {
	if w.hovered == hovered {
		return
	}
	w.hovered = hovered
	w.BaseWidget.Refresh() // skips refreshing the menu
}

func (w *IconButton) showsHover() bool {
	return !w.Disabled() && w.hovered && w.isTappable()
}

func (w *IconButton) CreateRenderer() fyne.WidgetRenderer {
	return newIconButtonRenderer(w)
}

// iconButtonRenderer is a custom [fyne.WidgetRenderer] for [IconButton].
//
// It shows the icon at a fixed size with a themed inner padding
// and a hover background behind it.
type iconButtonRenderer struct {
	background *canvas.Circle
	button     *IconButton
	icon       *canvas.Image
}

var _ fyne.WidgetRenderer = (*iconButtonRenderer)(nil)

func newIconButtonRenderer(w *IconButton) *iconButtonRenderer {
	i := canvas.NewImageFromResource(w.resource)
	i.FillMode = canvas.ImageFillContain
	r := &iconButtonRenderer{
		background: canvas.NewCircle(color.Transparent),
		button:     w,
		icon:       i,
	}
	r.updateState()
	return r
}

func (r *iconButtonRenderer) Destroy() {
}

func (r *iconButtonRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.icon}
}

// Layout centers icon and background at a fixed size, like Fyne's Button.
func (r *iconButtonRenderer) Layout(size fyne.Size) {
	bg := r.MinSize()
	r.background.Resize(bg)
	r.background.Move(centerIn(size, bg))
	icon := fyne.NewSquareSize(r.iconSize())
	r.icon.Resize(icon)
	r.icon.Move(centerIn(size, icon))
}

func (r *iconButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSquareSize(r.iconSize() + 2*r.padding())
}

func (r *iconButtonRenderer) iconSize() float32 {
	return r.button.Theme().Size(theme.SizeNameInlineIcon)
}

// padding returns the space around the icon, same as Fyne's icon-only buttons.
func (r *iconButtonRenderer) padding() float32 {
	return r.button.Theme().Size(theme.SizeNameInnerPadding)
}

func centerIn(outer, inner fyne.Size) fyne.Position {
	return fyne.NewPos((outer.Width-inner.Width)/2, (outer.Height-inner.Height)/2)
}

func (r *iconButtonRenderer) Refresh() {
	r.updateState()
	r.background.Refresh()
	r.icon.Refresh()
	canvas.Refresh(r.button.super())
}

// updateState applies the disabled and hover state.
func (r *iconButtonRenderer) updateState() {
	if r.button.Disabled() {
		r.icon.Resource = r.button.resourceDisabled
	} else {
		r.icon.Resource = r.button.resource
	}
	if r.button.showsHover() {
		v := fyne.CurrentApp().Settings().ThemeVariant()
		r.background.FillColor = r.button.Theme().Color(theme.ColorNameHover, v)
	} else {
		r.background.FillColor = color.Transparent
	}
}
