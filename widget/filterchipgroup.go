package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	kxlayout "github.com/ErikKalkoken/fyne-kx/layout"
)

// FilterChipGroup allows the user to toggle multiple filters with filter chips.
//
// The chips are automatically using the available horizontal space
// and wrapping into multiple rows as needed.
type FilterChipGroup struct {
	widget.DisableableWidget

	OnChanged func(selected []string)

	Options  []string // readonly TODO: Enable setting options
	Selected []string // readonly after first render

	chips    []*FilterChip
	options  []string
	selected []string
}

// NewFilterChipGroup returns a new [FilterChipGroup].
func NewFilterChipGroup(options []string, changed func([]string)) *FilterChipGroup {
	optionsCleaned := sliceDeleteFunc(sliceDeduplicate(options), func(v string) bool {
		return v == ""
	})
	w := &FilterChipGroup{
		chips:     make([]*FilterChip, 0),
		OnChanged: changed,
		options:   optionsCleaned,
		Options:   sliceClone(optionsCleaned),
		Selected:  make([]string, 0),
	}
	w.ExtendBaseWidget(w)
	for _, v := range w.options {
		v := v
		w.chips = append(w.chips, NewFilterChip(v, func(on bool) {
			isSelected := make(map[string]bool)
			for _, x := range w.selected {
				isSelected[x] = true
			}
			if on {
				isSelected[v] = true
			} else {
				isSelected[v] = false
			}
			w.updateSelected(isSelected)
			if w.OnChanged != nil {
				w.OnChanged(w.Selected)
			}
		}))
	}
	return w
}

func (w *FilterChipGroup) updateSelected(isSelected map[string]bool) {
	w.selected = make([]string, 0)
	for _, x := range w.options {
		if isSelected[x] {
			w.selected = append(w.selected, x)
		}
	}
	w.Selected = sliceClone(w.selected)
}

// SetSelected updates the selected options.
// Invalid elements including empty strings will be ignored.
func (w *FilterChipGroup) SetSelected(s []string) {
	w.setSelected(s)
	w.Refresh()
	if w.OnChanged != nil {
		w.OnChanged(w.Selected)
	}
}

func (w *FilterChipGroup) setSelected(s []string) {
	isValid := make(map[string]bool)
	for _, v := range w.options {
		isValid[v] = true
	}
	isSelected := make(map[string]bool)
	for _, v := range s {
		if !isValid[v] {
			continue
		}
		isSelected[v] = true
	}
	for i, v := range w.options {
		w.chips[i].On = isSelected[v]
		w.chips[i].Refresh()
	}
	w.updateSelected(isSelected)
}

func (w *FilterChipGroup) CreateRenderer() fyne.WidgetRenderer {
	w.setSelected(w.Selected)
	p := w.Theme().Size(theme.SizeNamePadding)
	box := container.New(kxlayout.NewRowWrapLayoutWithCustomPadding(2*p, 2*p))
	for _, c := range w.chips {
		box.Add(c)
	}
	return widget.NewSimpleRenderer(container.New(layout.NewCustomPaddedLayout(p, p, p, p), box))
}
