package widget

import (
	"testing"

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
