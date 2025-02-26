package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// NewToolbarActionMenu returns a ToolBarAction with a context menu.
func NewToolbarActionMenu(icon fyne.Resource, menu *fyne.Menu) *widget.ToolbarAction {
	a := widget.NewToolbarAction(icon, nil)
	o := a.ToolbarObject()
	a.OnActivated = func() {
		showContextMenu(o, menu)
	}
	return a
}

func showContextMenu(o fyne.CanvasObject, menu *fyne.Menu) {
	c := fyne.CurrentApp().Driver().CanvasForObject(o)
	m := widget.NewPopUpMenu(menu, c)
	right := o.Position().X > c.Size().Width/2
	bottom := o.Position().Y > c.Size().Height/2
	var x, y float32
	if right {
		x = -m.Size().Width + o.Size().Width
	} else {
		x = 0
	}
	if bottom {
		y = -m.Size().Height - o.Size().Height
	} else {
		y = o.Size().Height
	}
	m.ShowAtRelativePosition(fyne.NewPos(x, y), o)
}
