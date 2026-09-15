package widget_test

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func ExampleNewProgressButton() {
	a := app.New()
	w := a.NewWindow("Progress Button")

	button := kxwidget.NewProgressButton("Load file", nil, func(done func()) {
		go func() {
			time.Sleep(3 * time.Second) // simulate a long running process
			done()
		}()
	})

	w.SetContent(container.NewCenter(button))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
