package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestIconButton_CreateNormal(t *testing.T) {
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

func TestIconButton_MouseInOutTogglesPointerCursorWhenEnabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButton(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	assert.Equal(t, desktop.DefaultCursor, icon.Cursor())

	icon.MouseIn(&desktop.MouseEvent{})
	assert.Equal(t, desktop.PointerCursor, icon.Cursor())

	icon.MouseOut()
	assert.Equal(t, desktop.DefaultCursor, icon.Cursor())
}

func TestIconButton_MouseInDoesNotHoverWhenDisabled(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButton(theme.HomeIcon(), nil)
	icon.Disable()
	w := test.NewWindow(icon)
	defer w.Close()

	icon.MouseIn(&desktop.MouseEvent{})

	assert.Equal(t, desktop.DefaultCursor, icon.Cursor(), "a disabled button should never show a hover cursor")
}

func TestIconButton_SetMenuItemsDoesNothingWithoutMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButton(theme.HomeIcon(), nil)
	w := test.NewWindow(icon)
	defer w.Close()

	assert.NotPanics(t, func() {
		icon.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("new", nil)})
	})
}

func TestIconButton_SetMenuItemsReplacesItems(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	menu := fyne.NewMenu("", fyne.NewMenuItem("old", nil))
	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), menu)
	w := test.NewWindow(icon)
	defer w.Close()

	newItems := []*fyne.MenuItem{fyne.NewMenuItem("new", nil)}
	icon.SetMenuItems(newItems)

	assert.Equal(t, newItems, menu.Items)
}
