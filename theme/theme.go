// Package theme contains custome themes for the Fyne GUI toolkit.
package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fixedVariant struct {
	variant fyne.ThemeVariant
}

func (ct fixedVariant) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(c, ct.variant)
}

func (fixedVariant) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (fixedVariant) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (fixedVariant) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

// DefaultWithFixedVariant returns the default Fyne theme, but with a fixed theme variant.
// This allows apps to choose a light or dark mode independant of the current os settings.
//
// For example here is how to set an app to permament dark variant:
//
//	app.Settings().SetTheme(kxtheme.DefaultWithFixedVariant(theme.VariantDark))
func DefaultWithFixedVariant(v fyne.ThemeVariant) fyne.Theme {
	return &fixedVariant{variant: v}
}
