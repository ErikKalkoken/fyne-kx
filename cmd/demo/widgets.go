package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func makeBadge() fyne.CanvasObject {
	badges := container.NewVBox()
	badgesConfig := []struct {
		name       string
		importance widget.Importance
	}{
		{"danger", widget.DangerImportance},
		{"high", widget.HighImportance},
		{"low", widget.LowImportance},
		{"medium", widget.MediumImportance},
		{"success", widget.SuccessImportance},
		{"warning", widget.WarningImportance},
	}
	for _, bc := range badgesConfig {
		b := kxwidget.NewBadge("Alpha")
		b.Importance = bc.importance
		badges.Add(container.NewHBox(b, widget.NewLabel(bc.name+" importance")))
	}
	return badges
}

func makeSlider() fyne.CanvasObject {
	slider := kxwidget.NewSlider(0, 100)
	slider.SetValue(25)
	return slider
}

func makeSwitch() fyne.CanvasObject {
	textForBool := func(b bool) string {
		if b {
			return "on"
		}
		return "off"
	}
	switchLabel1 := widget.NewLabel("")
	switch1 := kxwidget.NewSwitch(func(on bool) {
		switchLabel1.SetText(textForBool(on))
	})
	switch1.On = true
	switchLabel1.Text = textForBool(switch1.State())
	switch1Box := container.NewHBox(switch1, switchLabel1)

	switchLabel2 := widget.NewLabel("")
	switch2 := kxwidget.NewSwitch(func(on bool) {
		switchLabel2.SetText(textForBool(on))
	})
	switchLabel2.Text = textForBool(switch2.State())
	switch2Box := container.NewHBox(switch2, switchLabel2)

	switch3 := kxwidget.NewSwitch(nil)
	switch3.On = true
	switch3.Disable()
	switch4 := kxwidget.NewSwitch(nil)
	switch4.Disable()
	addLabel := func(c fyne.CanvasObject, text string) fyne.CanvasObject {
		return container.NewHBox(c, widget.NewLabel(text))
	}

	return container.NewVBox(
		switch1Box,
		switch2Box,
		addLabel(switch3, "on disabled"),
		addLabel(switch4, "off disabled"),
	)
}

func makeTappableImage() fyne.CanvasObject {
	imgStandard := kxwidget.NewTappableImage(resourceIconPng, func() {
		log.Println("TappableImage tapped")
	})
	imgStandard.SetFillMode(canvas.ImageFillContain)
	imgStandard.SetMinSize(fyne.NewSize(100, 100))

	menu := fyne.NewMenu(
		"",
		fyne.NewMenuItem("First", func() {
			log.Println("first selected")
		}),
		fyne.NewMenuItem("Second", func() {
			log.Println("second selected")
		}),
	)
	im1 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im1.SetFillMode(canvas.ImageFillContain)
	im1.SetMinSize(fyne.NewSize(100, 100))
	im2 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im2.SetFillMode(canvas.ImageFillContain)
	im2.SetMinSize(fyne.NewSize(100, 100))
	im3 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im3.SetFillMode(canvas.ImageFillContain)
	im3.SetMinSize(fyne.NewSize(100, 100))
	im4 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im4.SetFillMode(canvas.ImageFillContain)
	im4.SetMinSize(fyne.NewSize(100, 100))
	return container.NewBorder(im1, im2, im3, im4, imgStandard)
}

func makeTappableIcon() fyne.CanvasObject {
	icon := kxwidget.NewTappableIcon(theme.AccountIcon(), func() {
		log.Println("TappableIcon tapped")
	})
	return container.NewVBox(icon)
}

func makeTappableLabel() fyne.CanvasObject {
	label := kxwidget.NewTappableLabel("Tap me", func() {
		log.Println("TappableLabel tapped")
	})
	return label
}
