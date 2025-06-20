package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestFilterChip_CanCreateEnabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	chip := kxwidget.NewFilterChip("Test", nil)
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/enabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanCreateDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	chip := kxwidget.NewFilterChip("Test", nil)
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()
	chip.Disable()

	test.AssertImageMatches(t, "filterchip/disabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanCreateEnabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	chip := kxwidget.NewFilterChip("Test", nil)
	chip.On = true
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/enabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_CanCreateDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	chip := kxwidget.NewFilterChip("Test", nil)
	chip.On = true
	chip.Disable()
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()

	test.AssertImageMatches(t, "filterchip/disabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_CanSwitchOnWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	chip := kxwidget.NewFilterChip("Test", func(on bool) {
		tapped = true
	})
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	test.Tap(chip)
	assert.True(t, tapped)
	test.AssertImageMatches(t, "filterchip/tapped_enabled_on.png", w.Canvas().Capture())
}

func TestFilterChip_CanSwitchOffWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	chip := kxwidget.NewFilterChip("Test", func(on bool) {
		tapped = true
	})
	chip.On = true
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()
	w.Resize(fyne.NewSize(150, 50))

	test.Tap(chip)
	assert.True(t, tapped)
	test.AssertImageMatches(t, "filterchip/tapped_enabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanNotSwitchWhenDisabledOff(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	chip := kxwidget.NewFilterChip("Test", func(on bool) {
		tapped = true
	})
	chip.Disable()
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()

	test.Tap(chip)
	assert.False(t, tapped)
	test.AssertImageMatches(t, "filterchip/disabled_off.png", w.Canvas().Capture())
}

func TestFilterChip_CanNotSwitchWhenDisabledOn(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	chip := kxwidget.NewFilterChip("Test", func(on bool) {
		tapped = true
	})
	chip.On = true
	chip.Disable()
	w := test.NewWindow(container.NewCenter(chip))
	defer w.Close()

	test.Tap(chip)
	assert.False(t, tapped)
	test.AssertImageMatches(t, "filterchip/disabled_on.png", w.Canvas().Capture())
}
