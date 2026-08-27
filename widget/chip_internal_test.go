package widget

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"

	"github.com/stretchr/testify/assert"
)

func TestChip_CanCreateEnabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "chip/enabled_off.png", w.Canvas().Capture())
}

func TestChip_CanCreateDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "chip/disabled_off.png", w.Canvas().Capture())
}

func TestChip_CanCreateEnabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	c.on = true
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "chip/enabled_on.png", w.Canvas().Capture())
}

func TestChip_CanCreateDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	c.on = true
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "chip/disabled_on.png", w.Canvas().Capture())
}

func TestChip_CanTapWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	c := newChip("Test", func() {
		tapped = true
	})
	w := test.NewWindow(c)
	defer w.Close()

	test.Tap(c)
	assert.True(t, tapped)
}

func TestChip_CanNotTapWhenDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	c := newChip("Test", func() {
		tapped = true
	})
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.Tap(c)
	assert.False(t, tapped)
}

func TestChip_WithLeadingAndTrailingIcons(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	c.leadingIcon = theme.InfoIcon()
	c.trailingIcon = theme.CancelIcon()
	c.Refresh()

	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "chip/with_icons.png", w.Canvas().Capture())
}

func TestChip_DynamicallyAddIcons(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	w := test.NewWindow(c)
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	c.leadingIcon = theme.ConfirmIcon()
	c.trailingIcon = theme.DeleteIcon()
	c.Refresh()

	test.AssertImageMatches(t, "chip/icons_added.png", w.Canvas().Capture())
}

func TestChip_DynamicallyRemoveIcons(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	c.leadingIcon = theme.ConfirmIcon()
	c.trailingIcon = theme.DeleteIcon()
	w := test.NewWindow(c)
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	c.leadingIcon = nil
	c.trailingIcon = nil
	c.Refresh()

	test.AssertImageMatches(t, "chip/icons_removed.png", w.Canvas().Capture())
}

func TestChip_Hover(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	w := test.NewWindow(c)
	defer w.Close()

	// Initial cursor when not hovered
	assert.Equal(t, desktop.DefaultCursor, c.Cursor())

	// Mouse inside enabled chip updates cursor
	c.MouseIn(&desktop.MouseEvent{})
	assert.Equal(t, desktop.PointerCursor, c.Cursor())

	// Mouse out resets cursor state
	c.MouseOut()
	assert.Equal(t, desktop.DefaultCursor, c.Cursor())

	// Mouse inside disabled chip stays default cursor
	c.Disable()
	c.MouseIn(&desktop.MouseEvent{})
	assert.Equal(t, desktop.DefaultCursor, c.Cursor())
}

func TestChip_Focus(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := newChip("Test", nil)
	w := test.NewWindow(c)
	defer w.Close()

	// Focus gained applies focus styling
	c.FocusGained()
	test.AssertImageMatches(t, "chip/focused.png", w.Canvas().Capture())

	// Focus lost restores default appearance
	c.FocusLost()
	test.AssertImageMatches(t, "chip/enabled_off.png", w.Canvas().Capture())

	// Disabled chip ignores focus gained
	c.Disable()
	c.FocusGained()
	test.AssertImageMatches(t, "chip/disabled_off.png", w.Canvas().Capture())
}

func TestChip_TypedRuneWhenFocused(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	var tapped bool
	c := newChip("Test", func() {
		tapped = true
	})
	w := test.NewWindow(c)
	defer w.Close()

	// Space key triggers tap action when enabled
	c.TypedRune(' ')
	assert.True(t, tapped)

	// Space key is ignored when disabled
	tapped = false
	c.Disable()
	c.TypedRune(' ')
	assert.False(t, tapped)
}
