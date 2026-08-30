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

// widthFor returns the configured width for column i, reusing the last
// defined width for any columns beyond len(l.widths).
func (l columnsLayout) widthFor(i int) float32 {
	if i < len(l.widths) {
		return l.widths[i]
	}
	return l.widths[len(l.widths)-1]
}

func (l columnsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	wTotal, hTotal := float32(0), float32(0)
	for _, o := range objects {
		hTotal = fyne.Max(hTotal, o.MinSize().Height)
	}
	for i := range objects {
		wTotal += l.widthFor(i)
		if i < len(objects)-1 {
			wTotal += theme.Padding()
		}
	}
	return fyne.NewSize(wTotal, hTotal)
}

func (l columnsLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	padding := theme.Padding()
	for i, o := range objects {
		size := o.MinSize()
		w1 := l.widthFor(i)
		var w2 float32
		if i < len(objects)-1 || containerSize.Width < 0 {
			w2 = w1
		} else {
			w2 = fyne.Max(containerSize.Width-pos.X-padding, w1)
		}
		o.Resize(fyne.Size{Width: w2, Height: size.Height})
		o.Move(pos)
		pos = pos.AddXY(w2+padding, 0)
	}
}
