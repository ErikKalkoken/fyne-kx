package widget

import (
	"testing"

	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

func TestTappableIcon_Disable_RecolorsSVGResource(t *testing.T) {
	test.NewApp()
	icon := NewTappableIcon(theme.HomeIcon(), nil)

	icon.Disable()

	assert.NotEqual(t, icon.resource, icon.Icon.Resource, "a disabled SVG resource must be recolored, not reused as-is")

	icon.Enable()
	assert.Equal(t, icon.resource, icon.Icon.Resource)
}

func TestTappableIcon_Disable_KeepsNonSVGResourceUnchanged(t *testing.T) {
	// iconQuestionmark32Png is a raster (PNG) resource. theme.NewDisabledResource
	// only knows how to recolor SVG content, so a raster resource must be left
	// as-is rather than passed through it (which would break the image).
	test.NewApp()
	icon := NewTappableIcon(iconQuestionmark32Png, nil)

	icon.Disable()

	assert.Equal(t, icon.resource, icon.Icon.Resource)
}

func TestTappableIcon_Disabled_CursorNeverShowsPointer(t *testing.T) {
	test.NewApp()
	icon := NewTappableIcon(theme.HomeIcon(), nil)
	icon.hovered = true
	icon.Disable()

	assert.Equal(t, desktop.DefaultCursor, icon.Cursor())
}

func TestTappableIcon_Disabled_MouseInDoesNotSetHovered(t *testing.T) {
	test.NewApp()
	icon := NewTappableIcon(theme.HomeIcon(), nil)
	icon.Disable()

	icon.MouseIn(&desktop.MouseEvent{})

	assert.False(t, icon.hovered)
}
