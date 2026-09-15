package modal_test

import (
	"errors"
	"testing"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ErikKalkoken/fyne-kx/modal"
)

func TestProgressModal_Success(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	successCalled := make(chan struct{})
	var finalProgress float64

	m := modal.NewProgress("Title", "Message", func(pg binding.Float, done func(error)) {
		err := pg.Set(50.0)
		require.NoError(t, err)
		val, err := pg.Get()
		require.NoError(t, err)
		finalProgress = val
		done(nil)
	}, 100.0, w)

	m.OnSuccess = func() {
		close(successCalled)
	}

	m.Start()

	select {
	case <-successCalled:
		assert.Equal(t, 50.0, finalProgress)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnSuccess callback")
	}
}

func TestProgressModal_ErrorHandling(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	expectedErr := errors.New("something went wrong")
	errChan := make(chan error, 1)

	m := modal.NewProgress("Title", "Message", func(pg binding.Float, done func(error)) {
		done(expectedErr)
	}, 100.0, w)

	m.OnError = func(err error) {
		errChan <- err
	}

	m.Start()

	select {
	case err := <-errChan:
		assert.Equal(t, expectedErr, err)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnError callback")
	}
}

func TestProgressModal_PreventsDoubleStart(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	executionCount := 0
	done := make(chan struct{})

	m := modal.NewProgressInfinite("Title", "Message", func(cb func(error)) {
		executionCount++
		cb(nil)
	}, w)

	m.OnSuccess = func() {
		close(done)
	}

	m.Start()
	m.Start() // Second call should be ignored immediately

	select {
	case <-done:
		// Give a tiny buffer to confirm second execution never happens
		time.Sleep(50 * time.Millisecond)
		assert.Equal(t, 1, executionCount)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for action completion")
	}
}

func TestProgressInfiniteModal_Success(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	successCalled := make(chan struct{})

	m := modal.NewProgressInfinite("Title", "Message", func(done func(error)) {
		done(nil)
	}, w)

	m.OnSuccess = func() {
		close(successCalled)
	}

	m.Start()

	select {
	case <-successCalled:
		// Success case passed
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnSuccess callback")
	}
}

func TestProgressCancelModal_Success(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	successCalled := make(chan struct{})
	cancelRegistered := make(chan struct{})

	m := modal.NewProgressWithCancel("Title", "Message", func(pg binding.Float, onCancel func(func()), done func(error)) {
		onCancel(func() {})
		close(cancelRegistered)
		done(nil)
	}, 100.0, w)

	m.OnSuccess = func() {
		close(successCalled)
	}

	m.Start()

	select {
	case <-successCalled:
		select {
		case <-cancelRegistered:
			// onCancel was called, as expected
		default:
			t.Error("onCancel should have been called by the action")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnSuccess callback")
	}
}

func TestProgressCancelModal_ErrorHandling(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	expectedErr := errors.New("something went wrong")
	errChan := make(chan error, 1)

	m := modal.NewProgressWithCancel("Title", "Message", func(pg binding.Float, onCancel func(func()), done func(error)) {
		done(expectedErr)
	}, 100.0, w)

	m.OnError = func(err error) {
		errChan <- err
	}

	m.Start()

	select {
	case err := <-errChan:
		assert.Equal(t, expectedErr, err)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnError callback")
	}
}

func TestProgressCancelModal_PreventsDoubleStart(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	executionCount := 0
	done := make(chan struct{})

	m := modal.NewProgressWithCancel("Title", "Message", func(pg binding.Float, onCancel func(func()), cb func(error)) {
		executionCount++
		cb(nil)
	}, 100.0, w)

	m.OnSuccess = func() {
		close(done)
	}

	m.Start()
	m.Start() // Second call should be ignored immediately

	select {
	case <-done:
		time.Sleep(50 * time.Millisecond)
		assert.Equal(t, 1, executionCount)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for action completion")
	}
}

func TestProgressInfiniteCancelModal_Success(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	successCalled := make(chan struct{})

	m := modal.NewProgressInfiniteWithCancel("Title", "Message", func(onCancel func(func()), done func(error)) {
		done(nil)
	}, w)

	m.OnSuccess = func() {
		close(successCalled)
	}

	m.Start()

	select {
	case <-successCalled:
		// Success case passed
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnSuccess callback")
	}
}

func TestProgressInfiniteCancelModal_ErrorHandling(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()
	w := testApp.NewWindow("Test")

	expectedErr := errors.New("something went wrong")
	errChan := make(chan error, 1)

	m := modal.NewProgressInfiniteWithCancel("Title", "Message", func(onCancel func(func()), done func(error)) {
		done(expectedErr)
	}, w)

	m.OnError = func(err error) {
		errChan <- err
	}

	m.Start()

	select {
	case err := <-errChan:
		assert.Equal(t, expectedErr, err)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for OnError callback")
	}
}
