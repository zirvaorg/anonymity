package anonymity

import "testing"

// TestUserAgent tests the UserAgent function
func TestUserAgent(t *testing.T) {
	a := New()
	userAgent := a.UserAgent()
	if userAgent == "" {
		t.Error("Expected a non-empty user agent string")
	}
}
