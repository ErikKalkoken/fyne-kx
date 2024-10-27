// Themeinsight is a Fyne app for showing details about the default Fyne theme.
package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

const (
	iconSizeStart = 40
)

func main() {
	app := app.New()
	w := app.NewWindow("Theme Insight")
	tabs := container.NewAppTabs(
		container.NewTabItem("Colors", makeColors()),
		container.NewTabItem("Icons", makeIcons()),
		container.NewTabItem("Sizes", makeSizes()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		tabs,
	))
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}

type colorRow struct {
	label string
	name  fyne.ThemeColorName
}

func makeColors() fyne.CanvasObject {
	colors := []colorRow{
		{"ColorNameBackground", theme.ColorNameBackground},
		{"ColorNameButton", theme.ColorNameButton},
		{"ColorNameDisabled", theme.ColorNameDisabled},
		{"ColorNameDisabledButton", theme.ColorNameDisabledButton},
		{"ColorNameError", theme.ColorNameError},
		{"ColorNameFocus", theme.ColorNameFocus},
		{"ColorNameForeground", theme.ColorNameForeground},
		{"ColorNameForegroundOnError", theme.ColorNameForegroundOnError},
		{"ColorNameForegroundOnPrimary", theme.ColorNameForegroundOnPrimary},
		{"ColorNameForegroundOnSuccess", theme.ColorNameForegroundOnSuccess},
		{"ColorNameForegroundOnWarning", theme.ColorNameForegroundOnWarning},
		{"ColorNameHeaderBackground", theme.ColorNameHeaderBackground},
		{"ColorNameHover", theme.ColorNameHover},
		{"ColorNameHyperlink", theme.ColorNameHyperlink},
		{"ColorNameInputBackground", theme.ColorNameInputBackground},
		{"ColorNameInputBorder", theme.ColorNameInputBorder},
		{"ColorNameMenuBackground", theme.ColorNameMenuBackground},
		{"ColorNameOverlayBackground", theme.ColorNameOverlayBackground},
		{"ColorNamePlaceHolder", theme.ColorNamePlaceHolder},
		{"ColorNamePressed", theme.ColorNamePressed},
		{"ColorNamePrimary", theme.ColorNamePrimary},
		{"ColorNameScrollBar", theme.ColorNameScrollBar},
		{"ColorNameSelection", theme.ColorNameSelection},
		{"ColorNameSeparator", theme.ColorNameSeparator},
		{"ColorNameShadow", theme.ColorNameShadow},
		{"ColorNameSuccess", theme.ColorNameSuccess},
		{"ColorNameWarning", theme.ColorNameWarning},
	}
	slices.SortFunc(colors, func(a, b colorRow) int {
		return strings.Compare(a.label, b.label)
	})
	colorsFiltered := slices.Clone(colors)
	list := widget.NewList(
		func() int {
			return len(colorsFiltered)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("Template"),
				layout.NewSpacer(),
				canvas.NewRectangle(color.Transparent),
			)
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			if id >= len(colorsFiltered) {
				return
			}
			c := colorsFiltered[id]
			row := co.(*fyne.Container).Objects
			label := row[0].(*widget.Label)
			label.SetText(c.label)
			r := row[2].(*canvas.Rectangle)
			r.FillColor = theme.Color(fyne.ThemeColorName(c.name))
			r.SetMinSize(fyne.NewSize(100, 30))
			r.StrokeColor = theme.Color(theme.ColorNameForeground)
			r.StrokeWidth = 1.6
		},
	)
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Search...")
	entry.OnChanged = func(s string) {
		colorsFiltered = make([]colorRow, 0)
		s2 := strings.ToLower(s)
		for _, c := range colors {
			if strings.Contains(strings.ToLower(c.label), s2) {
				colorsFiltered = append(colorsFiltered, c)
			}
		}
		list.Refresh()
	}
	return container.NewBorder(
		entry,
		nil,
		nil,
		nil,
		list,
	)
}

type sizeRow struct {
	label string
	name  fyne.ThemeSizeName
}

