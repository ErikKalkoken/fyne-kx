// Demo is a Fyne app for demonstrating the features provided by the fyne-kx library.
package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	kxdialog "github.com/ErikKalkoken/fyne-kx/dialog"
	kxmodal "github.com/ErikKalkoken/fyne-kx/modal"
	kxtheme "github.com/ErikKalkoken/fyne-kx/theme"
)

type treeItem struct {
	name    string
	content fyne.CanvasObject
}

func main() {
	app := app.New()
	w := app.NewWindow("KX Demo")

	pages := []treeItem{
		{"Badge", makeBadge()},
		{"Columns", makeColumns()},
		{"Dialogs", makeDialogs(w)},
		{"FilterChip", makeFilterChip()},
		{"FilterChipGroup", makeFilterChipGroup()},
		{"FilterChipSelect", makeFilterChipSelect(w)},
		{"IconButton", makeIconButton()},
		{"Modals", makeModals(w)},
		{"RowWrap", makeRowWrap()},
		{"Slider", makeSlider()},
		{"Switch", makeSwitch()},
		{"TappableIcon", makeTappableIcon()},
		{"TappableImage", makeTappableImage()},
		{"TappableLabel", makeTappableLabel()},
		{"ToolbarActionMenu", makeToolbarActionMenu()},
	}
	body := container.NewStack()
	pageIndexes := make(map[string]int)
	for i, it := range pages {
		title := widget.NewLabel(it.name)
		title.TextStyle.Bold = true
		p := container.NewBorder(
			container.NewVBox(title, widget.NewSeparator()),
			nil,
			nil,
			nil,
			it.content,
		)
		p.Hide()
		body.Add(p)
		pageIndexes[it.name] = i
	}
	currentPageIdx := -1
	nav := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				s := []widget.TreeNodeID{
					"Dialogs",
					"Layouts",
					"Modals",
					"Widgets",
				}
				return s
			case "Layouts":
				s := []widget.TreeNodeID{
					"Columns",
					"RowWrap",
				}
				return s
			case "Widgets":
				s := []widget.TreeNodeID{
					"Badge",
					"FilterChip",
					"FilterChipGroup",
					"FilterChipSelect",
					"IconButton",
					"Slider",
					"Switch",
					"TappableIcon",
					"TappableImage",
					"TappableLabel",
					"ToolbarActionMenu",
				}
				return s
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "Layouts" || id == "Widgets"
		},
		func(b bool) fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TreeNodeID, b bool, co fyne.CanvasObject) {
			text := id
			co.(*widget.Label).SetText(text)
		},
	)
	nav.OnSelected = func(id widget.TreeNodeID) {
		if nav.IsBranch(id) {
			nav.UnselectAll()
			return
		}
		if currentPageIdx >= 0 {
			body.Objects[currentPageIdx].Hide()
		}
		idx, found := pageIndexes[id]
		if !found {
			log.Fatalf("content not defined for ID %s", id)
		}
		body.Objects[idx].Show()
		currentPageIdx = idx
	}

	theme := widget.NewSelect([]string{"Auto", "Light", "Dark"}, func(s string) {
		switch s {
		case "Light":
			app.Settings().SetTheme(kxtheme.DefaultWithFixedVariant(theme.VariantLight))
		case "Dark":
			app.Settings().SetTheme(kxtheme.DefaultWithFixedVariant(theme.VariantDark))
		default:
			app.Settings().SetTheme(theme.DefaultTheme())
		}

	})
	theme.SetSelected("Auto")
	bottom := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Theme"),
			theme,
		),
	)

	main := container.NewHSplit(nav, body)
	main.SetOffset(0.33)
	w.SetContent(container.NewBorder(
		nil,
		bottom,
		nil,
		nil,
		main,
	))
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}

func makeDialogs(w fyne.Window) fyne.CanvasObject {
	c := container.NewVBox(
		widget.NewButton("Information Dialog with key handler", func() {
			d := dialog.NewInformation("Info", "You can close this dialog with the Escape key.", w)
			kxdialog.AddDialogKeyHandler(d, w)
			d.Show()
		}),
		widget.NewButton("Confirm Dialog with key handler", func() {
			d := dialog.NewConfirm("Confirm", "You can close this dialog with the Escape key.", func(b bool) {
				fmt.Printf("Confirm dialog: %v\n", b)
			}, w)
			kxdialog.AddDialogKeyHandler(d, w)
			d.Show()
		}),
	)
	return c
}

func makeModals(w fyne.Window) *fyne.Container {
	b1 := widget.NewButton("ProgressModal", func() {
		m := kxmodal.NewProgress(
			"ProgressModal",
			"Please wait...",
			func(progress binding.Float) error {
				for i := float64(1); i < 50; i++ {
					fyne.Do(func() {
						progress.Set(float64(i))
					})
					time.Sleep(100 * time.Millisecond)
				}
				return nil
			}, 50, w)
		m.Start()
	})

	b2 := widget.NewButton("ProgressCancelModal", func() {
		m := kxmodal.NewProgressWithCancel("ProgressCancelModal", "Please wait...", func(progress binding.Float, canceled chan struct{}) error {
			ticker := time.NewTicker(100 * time.Millisecond)
			for i := 1; i < 50; i++ {
				fyne.Do(func() {
					progress.Set(float64(i))
				})
				select {
				case <-canceled:
					return nil
				case <-ticker.C:
				}
			}
			return nil
		}, 50, w)
		m.Start()
	})

	b3 := widget.NewButton("ProgressInfiniteModal", func() {
		m := kxmodal.NewProgressInfinite("ProgressInfiniteModal", "Please wait...", func() error {
			time.Sleep(3 * time.Second)
			return nil
		}, w)
		m.Start()
	})

	b4 := widget.NewButton("ProgressInfiniteCancelModal", func() {
		m := kxmodal.NewProgressInfiniteWithCancel("ProgressInfiniteCancelModal", "Please wait...", func(canceled chan struct{}) error {
			ticker := time.NewTicker(100 * time.Millisecond)
			for i := 1; i < 50; i++ {
				select {
				case <-canceled:
					return nil
				case <-ticker.C:
				}
			}
			return nil
		}, w)
		m.Start()
	})

	return container.NewVBox(b1, b2, b3, b4)
}
