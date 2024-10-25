// App for demonstrating the features provided by the Fyne KX extension.
package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	kxmodal "github.com/ErikKalkoken/fyne-kx/modal"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("KX Demo")
	tabs := container.NewAppTabs(
		container.NewTabItem("Modals", makeModals(w)),
		container.NewTabItem("Widgets", makeWidgets()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

func makeWidgets() fyne.CanvasObject {
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
	f := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Badge",
				Widget: kxwidget.NewBadge("1234"),
			},
			{
				Text:   "SliderWithValue",
				Widget: kxwidget.NewSliderWithValue(0, 50),
			},
			{
				Text:   "TappableIcon",
				Widget: icon,
			},
			{
				Text:   "TappableImage",
				Widget: img,
			},
			{
				Text:   "TappableLabel",
				Widget: label,
			},
		},
	}
	return f
}

func makeModals(w fyne.Window) *fyne.Container {
	b1 := widget.NewButton("ProgressModal", func() {
		m := kxmodal.NewProgress("ProgressModal", "Please wait...", func(progress binding.Float) error {
			for i := 1; i < 10; i++ {
				progress.Set(float64(i))
				time.Sleep(250 * time.Millisecond)
			}
			return nil
		}, 10, w)
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
			for i := 1; i < 10; i++ {
				time.Sleep(250 * time.Millisecond)
			}
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
