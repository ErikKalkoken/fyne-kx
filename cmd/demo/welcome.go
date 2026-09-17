package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func parseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	return u
}

func makeWelcome() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("fyne-kx", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.SizeName = theme.SizeNameHeadingText

	description := widget.NewLabelWithStyle(
		"A library with extensions and tools for the Fyne GUI toolkit.\n\nSelect an item from the list on the left to see it in action.",
		fyne.TextAlignCenter,
		fyne.TextStyle{},
	)

	footer := container.NewHBox(
		layout.NewSpacer(),
		widget.NewHyperlink("GitHub", parseURL("https://github.com/ErikKalkoken/fyne-kx")),
		widget.NewLabel("-"),
		widget.NewHyperlink("Go Reference", parseURL("https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx")),
		widget.NewLabel("-"),
		widget.NewHyperlink("Fyne Documentation", parseURL("https://docs.fyne.io/")),
		layout.NewSpacer(),
	)

	content := container.NewCenter(container.NewVBox(
		title,
		description,
	))

	return container.NewBorder(nil, footer, nil, nil, content)
}
