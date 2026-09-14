package modal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseChannelIfOpen_ClosesOpenChannel(t *testing.T) {
	c := make(chan struct{})

	closeChannelIfOpen(c)

	select {
	case <-c:
		// channel is closed, as expected
	default:
		t.Fatal("channel should be closed")
	}
}

func TestCloseChannelIfOpen_IsSafeOnAlreadyClosedChannel(t *testing.T) {
	c := make(chan struct{})
	close(c)

	assert.NotPanics(t, func() {
		closeChannelIfOpen(c)
	})
}
