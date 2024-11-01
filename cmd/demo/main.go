// Demo is a Fyne app for demonstrating the features provided by the fyne-kx library.
package main

import (
	"log"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

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
	switch1 := kxwidget.NewSwitch(func(on bool) {
		log.Println("Switch 1: ", on)
	})
	switch1.SetState(true)
	switch2 := kxwidget.NewSwitch(func(on bool) {
		log.Println("Switch 2: ", on)
	})
	switch3 := kxwidget.NewSwitch(nil)
	switch3.Disable()
	switch4 := kxwidget.NewSwitch(nil)
	switch4.SetState(true)
	switch4.Disable()
	f := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Badge", Widget: badge},
			{Text: "", Widget: container.NewPadded()},
			{Text: "Slider", Widget: slider},
			{Text: "", Widget: container.NewPadded()},
			{Text: "Switch", Widget: container.NewVBox(switch1, switch2, switch3, switch4)},
			{Text: "", Widget: container.NewPadded()},
			{Text: "TappableIcon", Widget: container.NewHBox(icon)},
			{Text: "", Widget: container.NewPadded()},
			{Text: "TappableImage", Widget: container.NewHBox(img)},
			{Text: "", Widget: container.NewPadded()},
			{Text: "TappableLabel", Widget: label},
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
