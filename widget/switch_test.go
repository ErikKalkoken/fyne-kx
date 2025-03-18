package widget_test

import (
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestSwitch(t *testing.T) {
	t.Run("can create initial switch", func(t *testing.T) {
		w := widget.NewSwitch(nil)
		window := test.NewWindow(w)
		defer window.Close()
		assert.False(t, w.On)
	})
	t.Run("should set state and ignore undefined callback", func(t *testing.T) {
		w := widget.NewSwitch(nil)
		window := test.NewWindow(w)
		defer window.Close()
		w.SetOn(true)
		assert.True(t, w.On)
	})
	t.Run("should run callback when state changed", func(t *testing.T) {
		var hasRun bool
		w := widget.NewSwitch(func(on bool) {
			hasRun = true
		})
		window := test.NewWindow(w)
		defer window.Close()
		w.SetOn(true)
		assert.True(t, hasRun)
	})
	t.Run("should not run callback when state did not change", func(t *testing.T) {
		var hasRun bool
		w := widget.NewSwitch(func(on bool) {
			hasRun = true
		})
		w.On = true
		window := test.NewWindow(w)
		defer window.Close()
		w.SetOn(true)
		assert.False(t, hasRun)
	})
}
