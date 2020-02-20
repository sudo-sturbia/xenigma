package machine

import (
	"math/rand"

	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Plugboard is a component used in xenigma machine.
type Plugboard struct {
	connections [alphabetSize]int
}

// NewPlugboard creates and returns a plugboard with the given configuration.
// Returns an error if given configurations are incorrect.
func NewPlugboard(connections [alphabetSize]int) (*Plugboard, error) {
	p := new(Plugboard)
	if err := p.Set(connections); err != nil {
		return nil, err
	}

	return p, nil
}

// GeneratePlugboard generates a plugboard with random configurations and
// returns a pointer to it.
func GeneratePlugboard() *Plugboard {
	connections := [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

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

// Set verifies and sets given connections. Returns an error if
// connections are incorrect.
func (p *Plugboard) Set(connections [alphabetSize]int) error {
	if err := p.isGivenConfigCorrect(connections); err != nil {
		return err
	}

	p.connections = connections

	return nil
}

// plugIn changes a byte (character) to an int (0 -> 25) based on
// plugboard connections. Used when character is entered.
func (p *Plugboard) plugIn(char byte) int {
	return p.connections[int(char-'a')]
}

// plugOut changes an int to a byte (character) based on
// plugboard connections. Used when character is returned.
func (p *Plugboard) plugOut(char int) byte {
	return byte(p.connections[char] + 'a')
}

// IsConfigCorrect returns an error if plugboard's connections are incorrect.
func (p *Plugboard) IsConfigCorrect() error {
	return p.isGivenConfigCorrect(p.connections)
}

// isGivenConfigCorrect returns an error if given connections are incorrect.
func (p *Plugboard) isGivenConfigCorrect(connections [alphabetSize]int) error {
	if !helper.AreElementsIndices(connections[:]) {
		return &initError{"plugboard's connections are incorrect"}
	}

	if !helper.IsSymmetric(connections[:]) {
		return &initError{"plugboard's connections are not symmetric"}
	}

	return nil
}

// Connections returns plugboard's connections array.
func (p *Plugboard) Connections() [alphabetSize]int {
	return p.connections
}
