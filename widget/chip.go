package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// A chip is a compact, interactive widget that represents a single entity, attribute,
// action, or filter.
type chip struct {
	widget.DisableableWidget

	onTapped     func()
	on           bool
	leadingIcon  fyne.Resource
	trailingIcon fyne.Resource
	text         string

	focused bool
	hovered bool
}

var _ desktop.Hoverable = (*chip)(nil)
var _ fyne.Disableable = (*chip)(nil)
var _ fyne.Focusable = (*chip)(nil)
var _ fyne.Tappable = (*chip)(nil)
var _ fyne.Widget = (*chip)(nil)

// newChip returns a new [chip] object.
func newChip(text string, tapped func()) *chip {
	w := &chip{
		onTapped: tapped,
		text:     text,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *chip) Tapped(pe *fyne.PointEvent) {
	if w.Disabled() {
		return
	}
	if w.onTapped != nil {
		w.onTapped()
	}
}

func (w *chip) Cursor() desktop.Cursor {
	if !w.Disabled() && w.hovered {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

func (w *chip) MouseIn(me *desktop.MouseEvent) {
	if w.Disabled() {
		return
	}
	w.hovered = true
}

func (w *chip) MouseMoved(me *desktop.MouseEvent) {}

func (w *chip) MouseOut() {
	w.hovered = false
}

func (w *chip) FocusGained() {
	if w.Disabled() {
		return
	}
	w.focused = true
	w.Refresh()
}

func (w *chip) FocusLost() {
	w.focused = false
	w.Refresh()
}

func (w *chip) TypedRune(r rune) {
	if w.Disabled() {
		return
	}
	if r == ' ' && w.onTapped != nil {
		w.onTapped()
	}
}

func (w *chip) TypedKey(key *fyne.KeyEvent) {}

func (w *chip) CreateRenderer() fyne.WidgetRenderer {
	r := &chipRenderer{
		chip:         w,
		bg:           canvas.NewRectangle(color.Transparent),
		leadingIcon:  widget.NewIcon(nil),
		trailingIcon: widget.NewIcon(nil),
		label:        canvas.NewText("", color.Transparent),
	}

	r.updateState()
	return r
}

type chipRenderer struct {
	bg           *canvas.Rectangle
	chip         *chip
	label        *canvas.Text
	leadingIcon  *widget.Icon
	trailingIcon *widget.Icon
	lastAutoSize fyne.Size
}

func (r *chipRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.leadingIcon, r.label, r.trailingIcon}
}

func (r *chipRenderer) Destroy() {}

func (r *chipRenderer) Layout(size fyne.Size) {
	minSize := r.MinSize()
	size = fyne.NewSize(max(size.Width, minSize.Width), max(size.Height, minSize.Height))

	r.bg.Resize(size)
	r.bg.Move(fyne.NewPos(0, 0))

	textMin := r.label.MinSize()
	th := r.chip.Theme()
	gapSize := th.Size(theme.SizeNameInnerPadding) / 2

	var leadWidth, leadHeight float32
	var leadGap float32
	if !r.leadingIcon.Hidden {
		leadMin := r.leadingIcon.MinSize()
		leadWidth = leadMin.Width
		leadHeight = leadMin.Height
		if textMin.Width > 0 || !r.trailingIcon.Hidden {
			leadGap = gapSize
		}
	} else {
		r.leadingIcon.Resize(fyne.NewSize(0, 0))
	}

	var trailWidth, trailHeight float32
	var trailGap float32
	if !r.trailingIcon.Hidden {
		trailMin := r.trailingIcon.MinSize()
		trailWidth = trailMin.Width
		trailHeight = trailMin.Height
		if textMin.Width > 0 || !r.leadingIcon.Hidden {
			trailGap = gapSize
		}
	} else {
		r.trailingIcon.Resize(fyne.NewSize(0, 0))
	}

	contentWidth := leadWidth + leadGap + textMin.Width + trailGap + trailWidth
	contentHeight := max(max(leadHeight, textMin.Height), trailHeight)

	startX := (size.Width - contentWidth) / 2
	startY := (size.Height - contentHeight) / 2
	currentX := startX

	// Position Leading Icon
	if !r.leadingIcon.Hidden {
		iconY := startY + (contentHeight-leadHeight)/2
		r.leadingIcon.Resize(fyne.NewSize(leadWidth, leadHeight))
		r.leadingIcon.Move(fyne.NewPos(currentX, iconY))
		currentX += leadWidth + leadGap
	}

	// Position Label
	textY := (size.Height - textMin.Height) / 2
	r.label.Resize(textMin)
	r.label.Move(fyne.NewPos(currentX, textY))
	currentX += textMin.Width + trailGap

	// Position Trailing Icon
	if !r.trailingIcon.Hidden {
		iconY := startY + (contentHeight-trailHeight)/2
		r.trailingIcon.Resize(fyne.NewSize(trailWidth, trailHeight))
		r.trailingIcon.Move(fyne.NewPos(currentX, iconY))
	}
}

func (r *chipRenderer) MinSize() fyne.Size {
	textMin := r.label.MinSize()
	th := r.chip.Theme()
	innerPadding := th.Size(theme.SizeNameInnerPadding)
	gapSize := innerPadding / 2

	var leadWidth, leadHeight float32
	var leadGap float32
	if !r.leadingIcon.Hidden {
		leadMin := r.leadingIcon.MinSize()
		leadWidth = leadMin.Width
		leadHeight = leadMin.Height
		if textMin.Width > 0 || !r.trailingIcon.Hidden {
			leadGap = gapSize
		}
	}

	var trailWidth, trailHeight float32
	var trailGap float32
	if !r.trailingIcon.Hidden {
		trailMin := r.trailingIcon.MinSize()
		trailWidth = trailMin.Width
		trailHeight = trailMin.Height
		if textMin.Width > 0 || !r.leadingIcon.Hidden {
			trailGap = gapSize
		}
	}

	iconSize := th.Size(theme.SizeNameInlineIcon) // use as minimum height reference
	width := leadWidth + leadGap + textMin.Width + trailGap + trailWidth + (innerPadding * 2)
	height := max(leadHeight, textMin.Height, trailHeight, iconSize) + (innerPadding * 2)

	return fyne.NewSize(width, height)
}

func (r *chipRenderer) updateState() {
	w := r.chip
	th := w.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	// Update text properties
	r.label.Text = w.text
	r.label.TextSize = th.Size(theme.SizeNameText)

	// Update icon state
	r.updateIcon(r.leadingIcon, w.leadingIcon)
	r.updateIcon(r.trailingIcon, w.trailingIcon)

	// Background and text color updates
	if w.Disabled() {
		r.bg.StrokeColor = th.Color(theme.ColorNameDisabled, v)
		r.label.Color = th.Color(theme.ColorNameDisabled, v)
	} else {
		r.bg.StrokeColor = th.Color(theme.ColorNameInputBorder, v)
		r.label.Color = th.Color(theme.ColorNameForeground, v)
	}

	if w.on {
		if w.Disabled() {
			r.bg.FillColor = th.Color(theme.ColorNameDisabledButton, v)
			r.bg.StrokeColor = th.Color(theme.ColorNameDisabledButton, v)
		} else {
			r.bg.FillColor = th.Color(theme.ColorNameSelection, v)
			r.bg.StrokeColor = th.Color(theme.ColorNameSelection, v)
		}
	} else {
		r.bg.FillColor = color.Transparent
	}

	if w.focused {
		r.bg.StrokeColor = th.Color(theme.ColorNameFocus, v)
	}

	r.bg.StrokeWidth = th.Size(theme.SizeNameInputBorder)
	r.bg.CornerRadius = th.Size(theme.SizeNameInputRadius)
}

func (r *chipRenderer) updateIcon(icon *widget.Icon, res fyne.Resource) {
	if res == nil {
		icon.Hide()
		return
	}
	icon.Show()
	if r.chip.Disabled() {
		icon.SetResource(theme.NewDisabledResource(res))
	} else {
		icon.SetResource(res)
	}
}

func (r *chipRenderer) Refresh() {
	r.updateState()

	// Resize logic on state refresh
	newMin := r.MinSize()
	current := r.chip.Size()
	if current.IsZero() || current == r.lastAutoSize {
		r.chip.Resize(newMin)
		r.lastAutoSize = newMin
	} else {
		r.Layout(current)
	}

	r.label.Refresh()
	r.leadingIcon.Refresh()
	r.trailingIcon.Refresh()
	r.bg.Refresh()
}
