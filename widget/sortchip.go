package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// SortOrder represents the order for sorting.
type SortOrder int8

const (
	SortOrderNone SortOrder = iota
	SortOrderAscending
	SortOrderDescending
)

func (so SortOrder) String() string {
	switch so {
	case SortOrderAscending:
		return "Ascending"
	case SortOrderDescending:
		return "Descending"
	default:
		return "None"
	}
}

// SortChip is a chip for sorting.
//
// It shows the currently selected column and sort order.
// When clicked it will show the sorting options in a drop down menu.
//
// The chip is in off state when default sorting is selected.
// Otherwise it is in on state.
type SortChip struct {
	chip

	// The default sort column.
	// When not set or invalid will be set to first column on next refresh.
	DefaultColumn string

	// The default ordering.
	// When not set will be set to ascending on next refresh.
	DefaultOrder SortOrder

	// OnChanged is called when a different sorting was selected.
	OnChanged func(column string, order SortOrder)

	// Currently selected column for sorting.
	// When not set or invalid will be set to default column on next refresh.
	Column string

	// Currently selected sort order.
	// When not set will be set to ascending on next refresh.
	Order SortOrder

	columns       []string
	columnsLookup map[string]struct{}
}

// NewSortChip returns a new [SortChip] object.
//
// columns defines which columns can be sorted.
// Empty, invalid and duplicate columns will be removed. The order is preserved.
func NewSortChip(columns []string, changed func(col string, order SortOrder)) *SortChip {
	w := &SortChip{
		OnChanged:     changed,
		columns:       sliceUniqueNonEmpty(columns),
		columnsLookup: make(map[string]struct{}),
	}
	w.ExtendBaseWidget(w)
	w.onTapped = w.showMenu

	for _, v := range w.columns {
		w.columnsLookup[v] = struct{}{}
	}
	return w
}

func (w *SortChip) Refresh() {
	w.updateState()
	w.chip.Refresh()
}

// ResetSilent resets the sorting to default without calling OnChanged.
func (w *SortChip) ResetSilent() {
	w.resetInvalidDefaults()
	w.Column = w.DefaultColumn
	w.Order = w.DefaultOrder
	w.Refresh()
}

func (w *SortChip) showMenu() {
	if len(w.columns) == 0 {
		return
	}

	oldColumn := w.Column
	oldDirection := w.Order

	onChanged := func(column string, order SortOrder) {
		w.Column = column
		w.Order = order
		if oldColumn == w.Column && oldDirection == w.Order {
			return
		}
		w.Refresh()
		if w.OnChanged != nil {
			w.OnChanged(w.Column, w.Order)
		}
	}

	var items []*fyne.MenuItem

	sortTitle := fyne.NewMenuItem("Sort by ", nil)
	sortTitle.Disabled = true
	items = append(items, sortTitle)

	for _, c := range w.columns {
		it := fyne.NewMenuItem(c, func() {
			onChanged(c, w.Order)
		})
		if c == w.Column {
			it.Icon = theme.ConfirmIcon()
		} else {
			it.Icon = IconBlankSvg
		}
		items = append(items, it)
	}

	orderTitle := fyne.NewMenuItem("Order", nil)
	orderTitle.Disabled = true
	items = append(items, orderTitle)

	for _, o := range []SortOrder{SortOrderAscending, SortOrderDescending} {
		it := fyne.NewMenuItem(o.String(), func() {
			onChanged(w.Column, o)
		})
		if o == w.Order {
			it.Icon = theme.ConfirmIcon()
		} else {
			it.Icon = IconBlankSvg
		}
		items = append(items, it)
	}

	items = append(items, fyne.NewMenuItemSeparator())
	reset := fyne.NewMenuItem("Reset", func() {
		onChanged(w.DefaultColumn, w.DefaultOrder)
	})
	reset.Icon = theme.NewThemedResource(IconRestoreSvg)
	reset.Disabled = w.Column == w.DefaultColumn && w.Order == w.DefaultOrder
	items = append(items, reset)

	menu := fyne.NewMenu("", items...)
	showPopUpMenuBelowLeading(w, menu)
}

func (w *SortChip) CreateRenderer() fyne.WidgetRenderer {
	w.updateState()
	return w.chip.CreateRenderer()
}

func (w *SortChip) updateState() {
	if len(w.columns) == 0 {
		w.text = "(empty)"
		w.leadingIcon = theme.NewThemedResource(IconSortSvg)
		return
	}

	w.resetInvalidDefaults()
	if _, found := w.columnsLookup[w.Column]; !found {
		w.Column = w.DefaultColumn
	}
	if w.Order == SortOrderNone {
		w.Order = w.DefaultOrder
	}

	w.text = w.Column
	switch w.Order {
	case SortOrderAscending:
		w.leadingIcon = theme.NewThemedResource(IconSortAscendingSvg)
	case SortOrderDescending:
		w.leadingIcon = theme.NewThemedResource(IconSortDescendingSvg)
	default:
		w.leadingIcon = theme.NewThemedResource(IconSortSvg)
	}

	isDefault := w.Order == w.DefaultOrder && w.Column == w.DefaultColumn
	w.on = !isDefault
}

func (w *SortChip) resetInvalidDefaults() {
	if len(w.columns) == 0 {
		return
	}
	if _, found := w.columnsLookup[w.DefaultColumn]; !found {
		w.DefaultColumn = w.columns[0]
	}
	if w.DefaultOrder == SortOrderNone {
		w.DefaultOrder = SortOrderAscending
	}
}
