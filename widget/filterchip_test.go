package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestFilterChip_CanCreateEnabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChip("Test", nil)
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/enabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanCreateDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChip("Test", nil)
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/disabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanCreateEnabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChip("Test", nil)
	c.On = true
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/enabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_CanCreateDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChip("Test", nil)
	c.On = true
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/disabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_CanTurnOnWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	var currentState bool
	var called bool
	c := kxwidget.NewFilterChip("Test", func(on bool) {
		called = true
		currentState = on
	})
	w := test.NewWindow(c)
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	test.Tap(c)
	assert.True(t, called)
	assert.True(t, currentState)
	test.AssertImageMatches(t, "filterchip/tapped_enabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_CanTurnOffWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	var currentState bool
	var called bool
	c := kxwidget.NewFilterChip("Test", func(on bool) {
		called = true
		currentState = on
	})
	c.On = true
	w := test.NewWindow(c)
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	test.Tap(c)
	assert.True(t, called)
	assert.False(t, currentState)
	test.AssertImageMatches(t, "filterchip/tapped_enabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanNotChangeStateWhenDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	c := kxwidget.NewFilterChip("Test", func(on bool) {
		tapped = true
	})
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.Tap(c)
	assert.False(t, tapped)
	test.AssertImageMatches(t, "filterchip/disabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanNotChangeStateWhenDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	c := kxwidget.NewFilterChip("Test", func(on bool) {
		tapped = true
	})
	c.On = true
	c.Disable()
	w := test.NewWindow(c)
	defer w.Close()

	test.Tap(c)
	assert.False(t, tapped)
	test.AssertImageMatches(t, "filterchip/disabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_SetState(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	calls := 0
	c := kxwidget.NewFilterChip("Test", func(on bool) {
		calls++
	})

	c.SetState(true)
	assert.True(t, c.On)
	assert.Equal(t, 1, calls)

	// Re-setting same state should not re-trigger callback
	c.SetState(true)
	assert.Equal(t, 1, calls)
}

func TestFilterChip_SetText(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	c := kxwidget.NewFilterChip("Test", nil)
	w := test.NewWindow(c)
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	c.SetText("Alpha")
	assert.Equal(t, "Alpha", c.Text)
	test.AssertImageMatches(t, "filterchip/set_text.png", w.Canvas().Capture())
}
