package main

import (
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	kxlayout "github.com/ErikKalkoken/fyne-kx/layout"
)

func makeColumns() fyne.CanvasObject {
	layout := kxlayout.NewColumns(150, 100, 50)
	makeBox := func(h float32) fyne.CanvasObject {
		x := canvas.NewRectangle(theme.Color(theme.ColorNameInputBorder))
		w := rand.Float32()*100 + 50
		x.SetMinSize(fyne.NewSize(w, h))
		return x
	}
	return container.NewVBox(
		container.New(layout, makeBox(50), makeBox(50), makeBox(50)),
		container.New(layout, makeBox(150), makeBox(150), makeBox(150)),
		container.New(layout, makeBox(30), makeBox(30), makeBox(30)),
	)
}

func makeRowWrap() fyne.CanvasObject {
	layout := kxlayout.NewRowWrapLayout()
	makeBox := func() fyne.CanvasObject {
		x := canvas.NewRectangle(theme.Color(theme.ColorNameInputBorder))
		w := rand.Float32()*150 + 20
		x.SetMinSize(fyne.NewSize(w, 50))
		return x
	}
	c := container.New(layout)
	for i := 0; i < 20; i++ {
		c.Add(makeBox())
	}
	return c
}
