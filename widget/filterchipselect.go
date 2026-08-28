package widget

import (
	"slices"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// FilterChipSelect represents a filter chip widget that allows the user to select
// from a list of options.
type FilterChipSelect struct {
	chip

	// The label shown for clearing a selection.
	ClearLabel string

	// OnChanged is called when the selected option changed.
	OnChanged func(selected string)

	// Options is the list of options that can be selected
	// They are deduplicated and sorted alphabetically when shown to the user.
	// Empty option strings will be ignored.
	Options []string

	// The placeholder text is shown as label when nothing is selected.
	//
	// To create a filter which is always selected leave placeholder empty
	// and set an initial option.
	Placeholder string

	// The currently selected option or empty when nothing is selected.
	// This can also be used to set an initial option.
	Selected string

	// Whether to disable sorting of options.
	SortDisabled bool
}

// NewFilterChipSelect returns a new [FilterChipSelect] widget with a drop down menu.
func NewFilterChipSelect(placeholder string, options []string, changed func(selected string)) *FilterChipSelect {
	w := newFilterChipSelect(placeholder, options, changed, nil)
	return w
}

// NewFilterChipSelectWithSearch returns a new [FilterChipSelect] widget with a search dialog.
//
// placeholder is shown as label text when nothing is selected.
func NewFilterChipSelectWithSearch(placeholder string, options []string, changed func(selected string), window fyne.Window) *FilterChipSelect {
	w := newFilterChipSelect(placeholder, options, changed, window)
	return w
}

func newFilterChipSelect(placeholder string, options []string, changed func(selected string), window fyne.Window) *FilterChipSelect {
	w := &FilterChipSelect{
		OnChanged:   changed,
		Placeholder: placeholder,
		Options:     sliceUniqueNonEmpty(options),
	}
	w.ExtendBaseWidget(w)
	w.trailingIcon = theme.MenuDropDownIcon()
	if window == nil {
		// show drop down
		w.onTapped = w.showDropDownMenu
	} else {
		// show search dialog
		w.onTapped = func() {
			w.showSearchDialog(window)
		}
	}
	return w
}

// ClearSelected clears any selection.
func (w *FilterChipSelect) ClearSelected() {
	w.SetSelected("")
}

// SetSelected selects an option.
//
// An empty string will clear the selection.
// Invalid options will be ignored.
func (w *FilterChipSelect) SetSelected(v string) {
	if w.Selected == v {
		return
	}
	if v != "" && !slices.Contains(w.Options, v) {
		return
	}
	if v == "" && w.Placeholder == "" {
		return
	}
	w.Selected = v
	if w.OnChanged != nil {
		w.OnChanged(v)
	}
	w.Refresh()
}

func (w *FilterChipSelect) Refresh() {
	w.updateState()
	w.chip.Refresh()
}

// SetOptions sets the options.
func (w *FilterChipSelect) SetOptions(options []string) {
	w.Options = sliceUniqueNonEmpty(options)
	w.Refresh()
}

func (w *FilterChipSelect) showDropDownMenu() {
	items := make([]*fyne.MenuItem, 0)
	if w.text != "" && w.Selected != "" {
		it := fyne.NewMenuItem(w.ClearLabel, func() {
			w.SetSelected("")
		})
		it.Icon = theme.DeleteIcon()
		items = append(items, it)
		items = append(items, fyne.NewMenuItemSeparator())
	}
	options := sliceUniqueNonEmpty(w.Options)
	if w.Selected != "" && !slices.Contains(options, w.Selected) {
		options = append(options, w.Selected)
	}
	if len(options) == 0 {
		it := fyne.NewMenuItem("No entries", nil)
		it.Disabled = true
		items = append(items, it)
	} else {
		if !w.SortDisabled {
			sort.Slice(options, func(i, j int) bool {
				return strings.ToLower(options[i]) < strings.ToLower(options[j])
			})
		}
		for _, o := range options {
			it := fyne.NewMenuItem(o, func() {
				w.SetSelected(o)
			})
			if w.Selected != "" {
				if o == w.Selected {
					it.Icon = theme.ConfirmIcon()
				} else {
					it.Icon = IconBlankSvg
				}
			}
			items = append(items, it)
		}
	}
	showPopUpMenuBelowLeading(w, fyne.NewMenu("", items...))
}

func (w *FilterChipSelect) showSearchDialog(window fyne.Window) {
	options := sliceUniqueNonEmpty(w.Options)
	baseItems := slices.Clone(options)
	if w.Selected != "" && !slices.Contains(baseItems, w.Selected) {
		baseItems = append(baseItems, w.Selected)
	}
	if !w.SortDisabled {
		sort.Slice(baseItems, func(i, j int) bool {
			return strings.ToLower(baseItems[i]) < strings.ToLower(baseItems[j])
		})
	}

	itemsFiltered := slices.Clone(baseItems)

	var d dialog.Dialog
	list := widget.NewList(
		func() int {
			return len(itemsFiltered)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(IconBlankSvg)
			if w.Selected == "" {
				icon.Hide()
			} else {
				icon.Show()
			}
			return container.NewBorder(
				nil,
				nil,
				icon,
				nil,
				widget.NewLabel(""),
			)
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			if id >= len(itemsFiltered) {
				return
			}
			s := itemsFiltered[id]
			box := co.(*fyne.Container).Objects
			box[0].(*widget.Label).SetText(s)
			if w.Selected == "" {
				return
			}
			icon := box[1].(*widget.Icon)
			if s == w.Selected {
				icon.SetResource(theme.ConfirmIcon())
			} else {
				icon.SetResource(IconBlankSvg)
			}
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		if id >= len(itemsFiltered) {
			return
		}
		w.SetSelected(itemsFiltered[id])
		d.Hide()
	}
	list.HideSeparators = true

	entry := widget.NewEntry()
	entry.PlaceHolder = "Type to start searching..."
	clearButton := NewIconButton(theme.CancelIcon(), func() {
		entry.SetText("")
	})
	clearButton.Hide()
	entry.ActionItem = clearButton
	entry.OnChanged = func(search string) {
		if search != "" {
			clearButton.Show()
		} else {
			clearButton.Hide()
		}
		if len(search) < 2 {
			itemsFiltered = slices.Clone(baseItems)
			list.Refresh()
			return
		}
		itemsFiltered = make([]string, 0)
		search2 := strings.ToLower(search)
		for _, s := range baseItems {
			if strings.Contains(strings.ToLower(s), search2) {
				itemsFiltered = append(itemsFiltered, s)
			}
		}
		list.Refresh()
	}
	clear := widget.NewButton("Clear", func() {
		w.SetSelected("")
		d.Hide()
	})
	if w.Selected != "" {
		entry.Disable()
		clear.Show()
	} else {
		clear.Hide()
	}
	empty := widget.NewLabel("No entries")
	empty.Importance = widget.LowImportance
	if len(options) == 0 {
		empty.Show()
		entry.Disable()
	} else {
		empty.Hide()
	}
	c := container.NewBorder(
		container.NewBorder(
			nil,
			clear,
			nil,
			widget.NewButton("Cancel", func() {
				d.Hide()
			}),
			entry,
		),
		empty,
		nil,
		nil,
		list,
	)
	d = dialog.NewCustomWithoutButtons("Filter by "+w.text, c, window)
	_, s := window.Canvas().InteractiveArea()
	if fyne.CurrentDevice().IsMobile() {
		d.Resize(fyne.NewSize(s.Width, s.Height))
	} else {
		d.Resize(fyne.NewSize(600, max(400, s.Height*0.8)))
	}
	d.Show()
	window.Canvas().Focus(entry)
}

func (w *FilterChipSelect) CreateRenderer() fyne.WidgetRenderer {
	w.updateState()
	return w.chip.CreateRenderer()
}

func (w *FilterChipSelect) updateState() {
	if w.ClearLabel == "" {
		w.ClearLabel = "Clear"
	}
	if w.Selected == "" {
		w.text = w.Placeholder
		w.on = false
		w.leadingIcon = nil
	} else {
		w.text = w.Selected
		w.on = true
		w.leadingIcon = theme.ConfirmIcon()
	}
}
