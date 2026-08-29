// Fynetheme is a Fyne app for showing details about the default Fyne theme.
package main

import (
	"cmp"
	"fmt"
	"image/color"
	"slices"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	kxtheme "github.com/ErikKalkoken/fyne-kx/theme"
	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

const (
	iconSizeStart = 40
)

func main() {
	app := app.NewWithID("io.github.erikkalkoken.fyne-kx.fynetheme")
	w := app.NewWindow("Fyne Theme Browser")
	tabs := container.NewAppTabs(
		container.NewTabItem("Colors", makeColors()),
		container.NewTabItem("Icons", makeIcons()),
		container.NewTabItem("Sizes", makeSizes()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	theme := widget.NewSelect([]string{"Auto", "Light", "Dark"}, func(s string) {
		switch s {
		case "Light":
			app.Settings().SetTheme(kxtheme.DefaultWithFixedVariant(theme.VariantLight))
		case "Dark":
			app.Settings().SetTheme(kxtheme.DefaultWithFixedVariant(theme.VariantDark))
		default:
			app.Settings().SetTheme(theme.DefaultTheme())
		}

	})
	theme.SetSelected("Auto")
	bottom := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Theme"),
			theme,
		),
	)

	w.SetContent(container.NewBorder(
		nil,
		bottom,
		nil,
		nil,
		tabs,
	))
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}

func makeColors() fyne.CanvasObject {
	hasTransparencyDark := make(map[fyne.ThemeColorName]bool)
	hasTransparencyLight := make(map[fyne.ThemeColorName]bool)
	th := theme.Current()
	hasTransparency := func(name fyne.ThemeColorName, v fyne.ThemeVariant) bool {
		c := th.Color(fyne.ThemeColorName(name), v)
		_, _, _, a := c.RGBA()
		return a != 0xffff
	}
	for _, col := range colors {
		if hasTransparency(col.name, theme.VariantDark) {
			hasTransparencyDark[col.name] = true
		}
		if hasTransparency(col.name, theme.VariantLight) {
			hasTransparencyLight[col.name] = true
		}
	}

	var rowsFiltered []colorRow

	list := widget.NewList(
		func() int {
			return len(rowsFiltered)
		},
		func() fyne.CanvasObject {
			check1 := widget.NewCheck("", nil)
			check1.Disable()
			check2 := widget.NewCheck("", nil)
			check2.Disable()
			return container.NewHBox(
				widget.NewLabel("Template"),
				layout.NewSpacer(),
				canvas.NewRectangle(color.Transparent),
				check1,
				canvas.NewRectangle(color.Transparent),
				check2,
			)
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			if id >= len(rowsFiltered) {
				return
			}
			myColor := rowsFiltered[id]
			row := co.(*fyne.Container).Objects

			label := row[0].(*widget.Label)
			label.SetText(myColor.label)

			boxSize := fyne.NewSize(100, 30)
			const borderSize = 7.5
			colorRect1 := row[2].(*canvas.Rectangle)
			colorRect1.FillColor = th.Color(fyne.ThemeColorName(myColor.name), theme.VariantLight)
			colorRect1.SetMinSize(boxSize)
			colorRect1.StrokeColor = th.Color(fyne.ThemeColorName(theme.ColorNameBackground), theme.VariantLight)
			colorRect1.StrokeWidth = borderSize

			check1 := row[3].(*widget.Check)
			check1.SetChecked(hasTransparencyLight[myColor.name])

			colorRect2 := row[4].(*canvas.Rectangle)
			colorRect2.FillColor = th.Color(fyne.ThemeColorName(myColor.name), theme.VariantDark)
			colorRect2.SetMinSize(boxSize)
			colorRect2.StrokeColor = th.Color(fyne.ThemeColorName(theme.ColorNameBackground), theme.VariantDark)
			colorRect2.StrokeWidth = borderSize

			check2 := row[5].(*widget.Check)
			check2.SetChecked(hasTransparencyDark[myColor.name])
		},
	)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search...")
	transparencyFilter := kxwidget.NewFilterChipSelect("Transparency", []string{"Transparent", "Opaque"}, nil)
	sortChip := kxwidget.NewSortChip([]string{"Name"}, "Name", kxwidget.SortOrderAscending, nil)
	footerLabel := widget.NewLabel("")

	filterRows := func() {
		rows := slices.Clone(colors)

		if x := searchEntry.Text; x != "" {
			x = strings.ToLower(x)
			rows = slices.DeleteFunc(rows, func(r colorRow) bool {
				return !strings.Contains(strings.ToLower(r.label), x)
			})
		}

		if x := transparencyFilter.Selected; x != "" {
			rows = slices.DeleteFunc(rows, func(r colorRow) bool {
				switch x {
				case "Transparent":
					return !hasTransparencyLight[r.name]
				case "Opaque":
					return hasTransparencyLight[r.name]
				}
				return false
			})
		}

		slices.SortFunc(rows, func(a, b colorRow) int {
			c := sortChip.Column
			o := sortChip.Order
			switch {
			case c == "Name" && o == kxwidget.SortOrderAscending:
				return strings.Compare(a.label, b.label)
			case c == "Name" && o == kxwidget.SortOrderDescending:
				return strings.Compare(b.label, a.label)
			}
			return 0
		})

		rowsFiltered = rows
		list.Refresh()
		footerLabel.SetText(fmt.Sprintf("Showing %d / %d entries", len(rows), len(colors)))
	}

	searchEntry.OnChanged = func(s string) {
		filterRows()
	}
	sortChip.OnChanged = func(column string, order kxwidget.SortOrder) {
		filterRows()
	}
	transparencyFilter.OnChanged = func(s string) {
		filterRows()
	}
	filterRows()

	return container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			nil,
			container.NewHBox(transparencyFilter, sortChip),
			searchEntry,
		),
		footerLabel,
		nil,
		nil,
		list,
	)
}

