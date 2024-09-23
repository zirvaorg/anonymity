package anonymity

import (
	"math/rand"
	"time"
)

// Anonymity provides methods to generate random IP addresses, HTTP headers, and user agents.
type Anonymity struct {
	rand *rand.Rand
}

// NewAnonymity creates a new Anonymity instance.
func NewAnonymity() *Anonymity {
	source := rand.NewSource(time.Now().UnixNano())
	return &Anonymity{rand: rand.New(source)}
}
