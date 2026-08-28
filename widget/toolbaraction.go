package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// NewToolbarActionMenu returns a ToolBarAction with a context menu.
func NewToolbarActionMenu(icon fyne.Resource, menu *fyne.Menu) *widget.ToolbarAction {
	a := widget.NewToolbarAction(icon, nil)
	w := a.ToolbarObject()
	a.OnActivated = func() {
		showPopUpMenuBelowLeading(w, menu)
	}
	return a
}
