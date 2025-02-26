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
	onRightHalf := o.Position().X > c.Size().Width/2
	var x float32
	if onRightHalf {
		x = -m.Size().Width + o.Size().Width
	} else {
		x = 0
	}
	m.ShowAtRelativePosition(fyne.NewPos(x, o.Size().Height), o)
}
