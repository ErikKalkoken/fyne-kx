package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO: Add hover shadow

// Icon buttons are widget which help people take minor actions with one tap.
type IconButton struct {
	widget.DisableableWidget

	// This callback runs when the icon is tapped.
	OnTapped func()

	hovered          bool
	icon             *canvas.Image
	menu             *fyne.Menu
	resource         fyne.Resource
	resourceDisabled fyne.Resource
}

var _ fyne.Tappable = (*IconButton)(nil)
var _ desktop.Hoverable = (*IconButton)(nil)

// NewIconButton returns a new instance of an [IconButton].
func NewIconButton(icon fyne.Resource, tapped func()) *IconButton {
	// i := NewImageFromResource(icon, fyne.NewSquareSize(theme.Size(theme.SizeNameInlineIcon)))
	i := canvas.NewImageFromResource(icon)
	i.FillMode = canvas.ImageFillContain
	i.SetMinSize(fyne.NewSquareSize(theme.Size(theme.SizeNameInlineIcon)))
	w := &IconButton{
		OnTapped: tapped,
		icon:     i,
	}
	w.ExtendBaseWidget(w)
	w.setIconResource(icon)
	return w
}

// NewIconButtonWithMenu returns an [IconButton] with a context menu.
func NewIconButtonWithMenu(icon fyne.Resource, menu *fyne.Menu) *IconButton {
	w := NewIconButton(icon, nil)
	w.menu = menu
	w.OnTapped = func() {
		if len(w.menu.Items) == 0 {
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
func (w *IconButton) SetMenuItems(menuItems []*fyne.MenuItem) {
	if w.menu == nil {
		return
	}
	w.menu.Items = menuItems
	w.Refresh()
}

func (w *IconButton) Refresh() {
	w.updateState()
	w.icon.Refresh()
	if w.menu != nil {
		w.menu.Refresh()
	}
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
func (w *IconButton) MouseIn(e *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.hovered = true
}

func (w *IconButton) MouseMoved(*desktop.MouseEvent) {
	// needed to satisfy the interface only
}

// MouseOut is a hook that is called if the mouse pointer leaves the element.
func (w *IconButton) MouseOut() {
	w.hovered = false
}

func (w *IconButton) CreateRenderer() fyne.WidgetRenderer {
	w.updateState()
	return widget.NewSimpleRenderer(container.NewPadded(w.icon))
}

func (w *IconButton) updateState() {
	if w.Disabled() {
		w.icon.Resource = w.resourceDisabled
	} else {
		w.icon.Resource = w.resource
	}
}