func makeSizes() fyne.CanvasObject {
	var rowsFiltered []sizeRow

	list := widget.NewList(
		func() int {
			return len(rowsFiltered)
		},
		func() fyne.CanvasObject {
			size := widget.NewLabel("999")
			size.Alignment = fyne.TextAlignTrailing
			return container.NewHBox(
				widget.NewLabel("Template"),
				layout.NewSpacer(),
				size,
			)
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			if id >= len(rowsFiltered) {
				return
			}
			s := rowsFiltered[id]
			row := co.(*fyne.Container).Objects
			label := row[0].(*widget.Label)
			label.SetText(s.label)
			size := row[2].(*widget.Label)
			v := theme.Size(s.name)
			size.SetText(fmt.Sprint(v))
		},
	)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search...")
	sortChip := kxwidget.NewSortChip([]string{"Name", "Size"}, "Name", kxwidget.SortOrderAscending, nil)
	footerLabel := widget.NewLabel("")

	filterRows := func() {
		rows := slices.Clone(sizes)

		if x := searchEntry.Text; x != "" {
			x = strings.ToLower(x)
			rows = slices.DeleteFunc(rows, func(r sizeRow) bool {
				return !strings.Contains(strings.ToLower(r.label), x)
			})
		}

		slices.SortFunc(rows, func(a, b sizeRow) int {
			c := sortChip.Column
			o := sortChip.Order
			switch {
			case c == "Name" && o == kxwidget.SortOrderAscending:
				return strings.Compare(a.label, b.label)
			case c == "Name" && o == kxwidget.SortOrderDescending:
				return strings.Compare(b.label, a.label)
			case c == "Size" && o == kxwidget.SortOrderAscending:
				return cmp.Compare(theme.Size(a.name), theme.Size(b.name))
			case c == "Size" && o == kxwidget.SortOrderDescending:
				return cmp.Compare(theme.Size(b.name), theme.Size(a.name))
			}
			return 0
		})

		rowsFiltered = rows
		list.Refresh()
		footerLabel.SetText(fmt.Sprintf("Showing %d / %d entries", len(rows), len(sizes)))
	}

	searchEntry.OnChanged = func(s string) {
		filterRows()
	}
	sortChip.OnChanged = func(col string, order kxwidget.SortOrder) {
		filterRows()
	}
	filterRows()

	return container.NewBorder(
		container.NewBorder(nil, nil, nil, sortChip, searchEntry),
		footerLabel,
		nil,
		nil,
		list,
	)
}

