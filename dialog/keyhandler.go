// Package dialog contains helpers for building dialogs with the Fyne GUI toolkit.
package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// AddDialogKeyHandler adds a key handler to a dialog.
// It enables the user to close the dialog by pressing the escape key.
//
// Note that previously defined key events will be deactivated while the dialog is open.
func AddDialogKeyHandler(d dialog.Dialog, w fyne.Window) {
	originalEvent := w.Canvas().OnTypedKey() // == nil when not set
	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if d == nil {
			return
		}
		if ke.Name == fyne.KeyEscape {
			d.Hide()
		}
	})
	d.SetOnClosed(func() {
		w.Canvas().SetOnTypedKey(originalEvent)
	})
}
