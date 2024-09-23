package anonymity

import (
	"net"
)

// IP generates a random public IP address.
func (a *Generator) IP() string {
	for {
		ip := make(net.IP, 4)
		a.rand.Read(ip)
		if !isPrivateIP(ip) {
			return ip.String()
		}
	}
}

// isPrivateIP checks if the given IP address is a private IP.
func isPrivateIP(ip net.IP) bool {
	privateIPBlocks := []*net.IPNet{
		{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)},
		{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)},
		{IP: net.IPv4(127, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		{IP: net.IPv4(169, 254, 0, 0), Mask: net.CIDRMask(16, 32)},
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
