package anonymity

import "testing"

// TestNewAnonymity tests the creation of a new Anonymity instance.
func TestNewAnonymity(t *testing.T) {
	a := New()
	if a == nil {
		t.Error("Expected a non-nil Anonymity instance")
	}
}
