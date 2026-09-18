package widget

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestBadge_TextColorMatchesImportance(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	badge := NewBadge("Test")
	w := test.NewWindow(badge)
	defer w.Close()

	cases := []struct {
		importance widget.Importance
		colorName  fyne.ThemeColorName
	}{
		{widget.MediumImportance, theme.ColorNameForeground},
		{widget.LowImportance, theme.ColorNameForeground},
		{widget.HighImportance, theme.ColorNameForegroundOnPrimary},
		{widget.DangerImportance, theme.ColorNameForegroundOnError},
		{widget.WarningImportance, theme.ColorNameForegroundOnWarning},
		{widget.SuccessImportance, theme.ColorNameForegroundOnSuccess},
	}
	for _, c := range cases {
		badge.Importance = c.importance
		badge.Refresh()
		r := test.WidgetRenderer(badge).(*badgeRenderer)
		assert.Equal(t, c.colorName, r.segment.Style.ColorName, "importance %v", c.importance)
	}
}
