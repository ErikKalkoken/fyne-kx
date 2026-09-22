package widget

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// pollInterval and pollTimeout bound how long tests wait for the Snackbar's
// background goroutine (and its fyne.Do-queued UI updates) to reach an
// expected state in real time, since there is no virtual clock here.
const (
	pollInterval = 10 * time.Millisecond
	pollTimeout  = time.Second
)

func TestSnackbar_LifecycleAndTimeout(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	require.False(t, sb.popup.Visible(), "expected snackbar popup to be hidden initially")

	sb.Display("Hello, World!")
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval,
		"expected snackbar popup to be visible after Display")

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, snackbarTimeoutDefault+pollTimeout, pollInterval,
		"expected snackbar popup to auto-hide after default timeout")
}

func TestSnackbar_CustomTimeout(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	customTimeout := 300 * time.Millisecond
	sb.DisplayWithTimeout("Custom Timeout Message", customTimeout)
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval, "expected snackbar popup to be visible")

	// Stay short of the timeout.
	time.Sleep(customTimeout / 2)
	assert.True(t, sb.popup.Visible(), "expected snackbar popup to still be visible before timeout")

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, customTimeout+pollTimeout, pollInterval,
		"expected snackbar popup to hide after custom timeout")
}

func TestSnackbar_ManualDismissByTap(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	sb.Display("Tap me to dismiss")
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval, "expected snackbar popup to be visible")

	// Simulate user tapping the popup overlay.
	sb.popup.Tapped(&fyne.PointEvent{})
	require.Eventually(t, func() bool { return !sb.popup.Visible() }, pollTimeout, pollInterval,
		"expected snackbar popup to hide after tap")
}

func TestSnackbar_SequentialQueueing(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	// Enqueue two messages.
	sb.DisplayWithTimeout("Message 1", 200*time.Millisecond)
	sb.DisplayWithTimeout("Message 2", 200*time.Millisecond)

	require.Eventually(t, func() bool {
		return sb.popup.Visible() && sb.text.String() == "Message 1"
	}, pollTimeout, pollInterval, "expected first message to be shown")

	require.Eventually(t, func() bool {
		return sb.popup.Visible() && sb.text.String() == "Message 2"
	}, pollTimeout, pollInterval, "expected snackbar to show second message after first times out")

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, pollTimeout, pollInterval,
		"expected snackbar to hide after second message completes")
}

func TestSnackbar_QueueingBeforeStart(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")

	sb := NewSnackbar(window.Canvas())

	// Messages queued while stopped should sit in the queue.
	sb.DisplayWithTimeout("Queued Early", 100*time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	assert.False(t, sb.popup.Visible(), "expected snackbar to remain hidden before Start() is called")

	// Starting the snackbar should pick up queued messages.
	sb.Start()
	defer sb.Stop()

	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval,
		"expected snackbar to display queued message after Start()")

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, pollTimeout, pollInterval,
		"expected snackbar to hide after processing pre-queued item")
}

func TestSnackbar_StopAndRestart(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")

	sb := NewSnackbar(window.Canvas())
	sb.Start()

	sb.DisplayWithTimeout("Going to stop", 500*time.Millisecond)
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval, "expected snackbar to show active message")

	// Stop aborts current context.
	sb.Stop()
	require.Eventually(t, func() bool { return !sb.isRunning.Load() }, pollTimeout, pollInterval,
		"expected isRunning to be false after Stop()")

	// Restart and show new message.
	sb.Start()
	defer sb.Stop()

	sb.DisplayWithTimeout("Restarted Message", 200*time.Millisecond)
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval,
		"expected snackbar to function normally after restart")
}

func TestSnackbar_MessageContent(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	sb.DisplayWithTimeout("First message", 200*time.Millisecond)
	require.Eventually(t, func() bool { return sb.text.String() == "First message" }, pollTimeout, pollInterval)

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, pollTimeout, pollInterval)

	sb.DisplayWithTimeout("Second message", 200*time.Millisecond)
	require.Eventually(t, func() bool { return sb.text.String() == "Second message" }, pollTimeout, pollInterval)
}

func TestSnackbar_ConcurrentDisplay(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(400, 300))

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	const n = 10
	const timeout = 50 * time.Millisecond

	var wg sync.WaitGroup
	wg.Add(n)
	for i := range n {
		go func(i int) {
			defer wg.Done()
			sb.DisplayWithTimeout(fmt.Sprintf("message-%d", i), timeout)
		}(i)
	}
	wg.Wait()

	// Sample more often than the per-message timeout so no message's display
	// window can be skipped over, then keep sampling past the point where all
	// messages must have been processed.
	seen := make(map[string]bool)
	deadline := time.Now().Add(time.Duration(n)*timeout + pollTimeout)
	for time.Now().Before(deadline) && len(seen) < n {
		if sb.popup.Visible() {
			seen[sb.text.String()] = true
		}
		time.Sleep(timeout / 3)
	}

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, pollTimeout, pollInterval,
		"expected popup hidden after processing all concurrently queued messages")
	assert.Len(t, seen, n, "expected %d distinct messages to be displayed, got %v", n, seen)
}

func TestSnackbar_TextWrappingCalculation(t *testing.T) {
	app := test.NewTempApp(t)

	window := app.NewWindow("Test Window")
	window.Resize(fyne.NewSize(200, 300)) // Narrow canvas to force text wrapping

	sb := NewSnackbar(window.Canvas())
	sb.Start()
	defer sb.Stop()

	// Short text shouldn't trigger word wrap.
	sb.DisplayWithTimeout("Hi", 100*time.Millisecond)
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval)
	assert.Equal(t, fyne.TextWrapOff, sb.text.Wrapping)

	require.Eventually(t, func() bool { return !sb.popup.Visible() }, pollTimeout, pollInterval)

	// Very long text should trigger word wrap.
	longText := "This is an extremely long message that will exceed the width of the canvas and force word wrapping"
	sb.DisplayWithTimeout(longText, 100*time.Millisecond)
	require.Eventually(t, sb.popup.Visible, pollTimeout, pollInterval)
	assert.Equal(t, fyne.TextWrapWord, sb.text.Wrapping)
}
