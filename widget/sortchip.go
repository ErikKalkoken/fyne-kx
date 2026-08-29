package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// SortOrder represents a sort order.
type SortOrder int8

const (
	sortOrderUnset SortOrder = iota
	SortOrderAscending
	SortOrderDescending
)

func (so SortOrder) String() string {
	switch so {
	case SortOrderAscending:
		return "Ascending"
	case SortOrderDescending:
		return "Descending"
	}
	return "Unknown"
}

// SortChip is a chip widget that shows current sorting (column & order)
// and allows the user to change it by selecting from a drop down menu.
//
// The chip is in off mode when default sorting is selected,
// and in on mode when a different sorting is selected.
type SortChip struct {
	chip

	// The default sort column. Will be used for reset.
	DefaultColumn string

	// The default sort order. Will be used for reset.
	DefaultOrder SortOrder

	// OnChanged is called when a different sorting was selected.
	OnChanged func(column string, order SortOrder)

	// Currently selected sort column.
	Column string

	// Currently selected sort order.
	Order SortOrder

	columns       []string
	columnsLookup map[string]struct{}
}

// NewSortChip returns a new [SortChip] object.
//
// columns defines which columns can be sorted.
// Empty, invalid and duplicate columns will be removed. The order is preserved.
//
// defaultColumn, defaultOrder sets the initial sort column and sort order.
func NewSortChip(columns []string, defaultColumn string, defaultOrder SortOrder, changed func(col string, order SortOrder)) *SortChip {
	w := &SortChip{
		OnChanged:     changed,
		columns:       sliceUniqueNonEmpty(columns),
		columnsLookup: make(map[string]struct{}),
		DefaultColumn: defaultColumn,
		DefaultOrder:  defaultOrder,
		Column:        defaultColumn,
		Order:         defaultOrder,
	}
	w.ExtendBaseWidget(w)
	w.onTapped = w.showMenu

	for _, v := range w.columns {
		w.columnsLookup[v] = struct{}{}
	}
	if len(columns) == 0 {
		fyne.LogError("SortChip misconfigured: No columns.", nil)
	}
	return w
}

func (w *SortChip) Refresh() {
	w.updateState()
	w.chip.Refresh()
}

// ResetSilent resets the sorting to default without calling OnChanged.
func (w *SortChip) ResetSilent() {
	w.Column = w.effectiveDefaultColumn()
	w.Order = w.effectiveDefaultOrder()
	w.Refresh()
}

func (w *SortChip) showMenu() {
	if len(w.columns) == 0 {
		return
	}

	effectiveDefaultColumn := w.effectiveDefaultColumn()
	effectiveDefaultOrder := w.effectiveDefaultOrder()
	currentColumn := w.effectiveColumn()
	currentOrder := w.effectiveOrder()

	onChanged := func(column string, order SortOrder) {
		w.Column = column
		w.Order = order
		if currentColumn == w.Column && currentOrder == w.Order {
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
			onChanged(c, currentOrder)
		})
		if c == currentColumn {
			it.Icon = theme.ConfirmIcon()
		} else {
			it.Icon = iconBlankSvg
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
		if o == currentOrder {
			it.Icon = theme.ConfirmIcon()
		} else {
			it.Icon = iconBlankSvg
		}
		items = append(items, it)
	}

	items = append(items, fyne.NewMenuItemSeparator())
	reset := fyne.NewMenuItem("Reset", func() {
		onChanged(effectiveDefaultColumn, effectiveDefaultOrder)
	})
	reset.Icon = theme.NewThemedResource(iconRestoreSvg)
	reset.Disabled = w.isAtDefault()
	items = append(items, reset)

	showPopUpMenuBelowLeading(w, fyne.NewMenu("", items...))
}

func (w *SortChip) CreateRenderer() fyne.WidgetRenderer {
	w.updateState()
	return w.chip.CreateRenderer()
}

func (w *SortChip) updateState() {
	if len(w.columns) == 0 {
		w.text = "(empty)"
		w.leadingIcon = theme.NewThemedResource(iconSortSvg)
		return
	}

	if x := w.effectiveColumn(); x == "" {
		w.text = "(unset)"
	} else {
		w.text = x
	}

	switch w.effectiveOrder() {
	case SortOrderAscending:
		w.leadingIcon = theme.NewThemedResource(iconSortAscendingSvg)
	case SortOrderDescending:
		w.leadingIcon = theme.NewThemedResource(iconSortDescendingSvg)
	default:
		w.leadingIcon = theme.NewThemedResource(iconSortSvg)
	}

	w.on = !w.isAtDefault()
}

func (w *SortChip) isAtDefault() bool {
	return w.effectiveColumn() == w.effectiveDefaultColumn() && w.effectiveOrder() == w.effectiveDefaultOrder()
}

func (w *SortChip) effectiveColumn() string {
	if _, found := w.columnsLookup[w.Column]; !found {
		return w.effectiveDefaultColumn()
	}
	return w.Column
}

func (w *SortChip) effectiveDefaultColumn() string {
	if _, found := w.columnsLookup[w.DefaultColumn]; !found {
		return ""
	}
	return w.DefaultColumn
}

func (w *SortChip) effectiveOrder() SortOrder {
	if w.Order != SortOrderAscending && w.Order != SortOrderDescending {
		return w.effectiveDefaultOrder()
	}
	return w.Order
}

func (w *SortChip) effectiveDefaultOrder() SortOrder {
	if w.DefaultOrder != SortOrderAscending && w.DefaultOrder != SortOrderDescending {
		return sortOrderUnset
	}
	return w.DefaultOrder
}
