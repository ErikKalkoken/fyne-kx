package widget

import (
	"testing"

	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

type embeddedIconButton struct {
	IconButton
}

func TestIconButton_Super(t *testing.T) {
	t.Run("returns itself when used directly", func(t *testing.T) {
		w := NewIconButton(theme.HomeIcon(), nil)
		assert.Same(t, w, w.super())
	})
	t.Run("returns outer widget when embedded", func(t *testing.T) {
		w := &embeddedIconButton{}
		w.ExtendBaseWidget(w)
		assert.Same(t, w, w.super())
	})
}
