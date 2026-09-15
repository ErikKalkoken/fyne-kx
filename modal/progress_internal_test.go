package modal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCancelRegistry_RegisterThenCancel(t *testing.T) {
	c := &cancelRegistry{}
	called := make(chan struct{})

	c.register(func() {
		close(called)
	})
	c.requestCancel()

	select {
	case <-called:
	case <-time.After(2 * time.Second):
		t.Fatal("handler was not called")
	}
}

func TestCancelRegistry_CancelThenRegister(t *testing.T) {
	c := &cancelRegistry{}
	called := make(chan struct{})

	c.requestCancel()
	c.register(func() {
		close(called)
	})

	select {
	case <-called:
	case <-time.After(2 * time.Second):
		t.Fatal("handler registered after cancel should run immediately")
	}
}

func TestCancelRegistry_RequestCancelIsIdempotent(t *testing.T) {
	c := &cancelRegistry{}
	callCount := 0

	c.register(func() {
		callCount++
	})
	c.requestCancel()
	c.requestCancel()

	assert.Equal(t, 1, callCount)
}

func TestCancelRegistry_RegisterWithoutCancelDoesNothing(t *testing.T) {
	c := &cancelRegistry{}
	called := false

	c.register(func() {
		called = true
	})

	assert.False(t, called)
}
