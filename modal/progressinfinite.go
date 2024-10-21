// Package modal defines modals for the Fyne GUI toolkit.
package modal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ProgressInfiniteModal is a modal that shows an infinite progress indicator while a function is running.
type ProgressInfiniteModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action func() error
	d      *dialog.CustomDialog
	pb     *widget.ProgressBarInfinite
}

// NewProgressInfinite returns a new [ProgressInfiniteModal] instance.
func NewProgressInfinite(title, message string, action func() error, parent fyne.Window) *ProgressInfiniteModal {
	m := &ProgressInfiniteModal{
		action: action,
		pb:     widget.NewProgressBarInfinite(),
	}
	content := container.NewVBox(widget.NewLabel(message), m.pb)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Show shows the modal and runs the action function.
func (m *ProgressInfiniteModal) Show() {
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
