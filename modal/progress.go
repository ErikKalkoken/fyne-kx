// Package modal defines modals for the Fyne GUI toolkit.
package modal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

/*
ProgressModal is a modal with that shows a progress indicator while a function is running.

Example:

	package main

	import (
		"log"
		"time"

		kxmodal "github.com/ErikKalkoken/fyne-kx/modal"

		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/app"
		"fyne.io/fyne/v2/container"
		"fyne.io/fyne/v2/widget"
	)

	func main() {
		a := app.New()
		w := a.NewWindow("Modal example")

		b := widget.NewButton("Delete A", func() {
			m := kxmodal.NewProgressModal("Deleting", "Deleting A...", func() error {
				time.Sleep(2 * time.Second) // simulating work
				return nil
			}, w)
			m.Show()
		})

		w.SetContent(container.NewCenter(b))
		w.Resize(fyne.Size{Width: 500, Height: 300})
		w.ShowAndRun()
	}
*/
type ProgressModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	pb     *widget.ProgressBarInfinite
	d      *dialog.CustomDialog
	action func() error
}

// NewProgressModal returns a new [ProgressModal] instance.
func NewProgressModal(title, message string, action func() error, parent fyne.Window) *ProgressModal {
	m := &ProgressModal{
		action: action,
		pb:     widget.NewProgressBarInfinite(),
	}
	content := container.NewVBox(widget.NewLabel(message), m.pb)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Show shows the modal and runs the action function.
func (m *ProgressModal) Show() {
	m.pb.Start()
	m.d.Show()
	go func() {
		err := m.action()
		m.d.Hide()
		if err != nil {
			if m.OnError != nil {
				m.OnError(err)
			}
		} else {
			if m.OnSuccess != nil {
				m.OnSuccess()
			}
		}
	}()
}