func makeSizes() fyne.CanvasObject {
	sizes := []sizeRow{
		{"SizeNameCaptionText", theme.SizeNameCaptionText},
		{"SizeNameHeadingText", theme.SizeNameHeadingText},
		{"SizeNameInlineIcon", theme.SizeNameInlineIcon},
		{"SizeNameInnerPadding", theme.SizeNameInnerPadding},
		{"SizeNameInputBorder", theme.SizeNameInputBorder},
		{"SizeNameInputRadius", theme.SizeNameInputRadius},
		{"SizeNameLineSpacing", theme.SizeNameLineSpacing},
		{"SizeNamePadding", theme.SizeNamePadding},
		{"SizeNameScrollBar", theme.SizeNameScrollBar},
		{"SizeNameScrollBarRadius", theme.SizeNameScrollBarRadius},
		{"SizeNameScrollBarSmall", theme.SizeNameScrollBarSmall},
		{"SizeNameSelectionRadius", theme.SizeNameSelectionRadius},
		{"SizeNameSeparatorThickness", theme.SizeNameSeparatorThickness},
		{"SizeNameSubHeadingText", theme.SizeNameSubHeadingText},
		{"SizeNameText", theme.SizeNameText},
	}
	slices.SortFunc(sizes, func(a, b sizeRow) int {
		return strings.Compare(a.label, b.label)
	})
	sizesFiltered := slices.Clone(sizes)
	list := widget.NewList(
		func() int {
			return len(sizesFiltered)
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
			if id >= len(sizesFiltered) {
				return
			}
			s := sizesFiltered[id]
			row := co.(*fyne.Container).Objects
			label := row[0].(*widget.Label)
			label.SetText(s.label)
			size := row[2].(*widget.Label)
			v := theme.Size(s.name)
			size.SetText(fmt.Sprint(v))
		},
	)
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Search...")
	entry.OnChanged = func(s string) {
		sizesFiltered = make([]sizeRow, 0)
		s2 := strings.ToLower(s)
		for _, c := range sizes {
			if strings.Contains(strings.ToLower(c.label), s2) {
				sizesFiltered = append(sizesFiltered, c)
			}
		}
		list.Refresh()
	}
	return container.NewBorder(
		entry,
		nil,
		nil,
		nil,
		list,
	)
}

type iconRow struct {
	label string
	name  fyne.ThemeIconName
}

