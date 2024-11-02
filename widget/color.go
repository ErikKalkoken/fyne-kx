package widget

import "image/color"

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
		r2, g2, b2 = redistribute_rgb(r2, g2, b2)
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

func redistribute_rgb(r, g, b float32) (float32, float32, float32) {
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
