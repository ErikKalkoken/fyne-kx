package widget

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestFtoa(t *testing.T) {
	cases := []struct {
		name string
		in   float64
		want string
	}{
		{"200", 200, "200"},
		{"20", 20.0, "20"},
		{"2", 2, "2"},
		{"2.2", 2.2, "2.2"},
		{"2.02", 2.02, "2.02"},
		{"200.02", 200.02, "200.02"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ftoa(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSlider_DefaultStepMatchesFyneSliderDefault(t *testing.T) {
	test.NewTempApp(t)

	w := NewSlider(0, 10)
	r := w.CreateRenderer().(*sliderRenderer)

	assert.Equal(t, float64(1), r.slider.Step)
}

func TestSlider_QuantizesValueByDefaultStep(t *testing.T) {
	test.NewTempApp(t)

	w := NewSlider(0, 10)
	r := w.CreateRenderer().(*sliderRenderer)

	// Simulate the internal slider settling on an arbitrary, non-integer
	// value, as could happen from a drag landing between pixel steps.
	r.slider.SetValue(3.482759)

	assert.Equal(t, "3", r.label.Text)
}
