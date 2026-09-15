package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

// makeAll returns a scrollable page showing a compact example of every
// widget provided by this library, so users can see them all at a glance.
func makeAll() fyne.CanvasObject {
	makeRow := func(name string, example fyne.CanvasObject) fyne.CanvasObject {
		label := widget.NewLabel(name)
		label.TextStyle.Bold = true
		label.Alignment = fyne.TextAlignTrailing
		return container.NewBorder(nil, nil, container.NewGridWrap(fyne.NewSize(150, label.MinSize().Height), label), nil, example)
	}
	makeSpace := func() fyne.CanvasObject {
		r := canvas.NewRectangle(color.Transparent)
		r.SetMinSize(fyne.NewSize(1, theme.Padding()))
		return r
	}

	badge := kxwidget.NewBadge("Example")
	badge.Importance = widget.WarningImportance
	group := kxwidget.NewFilterChipGroup([]string{"Alpha", "Bravo", "Charlie"}, func(s []string) {})
	group.Selected = []string{"Bravo"}
	items := []fyne.CanvasObject{
		makeRow("Badge", badge),
		makeRow("FilterChip", kxwidget.NewFilterChip("Filter", func(on bool) {})),
		makeRow("FilterChipGroup", group),
		makeRow("FilterChipSelect", kxwidget.NewFilterChipSelect("Select", []string{"Alpha", "Bravo", "Charlie"}, func(s string) {})),
		makeRow("IconButton", kxwidget.NewIconButton(theme.AccountIcon(), func() {})),
		makeRow("ProgressButton", kxwidget.NewProgressButton("Run action", nil, func(done func()) {
			go func() {
				defer done()
				time.Sleep(simulatedWorkDuration)
			}()
		})),
		makeRow("Slider", func() fyne.CanvasObject {
			s := kxwidget.NewSlider(0, 100)
			s.SetValue(25)
			return s
		}()),
		makeRow("SortChip", kxwidget.NewSortChip([]string{"Name", "Age"}, "Name", kxwidget.SortOrderAscending, func(c string, o kxwidget.SortOrder) {})),
		makeRow("Switch", kxwidget.NewSwitch(func(on bool) {})),
		makeRow("TappableIcon", kxwidget.NewTappableIcon(theme.AccountIcon(), func() {})),
		makeRow("TappableImage", func() fyne.CanvasObject {
			img := kxwidget.NewTappableImage(resourceIconPng, func() {})
			img.SetMinSize(fyne.NewSize(48, 48))
			return img
		}()),
		makeRow("TappableLabel", kxwidget.NewTappableLabel("Tap me", func() {})),
		makeRow("ToolbarActionMenu", widget.NewToolbar(kxwidget.NewToolbarActionMenu(theme.MenuIcon(), fyne.NewMenu("",
			fyne.NewMenuItem("First", func() {}),
			fyne.NewMenuItem("Second", func() {}),
		)))),
	}

	rows := container.NewVBox()
	for i, it := range items {
		if i > 0 {
			rows.Add(makeSpace())
		}
		rows.Add(it)
	}

	return container.NewVScroll(rows)
}
