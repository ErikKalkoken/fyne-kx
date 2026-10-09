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

func TestIconButton_SetMenuItemsAddsMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	var tapped bool
	icon := kxwidget.NewIconButton(theme.HomeIcon(), func() {
		tapped = true
	})
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()

	icon.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("new", nil)})
	test.Tap(icon)

	assert.NotNil(t, w.Canvas().Overlays().Top(), "should show menu")
	assert.False(t, tapped, "should have cleared OnTapped")
}

func TestIconButton_OnTappedTakesPrecedenceOverMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()
	var tapped bool
	icon.OnTapped = func() {
		tapped = true
	}

	test.Tap(icon)

	assert.True(t, tapped)
	assert.Nil(t, w.Canvas().Overlays().Top(), "should not show menu")
}

func TestIconButton_ShowMenuAgainWhenOnTappedCleared(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), fyne.NewMenu("", fyne.NewMenuItem("item", nil)))
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()
	icon.OnTapped = func() {}
	icon.OnTapped = nil

	test.Tap(icon)

	assert.NotNil(t, w.Canvas().Overlays().Top(), "should show menu")
}

func TestIconButton_ShowNoMenuWhenMenuEmpty(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	icon := kxwidget.NewIconButtonWithMenu(theme.HomeIcon(), fyne.NewMenu(""))
	w := test.NewWindow(container.NewCenter(icon))
	defer w.Close()

	test.Tap(icon)

	assert.Nil(t, w.Canvas().Overlays().Top())
}

type embeddedIconButton struct {
	kxwidget.IconButton
}

func newEmbeddedIconButton() *embeddedIconButton {
	w := &embeddedIconButton{}
	w.ExtendBaseWidget(w)
	w.SetIcon(theme.HomeIcon())
	return w
}

func TestIconButton_Embedded(t *testing.T) {
	t.Run("can tap", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		var tapped bool
		icon := newEmbeddedIconButton()
		icon.OnTapped = func() {
			tapped = true
		}
		w := test.NewWindow(icon)
		defer w.Close()

		test.Tap(icon)

		assert.True(t, tapped)
		test.AssertImageMatches(t, "iconbutton/normal.png", w.Canvas().Capture())
	})
	t.Run("can show menu", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		icon := newEmbeddedIconButton()
		icon.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("item", nil)})
		w := test.NewWindow(container.NewCenter(icon))
		defer w.Close()
		w.Resize(fyne.NewSize(100, 150))

		test.Tap(icon)

		test.AssertImageMatches(t, "iconbutton/menu_enabled.png", w.Canvas().Capture())
	})
	t.Run("shows no menu when disabled", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		icon := newEmbeddedIconButton()
		icon.SetMenuItems([]*fyne.MenuItem{fyne.NewMenuItem("item", nil)})
		icon.Disable()
		w := test.NewWindow(container.NewCenter(icon))
		defer w.Close()
		w.Resize(fyne.NewSize(100, 150))

		test.Tap(icon)

		test.AssertImageMatches(t, "iconbutton/menu_disabled.png", w.Canvas().Capture())
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

func TestIconButton_Hover(t *testing.T) {
	t.Run("shows background when hovered", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		icon := kxwidget.NewIconButton(theme.HomeIcon(), func() {})
		w := test.NewWindow(icon)
		defer w.Close()

		icon.MouseIn(&desktop.MouseEvent{})
		test.AssertImageMatches(t, "iconbutton/hovered.png", w.Canvas().Capture())

		icon.MouseOut()
		test.AssertImageMatches(t, "iconbutton/normal.png", w.Canvas().Capture())
	})
	t.Run("shows background when embedded", func(t *testing.T) {
		test.NewTempApp(t)
		test.ApplyTheme(t, test.Theme())
		icon := newEmbeddedIconButton()
		icon.OnTapped = func() {}
		w := test.NewWindow(icon)
		defer w.Close()

		icon.MouseIn(&desktop.MouseEvent{})
		test.AssertImageMatches(t, "iconbutton/hovered.png", w.Canvas().Capture())
	})
}
