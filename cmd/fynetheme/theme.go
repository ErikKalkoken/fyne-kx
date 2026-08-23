package main

import "fyne.io/fyne/v2"

//go:generate go run ../../tools/gentheme/ -out theme_gen.go

type colorRow struct {
	label string
	name  fyne.ThemeColorName
}

type iconRow struct {
	label string
	name  fyne.ThemeIconName
}

type sizeRow struct {
	label string
	name  fyne.ThemeSizeName
}
