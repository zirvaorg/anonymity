package anonymity

import (
	"math/rand"
	"time"
)

// Generator provides methods to generate random IP addresses, HTTP headers, and user agents.
type Generator struct {
	rand *rand.Rand
}

// New creates a new Anonymity instance.
func New() *Generator {
	source := rand.NewSource(time.Now().UnixNano())
	return &Generator{rand: rand.New(source)}
}
