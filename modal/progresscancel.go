package modal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

/*
ProgressCancelModal is a modal that shows a progress indicator while a function is running and which can be canceled.

Example:

	package main

	import (
		"fmt"
		"log"
		"time"

		kmodal "github.com/ErikKalkoken/fyne-kx/modal"

		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/app"
		"fyne.io/fyne/v2/container"
		"fyne.io/fyne/v2/widget"
	)

	func main() {
		a := app.New()
		w := a.NewWindow("Hello World")

		b := widget.NewButton("Delete B", func() {
			timer := time.NewTimer(3 * time.Second)
			m := kmodal.NewProgressModalWithCancel(
				"Deleting", "Deleting B...",
				func(canceled chan struct{}) error {
					select {
					case <-canceled:
						return fmt.Errorf("canceled")
					case <-timer.C:
					}
					return nil
				}, w)
			m.OnSuccess = func() {
				log.Println("completed")
			}
			m.OnError = func(err error) {
				log.Println(err)
			}
			m.Show()
		})

		w.SetContent(container.NewCenter(b))
		w.Resize(fyne.Size{Width: 500, Height: 300})
		w.ShowAndRun()
	}
*/
type ProgressCancelModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	pb       *widget.ProgressBarInfinite
	d        *dialog.CustomDialog
	action   func(chan struct{}) error
	canceled chan struct{}
}

// NewProgressCancelModal returns a new [ProgressCancelModal] instance.
// The action function needs to check the canceled channel and abort if it is closed.
func NewProgressCancelModal(
	title, message string, action func(canceled chan struct{}) error, parent fyne.Window,
) *ProgressCancelModal {
	m := &ProgressCancelModal{
		action: action,
		pb:     widget.NewProgressBarInfinite(),
	}
	content := container.NewVBox(
		widget.NewLabel(message),
		m.pb,
		container.NewPadded(),
		container.NewCenter(widget.NewButton("Cancel", func() {
			close(m.canceled)
		})))
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Show shows the modal and runs the action function.
func (m *ProgressCancelModal) Show() {
	m.canceled = make(chan struct{})
	m.pb.Show()
	m.d.Show()
	go func() {
		err := m.action(m.canceled)
		m.d.Hide()
		if err != nil {
			if m.OnError != nil {
				m.OnError(err)
			}
		} else {
			select {
			case <-m.canceled:
			default:
				close(m.canceled)
			}
			if m.OnSuccess != nil {
				m.OnSuccess()
			}
		}
	}()
}
