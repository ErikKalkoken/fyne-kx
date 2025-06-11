package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	switchLabel1.Text = textForBool(switch1.On)
	switch1Box := container.NewHBox(switch1, switchLabel1)

	switchLabel2 := widget.NewLabel("")
	switch2 := kxwidget.NewSwitch(func(on bool) {
		switchLabel2.SetText(textForBool(on))
	})
	switchLabel2.Text = textForBool(switch2.On)
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
	size := fyne.NewSize(100, 100)
	imgStandard := kxwidget.NewTappableImage(resourceIconPng, func() {
		log.Println("TappableImage tapped")
	})
	imgStandard.SetFillMode(canvas.ImageFillContain)
	imgStandard.SetMinSize(size)

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
	im1.SetMinSize(size)
	im2 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im2.SetFillMode(canvas.ImageFillContain)
	im2.SetMinSize(size)
	im3 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im3.SetFillMode(canvas.ImageFillContain)
	im3.SetMinSize(size)
	im4 := kxwidget.NewTappableImageWithMenu(resourceIconPng, menu)
	im4.SetFillMode(canvas.ImageFillContain)
	im4.SetMinSize(size)
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
	return container.NewHBox(label, widget.NewLabel("<- tap"))
}

func makeToolbarActionMenu() fyne.CanvasObject {
	menu := kxwidget.NewToolbarActionMenu(theme.MenuIcon(), fyne.NewMenu(
		"",
		fyne.NewMenuItem("First", func() {
			log.Println("first selected")
		}),
		fyne.NewMenuItem("Second", func() {
			log.Println("second selected")
		}),
	))
	ntb := widget.NewToolbar(menu, widget.NewToolbarAction(theme.AccountIcon(), func() {
		log.Println("Account tapped")
	}))
	return container.NewVBox(ntb)
}

func makeFilterChip() fyne.CanvasObject {
	c1 := kxwidget.NewFilterChip("Disabled Off", nil)
	c1.Disable()
	c2 := kxwidget.NewFilterChip("Disabled On", nil)
	c2.On = true
	c2.Disable()
	c := container.NewVBox(
		kxwidget.NewFilterChip("Alpha", func(on bool) {
			log.Printf("Alpha: %v\n", on)
		}),
		c1,
		c2,
	)
	return c
}

func makeFilterChipGroup() fyne.CanvasObject {
	options := []string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel"}
	g := kxwidget.NewFilterChipGroup(options, func(s []string) {
		log.Println(s)
	})
	g.Selected = []string{"Bravo", "Golf"}
	c := container.NewVBox(g)
	return c
}

func makeFilterChipSelect(w fyne.Window) fyne.CanvasObject {
	options := []string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel"}
	s1 := kxwidget.NewFilterChipSelect("DropDown", options, func(s string) {
		log.Printf("DropDown: %s\n", s)
	})
	s2 := kxwidget.NewFilterChipSelectWithSearch("Search", options, func(s string) {
		log.Printf("Search: %s\n", s)
	},
		w,
	)
	s3 := kxwidget.NewFilterChipSelect("Disabled", options, nil)
	s3.Disable()
	c := container.NewVBox(s1, s2, s3)
	return c
}

func makeIconButton() fyne.CanvasObject {
	i1 := kxwidget.NewIconButton(theme.AccountIcon(), func() {
		log.Println("IconButton tapped")
	})
	i2 := kxwidget.NewIconButton(theme.AccountIcon(), nil)
	i2.Disable()
	c := container.NewVBox(
		container.NewHBox(i1, widget.NewLabel("Enabled"), layout.NewSpacer()),
		container.NewHBox(i2, widget.NewLabel("Disabled"), layout.NewSpacer()),
	)
	return c
}
