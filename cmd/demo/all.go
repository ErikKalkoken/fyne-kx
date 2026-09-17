package main

import (
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	group := kxwidget.NewFilterChipGroup([]string{"Alpha", "Bravo"}, func(s []string) {
		log.Printf("FilterChipGroup: %v\n", s)
	})
	group.Selected = []string{"Bravo"}
	sw := kxwidget.NewSwitch(func(on bool) {
		log.Printf("Switch: %v\n", on)
	})
	sw.On = true
	items := []fyne.CanvasObject{
		makeRow("Badge", badge),
		makeRow("FilterChip", kxwidget.NewFilterChip("Filter", func(on bool) {
			log.Printf("FilterChip: %v\n", on)
		})),
		makeRow("FilterChipGroup", group),
		makeRow("FilterChipSelect", kxwidget.NewFilterChipSelect("Select", []string{"Alpha", "Bravo", "Charlie"}, func(s string) {
			log.Printf("FilterChipSelect: %s\n", s)
		})),
		makeRow("IconButton", kxwidget.NewIconButton(theme.AccountIcon(), func() {
			log.Println("IconButton tapped")
		})),
		makeRow("LoadingButton", kxwidget.NewLoadingButton("Run action", nil, func(done func()) {
			log.Println("LoadingButton: action started")
			go func() {
				defer done()
				time.Sleep(simulatedWorkDuration)
				log.Println("LoadingButton: action completed")
			}()
		})),
		makeRow("Slider", func() fyne.CanvasObject {
			s := kxwidget.NewSlider(0, 100)
			s.SetValue(25)
			s.OnChangeEnded = func(v float64) {
				log.Printf("Slider: %v\n", v)
			}
			return s
		}()),
		makeRow("SortChip", kxwidget.NewSortChip([]string{"Name", "Age"}, "Name", kxwidget.SortOrderAscending, func(c string, o kxwidget.SortOrder) {
			log.Printf("SortChip: %s %v\n", c, o)
		})),
		makeRow("Spinner", func() fyne.CanvasObject {
			r := kxwidget.NewSpinner()
			r.Start()
			return r
		}()),
		makeRow("Switch", sw),
		makeRow("TappableIcon", kxwidget.NewTappableIcon(theme.AccountIcon(), func() {
			log.Println("TappableIcon tapped")
		})),
		makeRow("TappableImage", func() fyne.CanvasObject {
			img := kxwidget.NewTappableImage(resourceIconPng, func() {
				log.Println("TappableImage tapped")
			})
			img.SetMinSize(fyne.NewSize(48, 48))
			return img
		}()),
		makeRow("TappableLabel", kxwidget.NewTappableLabel("Tap me", func() {
			log.Println("TappableLabel tapped")
		})),
		makeRow("ToolbarActionMenu", widget.NewToolbar(kxwidget.NewToolbarActionMenu(theme.MenuIcon(), fyne.NewMenu("",
			fyne.NewMenuItem("First", func() { log.Println("ToolbarActionMenu: First selected") }),
			fyne.NewMenuItem("Second", func() { log.Println("ToolbarActionMenu: Second selected") }),
		)))),
	}

	makeColumn := func(items []fyne.CanvasObject) fyne.CanvasObject {
		col := container.NewVBox()
		for i, it := range items {
			if i > 0 {
				col.Add(makeSpace())
			}
			col.Add(it)
		}
		return col
	}

	mid := (len(items) + 1) / 2
	halfGap := theme.Padding() * 4
	left := container.New(layout.NewCustomPaddedLayout(0, 0, 0, halfGap), makeColumn(items[:mid]))
	right := container.New(layout.NewCustomPaddedLayout(0, 0, halfGap, 0), makeColumn(items[mid:]))
	columns := container.NewGridWithColumns(2, left, right)

	margin := theme.Size(theme.SizeNameScrollBar)
	padded := container.New(layout.NewCustomPaddedLayout(0, 0, 0, margin), columns)
	return container.NewVScroll(padded)
}
