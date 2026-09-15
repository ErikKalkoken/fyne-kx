/*
Package modal defines modals for the Fyne GUI toolkit.

# Modals

Modals are similar to Fyne dialogs, but do not require user interaction.
They are useful when you have a longer running process that the user needs to wait for
before they can continue. e.g. opening a large file.

# Progress modals

Progress modals are modals that show a progress indicator while an action function is running.
The are several variant, which all share a similar API:
  - Title and message
  - Action function callback
  - Callback hooks for success and error, e.g. to inform the user about an error
  - Start() method is called to start the action

The action function runs on the calling goroutine, consistent with how Fyne invokes other
widget callbacks (Start() is expected to be called from the main goroutine, e.g. from a
button's OnTapped). It receives a done function that must be called when the action has
finished, with a nil error on success or a non-nil error on failure. done is safe to call
from any goroutine, and more than once (only the first call has an effect). If the action
does long running work it must spawn its own goroutine to avoid blocking the UI.

Cancelable variants additionally receive an onCancel function. The action should call it
once, as early as convenient, with a function that aborts the running work (e.g. canceling
a context.Context). onCancel is safe to call from any goroutine, and works no matter
whether it is called before or after the user presses the Cancel button.

A progress modal can be used similar to Fyne dialogs. See the examples for [NewProgress],
[NewProgressWithCancel], [NewProgressInfinite] and [NewProgressInfiniteWithCancel] for
basic usage.
*/
package modal

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	kxdialog "github.com/ErikKalkoken/fyne-kx/dialog"
	"github.com/ErikKalkoken/fyne-kx/internal/stack"
)

// Keeping a stack of opened dialogs to make sure they are closed in LIFO order
// For more information see Fyne issue #5564
var openDialogs *stack.Stack[*dialog.CustomDialog]

func init() {
	openDialogs = stack.New[*dialog.CustomDialog]()
}

// ProgressModal is a modal that shows a progress indicator while an action function is running.
// The progress indicator must be updated by the action function.
type ProgressModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action  func(progress binding.Float, done func(error))
	d       *dialog.CustomDialog
	pb      *widget.ProgressBar
	pg      binding.Float
	started bool
}

// NewProgress returns a new [ProgressModal] instance.
func NewProgress(title, message string, action func(progress binding.Float, done func(error)), max float64, parent fyne.Window) *ProgressModal {
	m := &ProgressModal{
		action: action,
		pg:     binding.NewFloat(),
	}
	pb := widget.NewProgressBarWithData(m.pg)
	pb.Max = max
	m.pb = pb
	content := container.NewVBox(widget.NewLabel(message), m.pb)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Start starts the action function and shows the modal while it is running.
// A modal can be started once only. Must be called from the main goroutine.
func (m *ProgressModal) Start() {
	if m.started {
		return
	}
	m.started = true
	startAction(m.d, m.OnSuccess, m.OnError, func(done func(error)) {
		m.action(m.pg, done)
	})
}

// ProgressCancelModal is a modal that shows a progress indicator while a function is running.
// The progress indicator is updated by the function.
type ProgressCancelModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action  func(progress binding.Float, onCancel func(func()), done func(error))
	cancel  *cancelRegistry
	d       *dialog.CustomDialog
	pb      *widget.ProgressBar
	pg      binding.Float
	started bool
}

// NewProgressWithCancel returns a new [ProgressCancelModal] instance.
func NewProgressWithCancel(title, message string, action func(progress binding.Float, onCancel func(func()), done func(error)), max float64, parent fyne.Window) *ProgressCancelModal {
	m := &ProgressCancelModal{
		action: action,
		pg:     binding.NewFloat(),
		cancel: &cancelRegistry{},
	}
	pb := widget.NewProgressBarWithData(m.pg)
	pb.Max = max
	m.pb = pb
	content := container.NewVBox(
		widget.NewLabel(message),
		m.pb,
		container.NewPadded(),
		container.NewCenter(widget.NewButton("Cancel", func() {
			m.cancel.requestCancel()
		})),
	)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	kxdialog.AddDialogKeyHandler(m.d, parent)
	return m
}

// Start starts the action function and shows the modal while it is running.
// A modal can be started once only. Must be called from the main goroutine.
func (m *ProgressCancelModal) Start() {
	if m.started {
		return
	}
	m.started = true
	startAction(m.d, m.OnSuccess, m.OnError, func(done func(error)) {
		m.action(m.pg, m.cancel.register, done)
	})
}

