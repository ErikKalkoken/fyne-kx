package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO: Add hover shadow

// IconButton is a widget which help people take minor actions with one tap.
type IconButton struct {
	widget.DisableableWidget

	// This callback runs when the icon is tapped.
	OnTapped func()

	hovered          bool
	menu             *fyne.Menu
	resource         fyne.Resource
	resourceDisabled fyne.Resource
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
	w.OnTapped = func() {
		if w.menu == nil || len(w.menu.Items) == 0 {
			return
		}
		m := widget.NewPopUpMenu(menu, fyne.CurrentApp().Driver().CanvasForObject(w))
		m.ShowAtRelativePosition(
			fyne.NewPos(
				-m.Size().Width+w.Size().Width,
				w.Size().Height,
			),
			w,
		)
	}
	return w
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

// SetMenuItems replaces the menu items.
// Does nothing when the widget has not bee created with [NewIconButtonWithMenu].
func (w *IconButton) SetMenuItems(menuItems []*fyne.MenuItem) {
	if w.menu == nil {
		return
	}
	w.menu.Items = menuItems
	w.Refresh()
}

func (w *IconButton) Refresh() {
	if w.menu != nil {
		w.menu.Refresh()
	}
	w.BaseWidget.Refresh()
}

func (w *IconButton) Tapped(_ *fyne.PointEvent) {
	if !w.Disabled() && w.OnTapped != nil {
		w.OnTapped()
	}
}

func (w *IconButton) TappedSecondary(_ *fyne.PointEvent) {
}

// Cursor returns the cursor type of this widget
func (w *IconButton) Cursor() desktop.Cursor {
	if w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// MouseIn is a hook that is called if the mouse pointer enters the element.
func (w *IconButton) MouseIn(_ *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.hovered = true
}

func (w *IconButton) MouseMoved(_ *desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *IconButton) MouseOut() {
	w.hovered = false
}

func (w *IconButton) CreateRenderer() fyne.WidgetRenderer {
	return newIconButtonRenderer(w)
}

// iconButtonRenderer is a custom [fyne.WidgetRenderer] for [IconButton].
//
// It owns the rendered icon image and lays it out with a themed padding on
// every side, replicating the behavior previously provided by wrapping the
// icon in a [container.NewPadded].
type iconButtonRenderer struct {
	button *IconButton
	icon   *canvas.Image
}

var _ fyne.WidgetRenderer = (*iconButtonRenderer)(nil)

func newIconButtonRenderer(w *IconButton) *iconButtonRenderer {
	i := canvas.NewImageFromResource(w.resource)
	i.FillMode = canvas.ImageFillContain
	i.SetMinSize(fyne.NewSquareSize(theme.Size(theme.SizeNameInlineIcon)))
	r := &iconButtonRenderer{
		button: w,
		icon:   i,
	}
	r.updateState()
	return r
}

func (r *iconButtonRenderer) Destroy() {
}

func (r *iconButtonRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.icon}
}

func (r *iconButtonRenderer) Layout(size fyne.Size) {
	pad := theme.Padding()
	innerSize := fyne.NewSize(size.Width-2*pad, size.Height-2*pad)
	r.icon.Resize(innerSize)
	r.icon.Move(fyne.NewPos(pad, pad))
}

func (r *iconButtonRenderer) MinSize() fyne.Size {
	pad := theme.Padding()
	iconMin := r.icon.MinSize()
	return fyne.NewSize(iconMin.Width+2*pad, iconMin.Height+2*pad)
}

func (r *iconButtonRenderer) Refresh() {
	r.updateState()
	r.icon.Refresh()
	canvas.Refresh(r.button)
}

// updateState sets the icon's resource depending on whether the button is
// enabled or disabled.
func (r *iconButtonRenderer) updateState() {
	if r.button.Disabled() {
		r.icon.Resource = r.button.resourceDisabled
	} else {
		r.icon.Resource = r.button.resource
	}
}
