/*
Package modal defines modals for the Fyne GUI toolkit.

# Modals

Modals are similar to Fyne dialogs, but do not require user interaction.
They are useful when you have a longer running process that the user needs to wait for before he can continue. e.g. opening a large file.

# Progress modals

Progress modals are modals that show a progress indicator while an action function is running.
The are several variant, which all share a similar API:
  - Title and message
  - Action function callback that return an error
  - Callback hooks for success and error, e.g. to inform the user about an error
  - Start() method is called to start the action

Note that the action function will always be run as a goroutine.

A progress modal can be used similar to Fyne dialogs:

	m := kxmodal.NewProgressInfinite("Loading file", "Loading file XX. Please wait.", func() error {
		time.Sleep(3 * time.Second)  // simulate a long running process
		return nil
	}, w)
	m.Start()
*/
package modal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ProgressModal is a modal that shows a progress indicator while an action function is running.
// The progress indicator must be updated by the action function.
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

// Start starts the action function and shows the modal while it is running.
func (m *ProgressModal) Start() {
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

// ProgressCancelModal is a modal that shows a progress indicator while a function is running.
// The progress indicator is updated by the function.
type ProgressCancelModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action   func(binding.Float, chan struct{}) error
	canceled chan struct{}
	d        *dialog.CustomDialog
	pb       *widget.ProgressBar
	pg       binding.Float
}

// NewProgress returns a new [ProgressModal] instance.
func NewProgressWithCancel(title, message string, action func(progress binding.Float, canceled chan struct{}) error, max float64, parent fyne.Window) *ProgressCancelModal {
	m := &ProgressCancelModal{
		action: action,
		pg:     binding.NewFloat(),
	}
	m.pb = widget.NewProgressBarWithData(m.pg)
	m.pb.Max = max
	content := container.NewVBox(
		widget.NewLabel(message),
		m.pb,
		container.NewPadded(),
		container.NewCenter(widget.NewButton("Cancel", func() {
			closeChannelIfOpen(m.canceled)
		})))
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Start starts the action function and shows the modal while it is running.
func (m *ProgressCancelModal) Start() {
	m.canceled = make(chan struct{})
	m.d.Show()
	go func() {
		err := m.action(m.pg, m.canceled)
		m.d.Hide()
		if err != nil {
			if m.OnError != nil {
				m.OnError(err)
			}
		} else {
			closeChannelIfOpen(m.canceled)
			if m.OnSuccess != nil {
				m.OnSuccess()
			}
		}
	}()
}

func closeChannelIfOpen(c chan struct{}) {
	select {
	case <-c:
	default:
		close(c)
	}
}

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

// Start starts the action function and shows the modal while it is running.
func (m *ProgressInfiniteModal) Start() {
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
			closeChannelIfOpen(m.canceled)
		})))
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Start starts the action function and shows the modal while it is running.
func (m *ProgressInfiniteCancelModal) Start() {
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
			closeChannelIfOpen(m.canceled)
			if m.OnSuccess != nil {
				m.OnSuccess()
			}
		}
	}()
}
