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
	makeTopic := func(title string, topic fyne.CanvasObject) fyne.CanvasObject {
		label := widget.NewLabel(title)
		label.TextStyle.Bold = true
		c := container.NewBorder(
			label,
			nil,
			nil,
			nil,
			topic,
		)
		return c
	}

	importance := container.New(layout.NewRowWrapLayout())
	defineImportance := []struct {
		name       string
		importance widget.Importance
	}{
		{"Danger", widget.DangerImportance},
		{"High", widget.HighImportance},
		{"Low", widget.LowImportance},
		{"Medium", widget.MediumImportance},
		{"Success", widget.SuccessImportance},
		{"Warning", widget.WarningImportance},
	}
	for _, bc := range defineImportance {
		b := kxwidget.NewBadge(bc.name)
		b.Importance = bc.importance
		importance.Add(b)
	}
	t1 := makeTopic("Importance", importance)

	alignment := container.NewVBox()
	defineAlignment := []struct {
		name  string
		align fyne.TextAlign
	}{
		{"Leading", fyne.TextAlignLeading},
		{"Center", fyne.TextAlignCenter},
		{"Trailing", fyne.TextAlignTrailing},
	}
	for _, c := range defineAlignment {
		b := kxwidget.NewBadge(c.name)
		b.Alignment = c.align
		b.Importance = widget.HighImportance
		alignment.Add(b)
	}
	t2 := makeTopic("Alignment", alignment)

	sizes := container.NewVBox()
	defineSizes := []struct {
		name     string
		sizeName fyne.ThemeSizeName
	}{
		{"Text", theme.SizeNameText},
		{"Caption", theme.SizeNameCaptionText},
		{"Subheading", theme.SizeNameSubHeadingText},
	}
	for _, c := range defineSizes {
		b := kxwidget.NewBadge(c.name)
		b.SizeName = c.sizeName
		b.Importance = widget.WarningImportance
		sizes.Add(b)
	}
	t3 := makeTopic("Sizes", sizes)

	// b1 := kxwidget.NewBadge("Alpha")
	// b1.Importance = widget.HighImportance
	// b2 := kxwidget.NewBadge("Bravo")
	// b2.Importance = widget.HighImportance
	// other := container.NewVBox(
	// 	widget.NewLabel("Alpha"),
	// 	b1,
	// 	container.NewHBox(widget.NewLabel("Bravo"), b2),
	// )
	// t4 := makeTopic("Label comparison", other)

	return container.NewVBox(t1, t2, t3)
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
	c1 := kxwidget.NewFilterChip("Charlie", nil)
	c1.Disable()
	c2 := kxwidget.NewFilterChip("Delta", nil)
	c2.On = true
	c2.Disable()
	c3 := kxwidget.NewFilterChip("Bravo", func(on bool) {
		log.Printf("Bravo: %v\n", on)
	})
	c3.On = true
	c := container.NewVBox(
		kxwidget.NewFilterChip("Alpha", func(on bool) {
			log.Printf("Alpha: %v\n", on)
		}),
		c3,
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
	b := widget.NewButton("Disable", nil)
	b.OnTapped = func() {
		if g.Disabled() {
			g.Enable()
			b.SetText("Disable")
		} else {
			g.Disable()
			b.SetText("Enable")
		}
	}
	c := container.NewVBox(g, container.NewPadded(), b)
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
	b := widget.NewButton("Disable", nil)
	b.OnTapped = func() {
		if s1.Disabled() {
			s1.Enable()
			s2.Enable()
			b.SetText("Disable")
		} else {
			s1.Disable()
			s2.Disable()
			b.SetText("Enable")
		}
	}
	c := container.NewVBox(s1, s2, container.NewPadded(), b)
	return c
}

func makeIconButton() fyne.CanvasObject {
	i1 := kxwidget.NewIconButton(theme.AccountIcon(), func() {
		log.Println("IconButton tapped")
	})
	i2 := kxwidget.NewIconButton(theme.AccountIcon(), nil)
	i2.Disable()
	makeBorder := func() fyne.CanvasObject {
		r := canvas.NewRectangle(theme.Color(theme.ColorNameInputBackground))
		r.SetMinSize(fyne.NewSquareSize(theme.Padding() * 5))
		return r
	}
	i3 := kxwidget.NewIconButtonWithMenu(theme.FolderIcon(), fyne.NewMenu("",
		fyne.NewMenuItem("first", nil),
		fyne.NewMenuItem("second", nil),
	))
	c := container.NewGridWithRows(
		3,
		container.NewBorder(
			makeBorder(),
			makeBorder(),
			makeBorder(),
			container.NewHBox(makeBorder(), widget.NewLabel("Enabled")),
			i1,
		),
		container.NewBorder(
			makeBorder(),
			makeBorder(),
			makeBorder(),
			container.NewHBox(makeBorder(), widget.NewLabel("Disabled")),
			i2,
		),

		container.NewHBox(
			container.NewCenter(i3),
			layout.NewSpacer(),
			widget.NewLabel("With menu"),
		),
	)
	return c
}

func makeSortChip() fyne.CanvasObject {
	cols := []string{"Name", "Title", "Age", "Location"}
	s1 := kxwidget.NewSortChip(cols, "Name", kxwidget.SortOrderAscending, func(c string, o kxwidget.SortOrder) {
		log.Printf("Sort 1: column %s, descending %s\n", c, o)
	})
	s2 := kxwidget.NewSortChip(cols, "Name", kxwidget.SortOrderAscending, func(c string, o kxwidget.SortOrder) {
		log.Printf("Sort 2: column %s, descending %s\n", c, o)
	})
	s2.Column = "Title"
	s2.Order = kxwidget.SortOrderDescending

	b1 := widget.NewButton("Disable", nil)
	b1.OnTapped = func() {
		if s1.Disabled() {
			s1.Enable()
			s2.Enable()
			b1.SetText("Disable")
		} else {
			s1.Disable()
			s2.Disable()
			b1.SetText("Enable")
		}
	}
	c := container.NewVBox(s1, s2, container.NewPadded(), b1)
	return c
}
