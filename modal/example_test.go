package modal_test

import (
	"context"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	kxmodal "github.com/ErikKalkoken/fyne-kx/modal"
)

func ExampleNewProgress() {
	a := app.New()
	w := a.NewWindow("Progress Modal")

	button := widget.NewButton("Load file", func() {
		m := kxmodal.NewProgress("Loading file", "Loading file XX. Please wait.", func(progress binding.Float, done func(error)) {
			go func() {
				for i := 1; i <= 10; i++ {
					fyne.Do(func() {
						progress.Set(float64(i))
					})
					time.Sleep(300 * time.Millisecond) // simulate a chunk of work
				}
				done(nil)
			}()
		}, 10, w)
		m.Start()
	})

	w.SetContent(container.NewCenter(button))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

func ExampleNewProgressWithCancel() {
	a := app.New()
	w := a.NewWindow("Progress Modal With Cancel")

	button := widget.NewButton("Load file", func() {
		m := kxmodal.NewProgressWithCancel("Loading file", "Loading file XX. Please wait.", func(progress binding.Float, onCancel func(func()), done func(error)) {
			ctx, cancel := context.WithCancel(context.Background())
			onCancel(cancel)

			go func() {
				for i := 1; i <= 10; i++ {
					select {
					case <-ctx.Done():
						done(ctx.Err())
						return
					default:
					}
					fyne.Do(func() {
						progress.Set(float64(i))
					})
					time.Sleep(300 * time.Millisecond) // simulate a chunk of work
				}
				done(nil)
			}()
		}, 10, w)
		m.Start()
	})

	w.SetContent(container.NewCenter(button))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

func ExampleNewProgressInfinite() {
	a := app.New()
	w := a.NewWindow("Progress Infinite Modal")

	button := widget.NewButton("Load file", func() {
		m := kxmodal.NewProgressInfinite("Loading file", "Loading file XX. Please wait.", func(done func(error)) {
			go func() {
				time.Sleep(3 * time.Second) // simulate a long running process
				done(nil)
			}()
		}, w)
		m.Start()
	})

	w.SetContent(container.NewCenter(button))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

func ExampleNewProgressInfiniteWithCancel() {
	a := app.New()
	w := a.NewWindow("Progress Infinite Modal With Cancel")

	button := widget.NewButton("Load file", func() {
		m := kxmodal.NewProgressInfiniteWithCancel("Loading file", "Loading file XX. Please wait.", func(onCancel func(func()), done func(error)) {
			ctx, cancel := context.WithCancel(context.Background())
			onCancel(cancel)

			go func() {
				select {
				case <-ctx.Done():
					done(ctx.Err())
				case <-time.After(3 * time.Second): // simulate a long running process
					done(nil)
				}
			}()
		}, w)
		m.Start()
	})

	w.SetContent(container.NewCenter(button))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
