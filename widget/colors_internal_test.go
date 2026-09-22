package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNRGBA(t *testing.T) {
	cases := []struct {
		name       string
		c          color.Color
		r, g, b, a int
	}{
		{"NRGBA value", color.NRGBA{R: 10, G: 20, B: 30, A: 40}, 10, 20, 30, 40},
		{"NRGBA pointer", &color.NRGBA{R: 10, G: 20, B: 30, A: 40}, 10, 20, 30, 40},
		{"NRGBA64 value", color.NRGBA64{R: 0x1020, G: 0x3040, B: 0x5060, A: 0x7080}, 0x10, 0x30, 0x50, 0x70},
		{"NRGBA64 pointer", &color.NRGBA64{R: 0x1020, G: 0x3040, B: 0x5060, A: 0x7080}, 0x10, 0x30, 0x50, 0x70},
		{"Gray value", color.Gray{Y: 100}, 100, 100, 100, 0xff},
		{"Gray pointer", &color.Gray{Y: 100}, 100, 100, 100, 0xff},
		{"Gray16 value", color.Gray16{Y: 0x6478}, 0x64, 0x64, 0x64, 0xff},
		{"Gray16 pointer", &color.Gray16{Y: 0x6478}, 0x64, 0x64, 0x64, 0xff},
		{"Alpha value", color.Alpha{A: 128}, 0xff, 0xff, 0xff, 128},
		{"Alpha pointer", &color.Alpha{A: 128}, 0xff, 0xff, 0xff, 128},
		{"Alpha16 value", color.Alpha16{A: 0x8090}, 0xff, 0xff, 0xff, 0x80},
		{"Alpha16 pointer", &color.Alpha16{A: 0x8090}, 0xff, 0xff, 0xff, 0x80},
		// RGBA64 is not special-cased, so it exercises the default branch,
		// which unmultiplies alpha via unmultiplyAlpha.
		{"default branch, fully opaque", color.RGBA64{R: 0xABAB, G: 0x1234, B: 0x5678, A: 0xffff}, 0xAB, 0x12, 0x56, 0xff},
		{"default branch, fully transparent", color.RGBA64{R: 0, G: 0, B: 0, A: 0}, 0, 0, 0, 0},
		{"default branch, partial alpha", color.RGBA64{R: 0x8000, G: 0x8000, B: 0x8000, A: 0x8000}, 0xff, 0xff, 0xff, 0x80},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r, g, b, a := toNRGBA(tc.c)
			assert.Equal(t, tc.r, r, "r")
			assert.Equal(t, tc.g, g, "g")
			assert.Equal(t, tc.b, b, "b")
			assert.Equal(t, tc.a, a, "a")
		})
	}
}

func TestUnmultiplyAlpha(t *testing.T) {
	cases := []struct {
		name       string
		c          color.Color
		r, g, b, a int
	}{
		{"fully opaque: no unmultiplication needed", color.RGBA64{R: 0xABAB, G: 0x1234, B: 0x5678, A: 0xffff}, 0xAB, 0x12, 0x56, 0xff},
		{"fully transparent: division by zero avoided", color.RGBA64{R: 0, G: 0, B: 0, A: 0}, 0, 0, 0, 0},
		{"partial alpha: components unmultiplied", color.RGBA64{R: 0x8000, G: 0x8000, B: 0x8000, A: 0x8000}, 0xff, 0xff, 0xff, 0x80},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r, g, b, a := unmultiplyAlpha(tc.c)
			assert.Equal(t, tc.r, r, "r")
			assert.Equal(t, tc.g, g, "g")
			assert.Equal(t, tc.b, b, "b")
			assert.Equal(t, tc.a, a, "a")
		})
	}
}
