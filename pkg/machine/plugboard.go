package machine

import (
	"math/rand"
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
	connections := [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	rand.Shuffle(alphabetSize, func(i, j int) {
		connections[i], connections[j] =
			connections[j], connections[i]
	})

	// Assign symmetric values to plugboard's connections
	p := new(Plugboard)
	for i := 0; i < alphabetSize/2; i++ {
		p.connections[connections[i]], p.connections[connections[i+13]] =
			connections[i+13], connections[i]
	}

	return p
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
		return &initError{"plugboard's connections are incorrect"}
	}

	if !isSymmetric(connections[:]) {
		return &initError{"plugboard's connections are not symmetric"}
	}

	return nil
}
