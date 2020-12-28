package machine

// Reflector is a set of connections that maps two characters to each other.
// Reflector is used as a middle step in xenigma.
type Reflector struct {
	connections [alphabetSize]int
}

// NewReflector creates and returns a new reflector. An error is returned if
// given connections are incorrect.
func NewReflector(connections [alphabetSize]int) (*Reflector, error) {
	if err := verify(connections); err != nil {
		return nil, err
	}

	return &Reflector{
		connections: connections,
	}, nil
}

// GenerateReflector generates a reflector with random configurations and
// returns a pointer to it.
func GenerateReflector() *Reflector {
	return &Reflector{
		connections: generateConnections(),
	}
}

// Connections returns reflector's connections array.
func (r *Reflector) Connections() [alphabetSize]int {
	return r.connections
}

// reflect returns the reflection of the given character using reflector's
// connections array.
func (r *Reflector) Reflect(char int) int {
	return r.connections[char]
}

// Verify returns an error if Reflector's connections are incorrect.
func (r *Reflector) Verify() error {
	return verify(r.connections)
}
