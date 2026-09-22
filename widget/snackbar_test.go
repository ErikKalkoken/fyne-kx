package widget_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	"github.com/ErikKalkoken/fyne-kx/widget"
)

// runWithTimeout runs fn in a goroutine and fails the test instead of
// hanging forever if fn does not return within timeout. Snackbar's public
// API exposes no way to observe its internal state, so black-box tests can
// only assert that calling it doesn't panic or deadlock.
func runWithTimeout(t *testing.T, timeout time.Duration, fn func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		fn()
	}()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Fatal("timed out, suspected deadlock")
	}
}

func TestSnackbar_PublicAPILifecycleDoesNotPanicOrDeadlock(t *testing.T) {
	app := test.NewTempApp(t)
	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	assert.NotPanics(t, func() {
		runWithTimeout(t, 2*time.Second, func() {
			sb := widget.NewSnackbar(window.Canvas())
			sb.BottomMargin = 20
			sb.Start()
			sb.Display("Hello, World!")
			sb.DisplayWithTimeout("Custom timeout", 20*time.Millisecond)
			time.Sleep(50 * time.Millisecond)
			sb.Stop()
		})
	})
}

func TestSnackbar_DisplayBeforeStart_QueuesWithoutPanic(t *testing.T) {
	app := test.NewTempApp(t)
	window := app.NewWindow("Test Window")

	assert.NotPanics(t, func() {
		runWithTimeout(t, 2*time.Second, func() {
			sb := widget.NewSnackbar(window.Canvas())
			sb.Display("Queued before start")
			sb.Start()
			time.Sleep(50 * time.Millisecond)
			sb.Stop()
		})
	})
}

func TestSnackbar_StartCalledTwice_SecondCallIsNoop(t *testing.T) {
	app := test.NewTempApp(t)
	window := app.NewWindow("Test Window")

	assert.NotPanics(t, func() {
		runWithTimeout(t, 2*time.Second, func() {
			sb := widget.NewSnackbar(window.Canvas())
			sb.Start()
			sb.Start() // Second call should just log a warning.
			sb.Stop()
		})
	})
}

func TestSnackbar_ConcurrentDisplayCalls_DoNotRace(t *testing.T) {
	app := test.NewTempApp(t)
	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := widget.NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	const n = 10
	assert.NotPanics(t, func() {
		runWithTimeout(t, 2*time.Second, func() {
			var wg sync.WaitGroup
			wg.Add(n)
			for i := range n {
				go func(i int) {
					defer wg.Done()
					sb.DisplayWithTimeout(fmt.Sprintf("message-%d", i), 10*time.Millisecond)
				}(i)
			}
			wg.Wait()
		})
	})
}
