package widget

import (
	"maps"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// FilterChipGroup allows the user to toggle multiple filters with filter chips.
//
// The chips are automatically using the available horizontal space
// and wrapping into multiple rows as needed.
type FilterChipGroup struct {
	widget.BaseWidget

	OnChanged func(selected []string)

	Options  []string // readonly TODO: Enable setting options
	Selected []string // readonly after first render

	chips      []*FilterChip
	options    []string
	selected   []string
	isDisabled bool
}

var _ fyne.Disableable = (*FilterChipGroup)(nil)

// NewFilterChipGroup returns a new [FilterChipGroup].
func NewFilterChipGroup(options []string, changed func([]string)) *FilterChipGroup {
	optionsCleaned := sliceUniqueNonEmpty(options)
	w := &FilterChipGroup{
		chips:     make([]*FilterChip, 0),
		OnChanged: changed,
		options:   optionsCleaned,
		Options:   slices.Clone(optionsCleaned),
		Selected:  make([]string, 0),
	}
	w.ExtendBaseWidget(w)
	for _, v := range w.options {
		w.chips = append(w.chips, NewFilterChip(v, func(on bool) {
			isSelected := make(map[string]bool)
			for _, x := range w.selected {
				isSelected[x] = true
			}
			isSelected[v] = on
			w.updateSelected(isSelected)
			w.Refresh()
			if w.OnChanged != nil {
				w.OnChanged(w.Selected)
			}
		}))
	}
	return w
}

func (w *FilterChipGroup) Enable() {
	if !w.isDisabled {
		return
	}
	for _, c := range w.chips {
		c.Enable()
	}
	w.isDisabled = false
	w.Refresh()
}

func (w *FilterChipGroup) Disable() {
	if w.isDisabled {
		return
	}
	for _, c := range w.chips {
		c.Disable()
	}
	w.isDisabled = true
	w.Refresh()
}

func (w *FilterChipGroup) Disabled() bool {
	return w.isDisabled
}

// SetSelected updates the selected options.
// Invalid elements including empty strings will be ignored.
func (w *FilterChipGroup) SetSelected(s []string) {
	s2 := slices.DeleteFunc(slices.Clone(s), func(v string) bool {
		return v == ""
	})
	if !w.setSelected(s2) {
		return
	}
	w.Refresh()
	if w.OnChanged != nil {
		w.OnChanged(w.Selected)
	}
}

func (w *FilterChipGroup) CreateRenderer() fyne.WidgetRenderer {
	w.setSelected(w.Selected)
	p := w.Theme().Size(theme.SizeNamePadding)
	box := container.New(layout.NewRowWrapLayoutWithCustomPadding(p, p))
	for _, c := range w.chips {
		box.Add(c)
	}
	return widget.NewSimpleRenderer(box)
}

func (w *FilterChipGroup) setSelected(s []string) bool {
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
	currentSelected := make(map[string]bool)
	for _, v := range w.selected {
		currentSelected[v] = true
	}
	if maps.Equal(isSelected, currentSelected) {
		return false
	}
	for i, v := range w.options {
		w.chips[i].On = isSelected[v]
		w.chips[i].Refresh()
	}
	w.updateSelected(isSelected)
	return true
}

func (w *FilterChipGroup) updateSelected(isSelected map[string]bool) {
	w.selected = make([]string, 0)
	for _, x := range w.options {
		if isSelected[x] {
			w.selected = append(w.selected, x)
		}
	}
	w.Selected = slices.Clone(w.selected)
}
