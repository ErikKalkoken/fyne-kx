package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	fynewidget "fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func TestNewToolbarActionMenu_CanCreate(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	menu := fyne.NewMenu("", fyne.NewMenuItem("item", nil))
	action := kxwidget.NewToolbarActionMenu(theme.HomeIcon(), menu)

	assert.NotNil(t, action)
	assert.NotNil(t, action.OnActivated)
}

func TestNewToolbarActionMenu_ActivateShowsMenu(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	menu := fyne.NewMenu("", fyne.NewMenuItem("item", nil))
	action := kxwidget.NewToolbarActionMenu(theme.HomeIcon(), menu)
	tb := fynewidget.NewToolbar(action)
	w := test.NewWindow(tb)
	defer w.Close()
	w.Resize(fyne.NewSize(200, 200))

	assert.NotPanics(t, func() { action.OnActivated() })
}
