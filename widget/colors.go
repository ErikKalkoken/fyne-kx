package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type modifiedColorMode uint

const (
	modeBrighter modifiedColorMode = iota + 1
	modeDarker
)

type modifiedColor struct {
	c    color.Color
	t    float32
	mode modifiedColorMode
}

// newModifiedColor returns a modified instance of a color.
//
// The factor value is expected between 0 and 1. Larger and smaller numbers will be truncated.
func newModifiedColor(c color.Color, mode modifiedColorMode, factor float32) color.Color {
	if factor < 0 {
		factor = 0
	}
	if factor > 1 {
		factor = 1
	}
	return modifiedColor{c: c, t: factor, mode: mode}
}

func (mc modifiedColor) RGBA() (r, g, b, a uint32) {
	r, g, b, a = mc.c.RGBA()
	var r2, g2, b2, f float32
	f = 1 + mc.t
	switch mc.mode {
	case modeBrighter:
		r2 = float32(r) / 0xffff * f
		g2 = float32(g) / 0xffff * f
		b2 = float32(b) / 0xffff * f
		r2, g2, b2 = redistributeRGB(r2, g2, b2)
	case modeDarker:
		r2 = float32(r) / 0xffff / f
		g2 = float32(g) / 0xffff / f
		b2 = float32(b) / 0xffff / f
	}
	r = uint32(r2 * 0xffff)
	g = uint32(g2 * 0xffff)
	b = uint32(b2 * 0xffff)
	return
}

func redistributeRGB(r, g, b float32) (float32, float32, float32) {
	var threshold float32 = 1.0
	m := max(r, g, b)
	if m <= threshold {
		return r, g, b
	}
	total := r + g + b
	if total >= 3*threshold {
		return threshold, threshold, threshold
	}
	x := (3*threshold - total) / (3*m - total)
	gray := threshold - x*m
	return gray + x*r, gray + x*g, gray + x*b
}

func max(values ...float32) float32 {
	var m float32
	for _, v := range values {
		if m < v {
			m = v
		}
	}
	return m
}

// buttonForegroundColor returns the same foreground color name a Fyne
// button uses for its label and icon at the given importance, so other
// widgets can match a button's text color for a given importance.
func buttonForegroundColor(importance widget.Importance) fyne.ThemeColorName {
	switch importance {
	case widget.DangerImportance:
		return theme.ColorNameForegroundOnError
	case widget.HighImportance:
		return theme.ColorNameForegroundOnPrimary
	case widget.SuccessImportance:
		return theme.ColorNameForegroundOnSuccess
	case widget.WarningImportance:
		return theme.ColorNameForegroundOnWarning
	default: // MediumImportance, LowImportance
		return theme.ColorNameForeground
	}
}

// toNRGBA converts a color to RGBA values which are not premultiplied, unlike color.RGBA().
func toNRGBA(c color.Color) (r, g, b, a int) {
	// We use unmultiplyAlpha with RGBA, RGBA64, and unrecognized implementations of Color.
	// It works for all Colors whose RGBA() method is implemented according to spec, but is only necessary for those.
	// Only RGBA and RGBA64 have components which are already premultiplied.
	switch col := c.(type) {
	// NRGBA and NRGBA64 are not premultiplied
	case color.NRGBA:
		r = int(col.R)
		g = int(col.G)
		b = int(col.B)
		a = int(col.A)
	case *color.NRGBA:
		r = int(col.R)
		g = int(col.G)
		b = int(col.B)
		a = int(col.A)
	case color.NRGBA64:
		r = int(col.R) >> 8
		g = int(col.G) >> 8
		b = int(col.B) >> 8
		a = int(col.A) >> 8
	case *color.NRGBA64:
		r = int(col.R) >> 8
		g = int(col.G) >> 8
		b = int(col.B) >> 8
		a = int(col.A) >> 8
	// Gray and Gray16 have no alpha component
	case *color.Gray:
		r = int(col.Y)
		g = int(col.Y)
		b = int(col.Y)
		a = 0xff
	case color.Gray:
		r = int(col.Y)
		g = int(col.Y)
		b = int(col.Y)
		a = 0xff
	case *color.Gray16:
		r = int(col.Y) >> 8
		g = int(col.Y) >> 8
		b = int(col.Y) >> 8
		a = 0xff
	case color.Gray16:
		r = int(col.Y) >> 8
		g = int(col.Y) >> 8
		b = int(col.Y) >> 8
		a = 0xff
	// Alpha and Alpha16 contain only an alpha component.
	case color.Alpha:
		r = 0xff
		g = 0xff
		b = 0xff
		a = int(col.A)
	case *color.Alpha:
		r = 0xff
		g = 0xff
		b = 0xff
		a = int(col.A)
	case color.Alpha16:
		r = 0xff
		g = 0xff
		b = 0xff
		a = int(col.A) >> 8
	case *color.Alpha16:
		r = 0xff
		g = 0xff
		b = 0xff
		a = int(col.A) >> 8
	default: // RGBA, RGBA64, and unknown implementations of Color
		r, g, b, a = unmultiplyAlpha(c)
	}
	return r, g, b, a
}

// unmultiplyAlpha returns a color's RGBA components as 8-bit integers by calling c.RGBA() and then removing the alpha premultiplication.
// It is only used by toNRGBA.
func unmultiplyAlpha(c color.Color) (r, g, b, a int) {
	red, green, blue, alpha := c.RGBA()
	if alpha != 0 && alpha != 0xffff {
		red = (red * 0xffff) / alpha
		green = (green * 0xffff) / alpha
		blue = (blue * 0xffff) / alpha
	}
	// Convert from range 0-65535 to range 0-255
	r = int(red >> 8)
	g = int(green >> 8)
	b = int(blue >> 8)
	a = int(alpha >> 8)
	return r, g, b, a
}
