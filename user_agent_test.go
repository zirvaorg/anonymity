package anonymity

import "testing"

// TestUserAgent tests the UserAgent function
func TestUserAgent(t *testing.T) {
	a := NewAnonymity()
	userAgent := a.UserAgent()
	if userAgent == "" {
		t.Error("Expected a non-empty user agent string")
	}
}
