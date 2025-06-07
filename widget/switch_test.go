package widget_test

import (
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestSwitch_CanCreateEnabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	sw := widget.NewSwitch(nil)
	w := test.NewWindow(sw)
	defer w.Close()

	test.AssertImageMatches(t, "switch/enabled_off.png", w.Canvas().Capture())
}

func TestSwitch_CanCreateEnabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	sw := widget.NewSwitch(nil)
	sw.On = true
	w := test.NewWindow(sw)
	defer w.Close()

	test.AssertImageMatches(t, "switch/enabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanCreateDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	sw := widget.NewSwitch(nil)
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	test.AssertImageMatches(t, "switch/disabled_off.png", w.Canvas().Capture())
}

func TestSwitch_CanCreateDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	sw := widget.NewSwitch(nil)
	sw.On = true
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	test.AssertImageMatches(t, "switch/disabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanSwitchOnWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	w := test.NewWindow(sw)
	defer w.Close()

	test.Tap(sw)

	assert.True(t, tapped)
	assert.True(t, sw.On)
	test.AssertImageMatches(t, "switch/enabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanSwitchOffWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	sw.On = true
	w := test.NewWindow(sw)
	defer w.Close()

	test.Tap(sw)

	assert.True(t, tapped)
	assert.False(t, sw.On)
	test.AssertImageMatches(t, "switch/enabled_off.png", w.Canvas().Capture())
}

func TestSwitch_CanNotSwitchOnWhenDisabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	test.Tap(sw)

	assert.False(t, tapped)
	assert.False(t, sw.On)
	test.AssertImageMatches(t, "switch/disabled_off.png", w.Canvas().Capture())
}

func TestSwitch_CanNotSwitchOffWhenDisabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	sw.On = true
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	test.Tap(sw)

	assert.False(t, tapped)
	assert.True(t, sw.On)
	test.AssertImageMatches(t, "switch/disabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanIgnoreCallbackWhenNotDefined(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(nil)
	w := test.NewWindow(sw)
	defer w.Close()

	test.Tap(sw)

	assert.False(t, tapped)
	assert.True(t, sw.On)
}

func TestSwitch_CanSetOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	w := test.NewWindow(sw)
	defer w.Close()

	sw.SetOn(true)

	assert.True(t, tapped)
	assert.True(t, sw.On)
	test.AssertImageMatches(t, "switch/enabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanSetOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	sw.On = true
	w := test.NewWindow(sw)
	defer w.Close()

	sw.SetOn(false)

	assert.True(t, tapped)
	assert.False(t, sw.On)
	test.AssertImageMatches(t, "switch/enabled_off.png", w.Canvas().Capture())
}

func TestSwitch_SetState_DoNothingWhenSameState_1(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	w := test.NewWindow(sw)
	defer w.Close()

	sw.SetOn(false)

	assert.False(t, tapped)
	assert.False(t, sw.On)
	test.AssertImageMatches(t, "switch/enabled_off.png", w.Canvas().Capture())
}

func TestSwitch_SetState_DoNothingWhenSameState_2(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	sw := widget.NewSwitch(func(on bool) {
		tapped = true
	})
	sw.On = true
	w := test.NewWindow(sw)
	defer w.Close()

	sw.SetOn(true)

	assert.False(t, tapped)
	assert.True(t, sw.On)
	test.AssertImageMatches(t, "switch/enabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanFocusWhenEnabledAndOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	w := test.NewWindow(sw)
	defer w.Close()

	w.Canvas().Focus(sw)

	test.AssertImageMatches(t, "switch/focused_off.png", w.Canvas().Capture())
}

func TestSwitch_CanFocusWhenEnabledAndOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	sw.On = true
	w := test.NewWindow(sw)
	defer w.Close()

	w.Canvas().Focus(sw)

	test.AssertImageMatches(t, "switch/focused_on.png", w.Canvas().Capture())
}

func TestSwitch_CanNotFocusWhenDisabledAndOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	w.Canvas().Focus(sw)

	test.AssertImageMatches(t, "switch/disabled_off.png", w.Canvas().Capture())
}

func TestSwitch_CanNotFocusWhenDisabledAndOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	sw.On = true
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	w.Canvas().Focus(sw)

	test.AssertImageMatches(t, "switch/disabled_on.png", w.Canvas().Capture())
}

func TestSwitch_CanSwitchWithKeyWhenOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	w := test.NewWindow(sw)
	defer w.Close()

	test.Type(sw, " ")

	test.AssertImageMatches(t, "switch/focused_on.png", w.Canvas().Capture())
}

func TestSwitch_CanSwitchWithKeyWhenOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	sw.On = true
	w := test.NewWindow(sw)
	defer w.Close()

	test.Type(sw, " ")

	test.AssertImageMatches(t, "switch/focused_off.png", w.Canvas().Capture())
}

func TestSwitch_CanNotSwitchWithKeyWhenDisabledAndOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	test.Type(sw, " ")

	test.AssertImageMatches(t, "switch/disabled_off.png", w.Canvas().Capture())
}

func TestSwitch_CanNotSwitchWithKeyWhenDisabledAndOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	sw := widget.NewSwitch(nil)
	sw.On = true
	sw.Disable()
	w := test.NewWindow(sw)
	defer w.Close()

	test.Type(sw, " ")

	test.AssertImageMatches(t, "switch/disabled_on.png", w.Canvas().Capture())
}
