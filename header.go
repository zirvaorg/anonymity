package anonymity

// Header represents an HTTP header.
type Header struct {
	Key string
	Val string
}

// Header generates a random header with the given IP address.
func (a *Anonymity) Header(ip string) []Header {
	headers := []Header{
		{"X-Originating-IP", ip},
		{"X-Remote-IP", ip},
		{"X-Remote-Addr", ip},
		{"X-Client-IP", ip},
		{"X-Host", ip},
		{"X-Forwarded-Host", ip},
		{"X-Forwarded-For", ip},
		{"X-Forwarded-Server", ip},
	}

	return headers
}
