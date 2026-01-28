package assert

import (
	"testing"
)

// HOLY SHIT A GENERIC FUNCTION!!!!!

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got=%v, want=%v", actual, expected)
	}
}
