// App for demonstrating the features provided by the Fyne KX extension.
package main

import (
	"image/color"
	"log"
	"math/rand"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"

	kxlayout "github.com/ErikKalkoken/fyne-kx/layout"
	kxmodal "github.com/ErikKalkoken/fyne-kx/modal"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("KX Demo")
	tabs := container.NewAppTabs(
		container.NewTabItem("Layouts", makeLayouts()),
		container.NewTabItem("Modals", makeModals(w)),
		container.NewTabItem("Widgets", makeWidgets()),
		container.NewTabItem("Colors", makeThemeColors()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.SelectIndex(2)

	w.SetContent(container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		tabs,
	))
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}

func makeLayouts() fyne.CanvasObject {
	layout := kxlayout.NewColumns(150, 100, 50)
	makeBox := func(h float32) fyne.CanvasObject {
		x := canvas.NewRectangle(theme.Color(theme.ColorNameDisabled))
		w := rand.Float32()*100 + 50
		x.SetMinSize(fyne.NewSize(w, h))
		return x
	}
	c := container.NewVBox(
		container.New(layout, makeBox(50), makeBox(50), makeBox(50)),
		container.New(layout, makeBox(150), makeBox(150), makeBox(150)),
		container.New(layout, makeBox(30), makeBox(30), makeBox(30)),
	)
	x := widget.NewLabel("Columns")
	x.TextStyle.Bold = true
	return container.NewBorder(
		container.NewVBox(x, widget.NewSeparator()),
		nil,
		nil,
		nil,
		c,
	)
}

func makeWidgets() fyne.CanvasObject {
	badge := kxwidget.NewBadge("1234")
	img := kxwidget.NewTappableImage(theme.FyneLogo(), func() {
		log.Println("TappableImage")
	})
	img.SetFillMode(canvas.ImageFillContain)
	img.SetMinSize(fyne.NewSize(100, 100))
	icon := kxwidget.NewTappableIcon(theme.AccountIcon(), func() {
		log.Println("TappableIcon")
	})
	label := kxwidget.NewTappableLabel("Tap me", func() {
		log.Println("TappableLabel")
	})
	slider := kxwidget.NewSlider(0, 100)
	slider.SetValue(25)
	toggle1 := kxwidget.NewToggle(func(on bool) {
		log.Println("Toggle 1: ", on)
	})
	toggle1.On = true
	toggle2 := kxwidget.NewToggle(func(on bool) {
		log.Println("Toggle 2: ", on)
	})
	toggle3 := kxwidget.NewToggle(nil)
	toggle3.Disable()
	toggle4 := kxwidget.NewToggle(nil)
	toggle4.SetState(true)
	toggle4.Disable()
	f := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Badge", Widget: badge},
			{Text: "", Widget: container.NewPadded()},
			{Text: "Slider", Widget: slider},
			{Text: "", Widget: container.NewPadded()},
			{Text: "TappableIcon", Widget: container.NewHBox(icon)},
			{Text: "", Widget: container.NewPadded()},
			{Text: "TappableImage", Widget: container.NewHBox(img)},
			{Text: "", Widget: container.NewPadded()},
			{Text: "TappableLabel", Widget: label},
			{Text: "", Widget: container.NewPadded()},
			{Text: "Toggle", Widget: container.NewVBox(toggle1, toggle2, toggle3, toggle4)},
		},
	}
	return f
}

func makeModals(w fyne.Window) *fyne.Container {
	b1 := widget.NewButton("ProgressModal", func() {
		m := kxmodal.NewProgress("ProgressModal", "Please wait...", func(progress binding.Float) error {
			for i := 1; i < 50; i++ {
				progress.Set(float64(i))
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
				progress.Set(float64(i))
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

type colorRow struct {
	label string
	name  fyne.ThemeColorName
}

func makeThemeColors() fyne.CanvasObject {
	colors := []colorRow{
		{"ColorNameBackground", theme.ColorNameBackground},
		{"ColorNameButton", theme.ColorNameButton},
		{"ColorNameDisabled", theme.ColorNameDisabled},
		{"ColorNameDisabledButton", theme.ColorNameDisabledButton},
		{"ColorNameError", theme.ColorNameError},
		{"ColorNameFocus", theme.ColorNameFocus},
		{"ColorNameForeground", theme.ColorNameForeground},
		{"ColorNameForegroundOnError", theme.ColorNameForegroundOnError},
		{"ColorNameForegroundOnPrimary", theme.ColorNameForegroundOnPrimary},
		{"ColorNameForegroundOnSuccess", theme.ColorNameForegroundOnSuccess},
		{"ColorNameForegroundOnWarning", theme.ColorNameForegroundOnWarning},
		{"ColorNameHeaderBackground", theme.ColorNameHeaderBackground},
		{"ColorNameHover", theme.ColorNameHover},
		{"ColorNameHyperlink", theme.ColorNameHyperlink},
		{"ColorNameInputBackground", theme.ColorNameInputBackground},
		{"ColorNameInputBorder", theme.ColorNameInputBorder},
		{"ColorNameMenuBackground", theme.ColorNameMenuBackground},
		{"ColorNameOverlayBackground", theme.ColorNameOverlayBackground},
		{"ColorNamePlaceHolder", theme.ColorNamePlaceHolder},
		{"ColorNamePressed", theme.ColorNamePressed},
		{"ColorNamePrimary", theme.ColorNamePrimary},
		{"ColorNameScrollBar", theme.ColorNameScrollBar},
		{"ColorNameSelection", theme.ColorNameSelection},
		{"ColorNameSeparator", theme.ColorNameSeparator},
		{"ColorNameShadow", theme.ColorNameShadow},
		{"ColorNameSuccess", theme.ColorNameSuccess},
		{"ColorNameWarning", theme.ColorNameWarning},
	}
	colorsFiltered := slices.Clone(colors)
	list := widget.NewList(
		func() int {
			return len(colorsFiltered)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("Template"),
				layout.NewSpacer(),
				canvas.NewRectangle(color.Transparent),
			)
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			if id >= len(colorsFiltered) {
				return
			}
			c := colorsFiltered[id]
			row := co.(*fyne.Container).Objects
			label := row[0].(*widget.Label)
			r := row[2].(*canvas.Rectangle)
			label.SetText(c.label)
			r.FillColor = theme.Color(fyne.ThemeColorName(c.name))
			r.SetMinSize(fyne.NewSize(100, 30))
			r.StrokeColor = theme.Color(theme.ColorNameForeground)
			r.StrokeWidth = 1.6
		},
	)
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Search...")
	entry.OnChanged = func(s string) {
		colorsFiltered = make([]colorRow, 0)
		s2 := strings.ToLower(s)
		for _, c := range colors {
			if strings.Contains(strings.ToLower(c.label), s2) {
				colorsFiltered = append(colorsFiltered, c)
			}
		}
		list.Refresh()
	}
	return container.NewBorder(
		entry,
		nil,
		nil,
		nil,
		list,
	)
}
