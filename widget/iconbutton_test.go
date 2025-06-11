package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestIconButton_Create(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	icon := kxwidget.NewIconButton(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	test.AssertImageMatches(t, "iconbutton/normal.png", w.Canvas().Capture())
}

func TestIconButton_SetIcon(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButton(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	icon.SetIcon(theme.ComputerIcon())

	test.AssertImageMatches(t, "iconbutton/set_icon.png", w.Canvas().Capture())
}

func TestIconButton_TappableWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	icon := kxwidget.NewIconButton(theme.HomeIcon(), func() {
		tapped = true
	})
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
	assert.True(t, tapped)
}

func TestIconButton_IgnoreTapWhenNoCallback(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButton(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
}

func TestIconButton_NotTappableWhenDisabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	icon := kxwidget.NewIconButton(theme.HomeIcon(), func() {
		tapped = true
	})
	icon.Disable()
	w := test.NewWindow(icon)
	defer w.Close()

	test.Tap(icon)
	assert.False(t, tapped, "should not be tappable")
	test.AssertImageMatches(t, "iconbutton/disabled.png", w.Canvas().Capture())
}

func TestIconButton_CreateWithMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()
	w.Resize(fyne.NewSize(100, 150))

	test.AssertImageMatches(t, "iconbutton/create_menu.png", w.Canvas().Capture())
}

func TestIconButton_ShowMenuWhenTappedAndEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()
	w.Resize(fyne.NewSize(100, 150))

	test.Tap(icon)

	test.AssertImageMatches(t, "iconbutton/menu_enabled.png", w.Canvas().Capture())
}

func TestIconButton_ShowNoMenuWhenTappedAndDisabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()
	w.Resize(fyne.NewSize(100, 150))
	icon.Disable()

	test.Tap(icon)

	test.AssertImageMatches(t, "iconbutton/menu_disabled.png", w.Canvas().Capture())
}
