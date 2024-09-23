package anonymity

import (
	"net"
	"testing"
)

// TestRandomIP tests the IP function
func TestIP(t *testing.T) {
	a := NewAnonymity()
	ip := a.IP()
	if net.ParseIP(ip) == nil {
		t.Errorf("Expected a valid IP address, got %s", ip)
	}
}
