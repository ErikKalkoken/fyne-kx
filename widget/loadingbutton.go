package widget

import (
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// LoadingButton represents a button widget which shows a loading indicator
// while its action is running.
type LoadingButton struct {
	widget.BaseWidget

	// OnAction is called when the button is tapped and runs on the main
	// goroutine, consistent with other Fyne widget callbacks. The button is
	// locked and shows a loading indicator until the received done function
	// is called; for long-running work, call it from your own goroutine.
	// done is safe to call from any goroutine and more than once.
	OnAction func(done func())

	button        *lockableButton
	disabledTemp  bool
	icon          fyne.Resource
	label         string
	activity      *widget.Activity
	activityTheme *activityColorTheme
	activityWrap  *container.ThemeOverride
	spacer        *canvas.Rectangle
}

var _ fyne.Accessible = (*LoadingButton)(nil)
var _ fyne.Disableable = (*LoadingButton)(nil)
var _ fyne.Widget = (*LoadingButton)(nil)

// NewLoadingButton creates and returns a new [LoadingButton].
func NewLoadingButton(label string, icon fyne.Resource, action func(done func())) *LoadingButton {
	w := &LoadingButton{
		button:   newLockableButton(label, icon, nil),
		activity: widget.NewActivity(),
		spacer:   canvas.NewRectangle(color.Transparent),
		label:    label,
		icon:     icon,
		OnAction: action,
	}
	w.ExtendBaseWidget(w)
	w.activity.Hide()
	w.activity.Stop()
	w.button.OnTapped = func() {
		if w.OnAction == nil {
			return
		}
		w.button.lock()
		// clear button
		w.spacer.SetMinSize(w.button.MinSize())
		w.button.Text = ""
		w.button.Icon = nil
		w.button.Refresh()
		// show progress
		w.activity.Show()
		w.activity.Start()

		var once sync.Once
		done := func() {
			once.Do(func() {
				fyne.Do(w.restore)
			})
		}
		w.OnAction(done)
	}
	return w
}

// restore returns the button to its normal state and hides the progress
// indicator. Must run on the main goroutine.
func (w *LoadingButton) restore() {
	w.button.Text = w.label
	w.button.Icon = w.icon
	if w.disabledTemp {
		w.button.Disable()
		w.disabledTemp = false
	} else {
		w.button.Refresh()
	}
	w.spacer.SetMinSize(fyne.Size{})
	w.activity.Stop()
	w.activity.Hide()
	w.button.unlock()
}

func (w *LoadingButton) CreateRenderer() fyne.WidgetRenderer {
	if w.activityWrap == nil {
		w.activityTheme = &activityColorTheme{
			Theme:     w.Theme(),
			colorName: loadingButtonForegroundColor(w.button.Importance),
		}
		w.activityWrap = container.NewThemeOverride(w.activity, w.activityTheme)
	}
	c := container.NewStack(w.spacer, w.button, w.activityWrap)
	return widget.NewSimpleRenderer(c)
}

// Refresh resyncs the loading indicator's theme override with the current
// theme (e.g. after an app-wide theme change) before delegating to the
// default widget refresh.
func (w *LoadingButton) Refresh() {
	if w.activityTheme != nil {
		w.activityTheme.Theme = w.Theme()
		w.activityWrap.Refresh()
	}
	w.BaseWidget.Refresh()
}

// SetImportance sets the importance of the button.
// Unlike SetText/SetIcon, this applies immediately even while locked.
func (w *LoadingButton) SetImportance(v widget.Importance) {
	w.button.Importance = v
	if w.activityTheme != nil {
		w.activityTheme.colorName = loadingButtonForegroundColor(v)
		w.activityWrap.Refresh()
	}
	w.button.Refresh()
}

// loadingButtonForegroundColor returns the same foreground color name a
// Fyne button uses for its label and icon at the given importance, so the
// loading indicator's dots can be made to match it.
func loadingButtonForegroundColor(importance widget.Importance) fyne.ThemeColorName {
	switch importance {
	case widget.DangerImportance:
		return theme.ColorNameForegroundOnError
	case widget.HighImportance:
		return theme.ColorNameForegroundOnPrimary
	case widget.SuccessImportance:
		return theme.ColorNameForegroundOnSuccess
	case widget.WarningImportance:
		return theme.ColorNameForegroundOnWarning
	default: // MediumImportance, LowImportance
		return theme.ColorNameForeground
	}
}

// activityColorTheme redirects theme.ColorNameForeground to colorName, so a
// [widget.Activity] wrapped in a [container.ThemeOverride] using this theme
// draws its dots in colorName instead of the plain foreground color.
type activityColorTheme struct {
	fyne.Theme
	colorName fyne.ThemeColorName
}

func (t *activityColorTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameForeground {
		return t.Theme.Color(t.colorName, variant)
	}
	return t.Theme.Color(name, variant)
}

// SetText sets the text of the button. While locked, the new text is
// stored and applied on unlock.
func (w *LoadingButton) SetText(label string) {
	w.label = label
	if w.button.locked {
		return
	}
	w.button.SetText(label)
}

// SetIcon sets the icon of the button. While locked, the new icon is
// stored and applied on unlock.
func (w *LoadingButton) SetIcon(icon fyne.Resource) {
	w.icon = icon
	if w.button.locked {
		return
	}
	w.button.SetIcon(icon)
}

// Disabled reports whether this widget is disabled.
func (w *LoadingButton) Disabled() bool {
	return w.disabledTemp || w.button.Disabled()
}

// Disable disables this widget.
func (w *LoadingButton) Disable() {
	if w.button.locked {
		w.disabledTemp = true
		return
	}
	w.button.Disable()
}

// Enable enables this widget.
func (w *LoadingButton) Enable() {
	if w.button.locked {
		w.disabledTemp = false
		return
	}
	w.button.Enable()
}

// AccessibilityLabel returns the label, or if there is none the name of the
// icon, that assistive technology should announce for this widget.
func (w *LoadingButton) AccessibilityLabel() string {
	if w.label != "" {
		return w.label
	}
	if w.icon != nil {
		return w.icon.Name()
	}
	return ""
}

// AccessibilityRole returns the accessibility role for this widget.
func (w *LoadingButton) AccessibilityRole() fyne.AccessibleRole {
	return fyne.AccessibleRoleButton
}

// lockableButton is an extension of the Fyne button which can be
// disabled / locked without changing it's appearance.
//
// This feature is used by the LoadingButton.
type lockableButton struct {
	widget.Button
	locked bool
}

func newLockableButton(label string, icon fyne.Resource, tapped func()) *lockableButton {
	w := &lockableButton{
		Button: widget.Button{
			Text:     label,
			Icon:     icon,
			OnTapped: tapped,
		},
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *lockableButton) unlock() {
	w.locked = false
}

func (w *lockableButton) lock() {
	w.locked = true
}

func (w *lockableButton) Tapped(pe *fyne.PointEvent) {
	if w.locked {
		return
	}
	w.Button.Tapped(pe)
}

func (w *lockableButton) MouseIn(me *desktop.MouseEvent) {
	if w.locked {
		return
	}
	w.Button.MouseIn(me)
}

func (w *lockableButton) MouseOut() {
	// Always forward, unlike MouseIn: guarding this would leave hover stuck.
	w.Button.MouseOut()
}

func (w *lockableButton) TypedKey(key *fyne.KeyEvent) {
	if w.locked {
		return
	}
	w.Button.TypedKey(key)
}

func (w *lockableButton) TypedRune(r rune) {
	if w.locked {
		return
	}
	w.Button.TypedRune(r)
}

func (w *lockableButton) FocusGained() {
	if w.locked {
		return
	}
	w.Button.FocusGained()
}

func (w *lockableButton) FocusLost() {
	// Always forward, unlike FocusGained: Fyne's FocusManager calls this only
	// once when focus moves away and never retries, so swallowing it while
	// locked would leave the button's focused state stuck permanently.
	w.Button.FocusLost()
}
