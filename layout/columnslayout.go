package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type columnsLayout struct {
	widths []float32
}

// NewColumns returns a new columns layout.
//
// Columns arranges all objects in a row, with each in their own column with a given minimum width.
// It can be used to arrange subsequent rows of objects in columns.
//
// The layout will fill the available space. This means that the trailing column might be wider,
// when the parent container has more space available. But it can never shrink below the given width.
// The last width will be re-used for additional columns if needed.
func NewColumns(widths ...float32) fyne.Layout {
	if len(widths) == 0 {
		panic("Need to define at least one width")
	}
	l := columnsLayout{
		widths: widths,
	}
	return l
}

func (l columnsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	wTotal, hTotal := float32(0), float32(0)
	if len(objects) > 0 {
		hTotal = objects[0].MinSize().Height
	}
	var w float32
	for i := range objects {
		if i < len(l.widths) {
			w = l.widths[i]
		}
		wTotal += w
		if i < len(l.widths) {
			wTotal += theme.Padding()
		}
	}
	return fyne.NewSize(wTotal, hTotal)
}

func (l columnsLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var lastX float32
	pos := fyne.NewPos(0, 0)
	padding := theme.Padding()
	var w1 float32
	for i, o := range objects {
		size := o.MinSize()
		if i < len(l.widths) {
			w1 = l.widths[i]
		}
		var w2 float32
		if i < len(objects)-1 || containerSize.Width < 0 {
			w2 = w1
		} else {
			w2 = max(containerSize.Width-pos.X-padding, w1)
		}
		o.Resize(fyne.Size{Width: w2, Height: size.Height})
		o.Move(pos)
		var x float32
		if len(l.widths) > i {
			x = w2
			lastX = x
		} else {
			x = lastX
		}
		pos = pos.AddXY(x+padding, 0)
	}
}
