package machine

import (
	"fmt"
)

// Plugboard is a set of connections that maps different characters to each
// other. Plugboard is used as an initial step in xenigma.
type Plugboard struct {
	connections [alphabetSize]int
}

// NewPlugboard creates and returns a new plugboard, and an error if given
// connections are incorrect.
func NewPlugboard(connections [alphabetSize]int) (*Plugboard, error) {
	if err := verify(connections); err != nil {
		return nil, err
	}

	return &Plugboard{
		connections: connections,
	}, nil
}

// GeneratePlugboard generates a plugboard with random configurations and
// returns a pointer to it.
func GeneratePlugboard() *Plugboard {
	return &Plugboard{
		connections: generateConnections(),
	}
}

// Connections returns plugboard's connections array.
func (p *Plugboard) Connections() [alphabetSize]int {
	return p.connections
}

// PlugIn returns the int mapped to char based on plugboard's
// connections. Should be used when a character is entered.
func (p *Plugboard) PlugIn(char byte) int {
	return p.connections[int(char-'a')]
}

// PlugOut returns the byte mapped to char based on plugboard's
// connections. Should be used when a character is returned.
func (p *Plugboard) plugOut(char int) byte {
	return byte(p.connections[char] + 'a')
}

// Verify returns an error if plugboard's connections are incorrect.
func (p *Plugboard) Verify() error {
	return verify(p.connections)
}

// verify verifies that given connections are valid, and returns an error
// if not.
func verify(connections [alphabetSize]int) error {
	if !areElementsIndices(connections[:]) {
		return fmt.Errorf("connections are invalid")
	}

	if !isSymmetric(connections[:]) {
		return fmt.Errorf("connections are not symmetric")
	}

	return nil
}