// ProgressInfiniteModal is a modal that shows an infinite progress indicator while a function is running.
type ProgressInfiniteModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action  func(done func(error))
	d       *dialog.CustomDialog
	pb      *widget.ProgressBarInfinite
	started bool
}

// NewProgressInfinite returns a new [ProgressInfiniteModal] instance.
func NewProgressInfinite(title, message string, action func(done func(error)), parent fyne.Window) *ProgressInfiniteModal {
	m := &ProgressInfiniteModal{
		action: action,
		pb:     widget.NewProgressBarInfinite(),
	}
	content := container.NewVBox(widget.NewLabel(message), m.pb)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	return m
}

// Start starts the action function and shows the modal while it is running.
// A modal can be started once only. Must be called from the main goroutine.
func (m *ProgressInfiniteModal) Start() {
	if m.started {
		return
	}
	m.started = true
	startAction(m.d, m.OnSuccess, m.OnError, m.action)
}

// ProgressInfiniteCancelModal is a modal that shows an infinite progress indicator while a function is running.
// The modal has a button for canceling the function.
type ProgressInfiniteCancelModal struct {
	// Optional callback when the action failed.
	OnError func(err error)

	// Optional callback when the action succeeded.
	OnSuccess func()

	action  func(onCancel func(func()), done func(error))
	cancel  *cancelRegistry
	d       *dialog.CustomDialog
	pb      *widget.ProgressBarInfinite
	started bool
}

// NewProgressInfiniteWithCancel returns a new [ProgressInfiniteCancelModal] instance.
func NewProgressInfiniteWithCancel(
	title, message string, action func(onCancel func(func()), done func(error)), parent fyne.Window,
) *ProgressInfiniteCancelModal {
	m := &ProgressInfiniteCancelModal{
		action: action,
		pb:     widget.NewProgressBarInfinite(),
		cancel: &cancelRegistry{},
	}
	content := container.NewVBox(
		widget.NewLabel(message),
		m.pb,
		container.NewPadded(),
		container.NewCenter(widget.NewButton("Cancel", func() {
			m.cancel.requestCancel()
		})),
	)
	m.d = dialog.NewCustomWithoutButtons(title, content, parent)
	kxdialog.AddDialogKeyHandler(m.d, parent)
	return m
}

// Start starts the action function and shows the modal while it is running.
// A modal can be started once only. Must be called from the main goroutine.
func (m *ProgressInfiniteCancelModal) Start() {
	if m.started {
		return
	}
	m.started = true
	startAction(m.d, m.OnSuccess, m.OnError, func(done func(error)) {
		m.action(m.cancel.register, done)
	})
}

// startAction shows d, then calls run on the calling goroutine, consistent with how Fyne
// invokes other widget callbacks. run receives a done function that reports the outcome of
// the action; done is safe to call from any goroutine, and more than once (only the first
// call has an effect).
func startAction(d *dialog.CustomDialog, onSuccess func(), onError func(error), run func(done func(error))) {
	openDialogs.Push(d)
	d.Show()
	var once sync.Once
	done := func(err error) {
		once.Do(func() {
			fyne.Do(func() {
				if dd, ok := openDialogs.Pop(); ok {
					dd.Hide()
				} else {
					fyne.LogError("Failed to hide dialog of progress modal", nil)
				}
				if err != nil {
					if onError != nil {
						onError(err)
					}
				} else if onSuccess != nil {
					onSuccess()
				}
			})
		})
	}
	run(done)
}

// cancelRegistry coordinates a Cancel button with an action's abort handler, regardless of
// whether the action registers its handler before or after the button is pressed.
type cancelRegistry struct {
	mu       sync.Mutex
	handler  func()
	canceled bool
}

// register stores f as the abort handler to run when Cancel is pressed. If Cancel was
// already pressed, f runs immediately instead. Safe to call from any goroutine.
func (c *cancelRegistry) register(f func()) {
	c.mu.Lock()
	already := c.canceled
	if !already {
		c.handler = f
	}
	c.mu.Unlock()
	if already && f != nil {
		f()
	}
}

// requestCancel runs the registered abort handler, if any has been registered yet. Safe to
// call from any goroutine, and more than once (only the first call has an effect).
func (c *cancelRegistry) requestCancel() {
	c.mu.Lock()
	if c.canceled {
		c.mu.Unlock()
		return
	}
	c.canceled = true
	h := c.handler
	c.mu.Unlock()
	if h != nil {
		h()
	}
}
