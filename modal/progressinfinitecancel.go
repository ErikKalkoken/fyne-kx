package modal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ProgressInfiniteCancelModal is a modal that shows an infinite progress indicator while a function is running.
// The modal has a button for canceling the function.
type ProgressInfiniteCancelModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action   func(chan struct{}) error
	canceled chan struct{}
	d        *dialog.CustomDialog
	pb       *widget.ProgressBarInfinite
}

// NewProgressInfiniteWithCancel returns a new [ProgressInfiniteCancelModal] instance.
// The action function needs to check the canceled channel and abort if it is closed.
func NewProgressInfiniteWithCancel(
	title, message string, action func(canceled chan struct{}) error, parent fyne.Window,
) *ProgressInfiniteCancelModal {
	m := &ProgressInfiniteCancelModal{
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
func (m *ProgressInfiniteCancelModal) Show() {
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
