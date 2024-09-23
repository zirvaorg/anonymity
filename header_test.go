package anonymity

import "testing"

// TestRandomHeader tests the RandomHeader method.
func TestHeader(t *testing.T) {
	a := New()
	ip := "192.0.2.1"
	headers := a.Header(ip)
	if len(headers) != 8 {
		t.Errorf("Expected 8 headers, got %d", len(headers))
	}
	for _, header := range headers {
		if header.Val != ip {
			t.Errorf("Expected header value %s, got %s", ip, header.Val)
		}
	}
}
