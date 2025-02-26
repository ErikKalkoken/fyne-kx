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
		c := fyne.CurrentApp().Driver().CanvasForObject(o)
		m := widget.NewPopUpMenu(menu, c)
		m.ShowAtRelativePosition(fyne.Position{}, o)
	}
	return a
}
