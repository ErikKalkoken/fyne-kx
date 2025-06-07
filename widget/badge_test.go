package widget_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
	"github.com/stretchr/testify/assert"
)

func TestBadge_CanCreateWithDefaults(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	badge := kxwidget.NewBadge("Test")
	w := test.NewWindow(badge)
	defer w.Close()

	assert.Equal(t, "Test", badge.Text)
	test.AssertImageMatches(t, "badge/default.png", w.Canvas().Capture())
}

func TestBadge_CanSetText(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	badge := kxwidget.NewBadge("Alpha")
	w := test.NewWindow(badge)
	defer w.Close()

	badge.SetText("Bravo")

	assert.Equal(t, "Bravo", badge.Text)
	test.AssertImageMatches(t, "badge/set_text.png", w.Canvas().Capture())
}

func TestBadge_CanUpdateText(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	badge := kxwidget.NewBadge("Alpha")
	w := test.NewWindow(badge)
	defer w.Close()

	badge.Text = "Bravo"
	badge.Refresh()

	assert.Equal(t, "Bravo", badge.Text)
	test.AssertImageMatches(t, "badge/set_text.png", w.Canvas().Capture())
}

func TestBadge_CanUpdateImportance(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())
	badge := kxwidget.NewBadge("Test")
	w := test.NewWindow(badge)
	defer w.Close()

	cases := []struct {
		importance widget.Importance
		filename   string
	}{
		{widget.DangerImportance, "danger"},
		{widget.HighImportance, "primary"},
		{widget.LowImportance, "disabled"},
		{widget.MediumImportance, "default"},
		{widget.SuccessImportance, "success"},
		{widget.WarningImportance, "warning"},
	}
	for _, tc := range cases {
		badge.Importance = tc.importance
		badge.Refresh()
		test.AssertImageMatches(t, "badge/"+tc.filename+".png", w.Canvas().Capture())
	}
}
