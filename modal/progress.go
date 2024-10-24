package modal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ProgressModal is a modal that shows a progress indicator while a function is running.
// The progress indicator is updated by the function.
type ProgressModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action func(binding.Float) error
	d      *dialog.CustomDialog
	pb     *widget.ProgressBar
	pg     binding.Float
}

// NewProgress returns a new [ProgressModal] instance.
func NewProgress(title, message string, action func(progress binding.Float) error, max float64, parent fyne.Window) *ProgressModal {
	m := &ProgressModal{
		action: action,
		pg:     binding.NewFloat(),
	}
	m.pb = widget.NewProgressBarWithData(m.pg)
	m.pb.Max = max
	content := container.NewVBox(widget.NewLabel(message), m.pb)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Show shows the modal and runs the action function.
func (m *ProgressModal) Show() {
	m.d.Show()
	go func() {
		err := m.action(m.pg)
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