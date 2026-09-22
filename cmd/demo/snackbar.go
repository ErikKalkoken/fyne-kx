package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func makeSnackbar(w fyne.Window) fyne.CanvasObject {
	sb := kxwidget.NewSnackbar(w.Canvas())
	sb.Start()

	b1 := widget.NewButton("Show message", func() {
		sb.Display("This is a snackbar message.")
	})

	b2 := widget.NewButton("Show message with long timeout", func() {
		sb.DisplayWithTimeout("This message stays for 10 seconds.", 10*time.Second)
	})

	b3 := widget.NewButton("Show wrapping message", func() {
		sb.Display("This is a much longer snackbar message that is expected to wrap onto multiple lines once it no longer fits within the width of the window.")
	})

	b4 := widget.NewButton("Show several messages", func() {
		for i := 1; i <= 3; i++ {
			sb.Display(fmt.Sprintf("Queued message #%d", i))
		}
	})

	return container.NewVBox(b1, b2, b3, b4)
}
