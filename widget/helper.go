package widget

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// sliceDeduplicate returns a new slice where all duplicate elements have been removed.
// The order of the elements is not changed, but the new slice can be shorter.
func sliceDeduplicate[S ~[]E, E comparable](s S) []E {
	seen := make(map[E]bool, len(s))
	s2 := make([]E, 0, len(s))
	for _, v := range s {
		if seen[v] {
			continue
		}
		s2 = append(s2, v)
		seen[v] = true
	}
	return slices.Clip(s2)
}

func showPopUpMenuBelowLeading(w fyne.CanvasObject, m *fyne.Menu) {
	if m == nil {
		return
	}
	pos := fyne.NewPos(0, w.Size().Height)
	widget.ShowPopUpMenuAtRelativePosition(
		m,
		fyne.CurrentApp().Driver().CanvasForObject(w),
		pos,
		w,
	)
}

// slicesUniqueNonEmpty returns a new slice where duplicate and zero elements have been removed.
func sliceUniqueNonEmpty[S ~[]E, E comparable](s S) []E {
	var z E
	s = slices.DeleteFunc(s, func(s E) bool {
		return s == z
	})
	return sliceDeduplicate(s)
}