func makeIcons() fyne.CanvasObject {
	var rowsFiltered []iconRow

	var iconSize float32 = iconSizeStart
	iconColors := []string{"Default", "Disabled", "Error", "Primary", "Success", "Warning"}
	var iconColor = "Default"

	grid := widget.NewGridWrap(
		func() int {
			return len(rowsFiltered)
		},
		func() fyne.CanvasObject {
			image := canvas.NewImageFromResource(theme.BrokenImageIcon())
			image.FillMode = canvas.ImageFillContain
			image.SetMinSize(fyne.NewSquareSize(iconSize))
			label := widget.NewLabel("IconNameRadioButtonChecked")
			label.Alignment = fyne.TextAlignCenter
			return container.NewBorder(
				nil,
				container.NewVBox(label, container.NewPadded()),
				nil,
				nil,
				image,
			)
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			if id >= len(rowsFiltered) {
				return
			}
			s := rowsFiltered[id]
			c := co.(*fyne.Container).Objects
			image := c[0].(*canvas.Image)
			r := theme.Icon(s.name)
			switch iconColor {
			case "Disabled":
				image.Resource = theme.NewDisabledResource(r)
			case "Error":
				image.Resource = theme.NewErrorThemedResource(r)
			case "Primary":
				image.Resource = theme.NewPrimaryThemedResource(r)
			case "Success":
				image.Resource = theme.NewSuccessThemedResource(r)
			case "Warning":
				image.Resource = theme.NewWarningThemedResource(r)
			default:
				image.Resource = theme.NewThemedResource(r)
			}
			image.Refresh()
			label := c[1].(*fyne.Container).Objects[0].(*widget.Label)
			label.SetText(s.label)
		},
	)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search...")
	footerLabel := widget.NewLabel("")
	sortChip := kxwidget.NewSortChip([]string{"Name"}, "Name", kxwidget.SortOrderAscending, nil)

	filterRows := func() {
		rows := slices.Clone(icons)

		if x := searchEntry.Text; x != "" {
			x = strings.ToLower(x)
			rows = slices.DeleteFunc(rows, func(r iconRow) bool {
				return !strings.Contains(strings.ToLower(r.label), x)
			})
		}

		slices.SortFunc(rows, func(a, b iconRow) int {
			c := sortChip.Column
			o := sortChip.Order
			switch {
			case c == "Name" && o == kxwidget.SortOrderAscending:
				return strings.Compare(a.label, b.label)
			case c == "Name" && o == kxwidget.SortOrderDescending:
				return strings.Compare(b.label, a.label)
			}
			return 0
		})

		rowsFiltered = rows
		grid.Refresh()
		footerLabel.SetText(fmt.Sprintf("Showing %d / %d entries", len(rows), len(icons)))
	}

	searchEntry.OnChanged = func(s string) {
		filterRows()
	}
	sortChip.OnChanged = func(column string, order kxwidget.SortOrder) {
		filterRows()
	}
	filterRows()

	slider := widget.NewSlider(2, 128)
	slider.Step = 4
	slider.OnChanged = func(v float64) {
		iconSize = float32(v)
		grid.Refresh()
	}
	slider.SetValue(float64(iconSize))
	sliderBox := container.NewBorder(nil, nil, widget.NewLabel("Size"), nil, slider)
	themeSelect := widget.NewSelect(iconColors, func(s string) {
		iconColor = s
		grid.Refresh()
	})
	themeSelect.SetSelected("Default")
	themeBox := container.NewHBox(widget.NewLabel("Color"), themeSelect)
	return container.NewBorder(
		container.NewVBox(
			container.NewBorder(nil, nil, nil, sortChip, searchEntry),
			container.NewGridWithColumns(2, themeBox, sliderBox),
		),
		footerLabel,
		nil,
		nil,
		grid,
	)
}
