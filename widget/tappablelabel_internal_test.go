package widget

import (
	"testing"

	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestTappableLabel_DisableEnableTogglesImportance(t *testing.T) {
	test.NewApp()
	label := NewTappableLabel("Test", nil)

	label.Disable()
	assert.Equal(t, widget.LowImportance, label.Importance)

	label.Enable()
	assert.Equal(t, widget.MediumImportance, label.Importance)
}

func TestTappableLabel_Disabled_CursorNeverShowsPointer(t *testing.T) {
	test.NewApp()
	label := NewTappableLabel("Test", nil)
	label.hovered = true
	label.Disable()

	assert.Equal(t, desktop.DefaultCursor, label.Cursor())
}

func TestTappableLabel_Disabled_MouseInDoesNotSetHovered(t *testing.T) {
	test.NewApp()
	label := NewTappableLabel("Test", nil)
	label.Disable()

	label.MouseIn(&desktop.MouseEvent{})

	assert.False(t, label.hovered)
}