func makeIcons() fyne.CanvasObject {
	sizes := []iconRow{
		{"IconNameAccount", theme.IconNameAccount},
		{"IconNameArrowDropDown", theme.IconNameArrowDropDown},
		{"IconNameArrowDropUp", theme.IconNameArrowDropUp},
		{"IconNameBrokenImage", theme.IconNameBrokenImage},
		{"IconNameCancel", theme.IconNameCancel},
		{"IconNameCheckButton", theme.IconNameCheckButton},
		{"IconNameCheckButtonChecked", theme.IconNameCheckButtonChecked},
		{"IconNameCheckButtonFill", theme.IconNameCheckButtonFill},
		{"IconNameColorAchromatic", theme.IconNameColorAchromatic},
		{"IconNameColorChromatic", theme.IconNameColorChromatic},
		{"IconNameColorPalette", theme.IconNameColorPalette},
		{"IconNameComputer", theme.IconNameComputer},
		{"IconNameConfirm", theme.IconNameConfirm},
		{"IconNameContentAdd", theme.IconNameContentAdd},
		{"IconNameContentClear", theme.IconNameContentClear},
		{"IconNameContentCopy", theme.IconNameContentCopy},
		{"IconNameContentCut", theme.IconNameContentCut},
		{"IconNameContentPaste", theme.IconNameContentPaste},
		{"IconNameContentRedo", theme.IconNameContentRedo},
		{"IconNameContentRemove", theme.IconNameContentRemove},
		{"IconNameContentUndo", theme.IconNameContentUndo},
		{"IconNameDelete", theme.IconNameDelete},
		{"IconNameDesktop", theme.IconNameDesktop},
		{"IconNameDocument", theme.IconNameDocument},
		{"IconNameDocumentCreate", theme.IconNameDocumentCreate},
		{"IconNameDocumentPrint", theme.IconNameDocumentPrint},
		{"IconNameDocumentSave", theme.IconNameDocumentSave},
		{"IconNameDownload", theme.IconNameDownload},
		{"IconNameDragCornerIndicator", theme.IconNameDragCornerIndicator},
		{"IconNameError", theme.IconNameError},
		{"IconNameFile", theme.IconNameFile},
		{"IconNameFileApplication", theme.IconNameFileApplication},
		{"IconNameFileAudio", theme.IconNameFileAudio},
		{"IconNameFileImage", theme.IconNameFileImage},
		{"IconNameFileText", theme.IconNameFileText},
		{"IconNameFileVideo", theme.IconNameFileVideo},
		{"IconNameFolder", theme.IconNameFolder},
		{"IconNameFolderNew", theme.IconNameFolderNew},
		{"IconNameFolderOpen", theme.IconNameFolderOpen},
		{"IconNameGrid", theme.IconNameGrid},
		{"IconNameHelp", theme.IconNameHelp},
		{"IconNameHistory", theme.IconNameHistory},
		{"IconNameHome", theme.IconNameHome},
		{"IconNameInfo", theme.IconNameInfo},
		{"IconNameList", theme.IconNameList},
		{"IconNameLogin", theme.IconNameLogin},
		{"IconNameLogout", theme.IconNameLogout},
		{"IconNameMailAttachment", theme.IconNameMailAttachment},
		{"IconNameMailCompose", theme.IconNameMailCompose},
		{"IconNameMailForward", theme.IconNameMailForward},
		{"IconNameMailReply", theme.IconNameMailReply},
		{"IconNameMailReplyAll", theme.IconNameMailReplyAll},
		{"IconNameMailSend", theme.IconNameMailSend},
		{"IconNameMediaFastForward", theme.IconNameMediaFastForward},
		{"IconNameMediaFastRewind", theme.IconNameMediaFastRewind},
		{"IconNameMediaMusic", theme.IconNameMediaMusic},
		{"IconNameMediaPause", theme.IconNameMediaPause},
		{"IconNameMediaPhoto", theme.IconNameMediaPhoto},
		{"IconNameMediaPlay", theme.IconNameMediaPlay},
		{"IconNameMediaRecord", theme.IconNameMediaRecord},
		{"IconNameMediaReplay", theme.IconNameMediaReplay},
		{"IconNameMediaSkipNext", theme.IconNameMediaSkipNext},
		{"IconNameMediaSkipPrevious", theme.IconNameMediaSkipPrevious},
		{"IconNameMediaStop", theme.IconNameMediaStop},
		{"IconNameMediaVideo", theme.IconNameMediaVideo},
		{"IconNameMenu", theme.IconNameMenu},
		{"IconNameMenuExpand", theme.IconNameMenuExpand},
		{"IconNameMoreHorizontal", theme.IconNameMoreHorizontal},
		{"IconNameMoreVertical", theme.IconNameMoreVertical},
		{"IconNameMoveDown", theme.IconNameMoveDown},
		{"IconNameMoveUp", theme.IconNameMoveUp},
		{"IconNameNavigateBack", theme.IconNameNavigateBack},
		{"IconNameNavigateNext", theme.IconNameNavigateNext},
		{"IconNameQuestion", theme.IconNameQuestion},
		{"IconNameRadioButton", theme.IconNameRadioButton},
		{"IconNameRadioButtonChecked", theme.IconNameRadioButtonChecked},
		{"IconNameRadioButtonFill", theme.IconNameRadioButtonFill},
		{"IconNameSearch", theme.IconNameSearch},
		{"IconNameSearchReplace", theme.IconNameSearchReplace},
		{"IconNameSettings", theme.IconNameSettings},
		{"IconNameStorage", theme.IconNameStorage},
		{"IconNameUpload", theme.IconNameUpload},
		{"IconNameViewFullScreen", theme.IconNameViewFullScreen},
		{"IconNameViewRefresh", theme.IconNameViewRefresh},
		{"IconNameViewRestore", theme.IconNameViewRestore},
		{"IconNameViewZoomFit", theme.IconNameViewZoomFit},
		{"IconNameViewZoomIn", theme.IconNameViewZoomIn},
		{"IconNameViewZoomOut", theme.IconNameViewZoomOut},
		{"IconNameVisibility", theme.IconNameVisibility},
		{"IconNameVisibilityOff", theme.IconNameVisibilityOff},
		{"IconNameVolumeDown", theme.IconNameVolumeDown},
		{"IconNameVolumeMute", theme.IconNameVolumeMute},
		{"IconNameVolumeUp", theme.IconNameVolumeUp},
		{"IconNameWarning", theme.IconNameWarning},
		{"IconNameWindowClose", theme.IconNameWindowClose},
		{"IconNameWindowMaximize", theme.IconNameWindowMaximize},
		{"IconNameWindowMinimize", theme.IconNameWindowMinimize},
	}
	slices.SortFunc(sizes, func(a, b iconRow) int {
		return strings.Compare(a.label, b.label)
	})
	iconsFiltered := slices.Clone(sizes)
	var iconSize float32 = iconSizeStart
	iconColors := []string{"Default", "Disabled", "Error", "Primary", "Success", "Warning"}
	var iconColor = "Default"
	grid := widget.NewGridWrap(
		func() int {
			return len(iconsFiltered)
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
			if id >= len(iconsFiltered) {
				return
			}
			s := iconsFiltered[id]
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
	search := widget.NewEntry()
	search.SetPlaceHolder("Search...")
	search.OnChanged = func(s string) {
		iconsFiltered = make([]iconRow, 0)
		s2 := strings.ToLower(s)
		for _, c := range sizes {
			if strings.Contains(strings.ToLower(c.label), s2) {
				iconsFiltered = append(iconsFiltered, c)
			}
		}
		grid.Refresh()
	}
	slider := kxwidget.NewSlider(8, 128)
	slider.SetStep(4)
	slider.OnChangeEnded = func(v float64) {
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
		search,
		container.NewBorder(nil, nil, nil, themeBox, sliderBox),
		nil,
		nil,
		grid,
	)
}