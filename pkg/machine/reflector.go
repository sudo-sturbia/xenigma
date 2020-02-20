package machine

import (
	"math/rand"

	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Reflector is a component used in xenigma machine.
type Reflector struct {
	connections [alphabetSize]int
}

// NewReflector creates and returns a plugboard with the given configuration.
// Returns an error if given configurations are incorrect.
func NewReflector(connections [alphabetSize]int) (*Reflector, error) {
	r := new(Reflector)
	if err := r.Set(connections); err != nil {
		return nil, err
	}

	return r, nil
}

// GenerateReflector generates a reflector with random configurations and
// returns a pointer to it.
func GenerateReflector() *Reflector {
	connections := [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	rand.Shuffle(alphabetSize, func(i, j int) {
		connections[i], connections[j] =
			connections[j], connections[i]
	})

	// Assign symmetric values to plugboard's connections
	r := new(Reflector)
	for i := 0; i < alphabetSize/2; i++ {
		r.connections[connections[i]], r.connections[connections[i+13]] =
			connections[i+13], connections[i]
	}

	return r
}

// Set verifies and sets given connections. Returns an error if
// connections are incorrect.
func (r *Reflector) Set(connections [alphabetSize]int) error {
	if err := r.isGivenConfigCorrect(connections); err != nil {
		return err
	}

	r.connections = connections

	return nil
}

// reflect returns the reflection of the given character using reflector's
// connections array.
func (r *Reflector) reflect(char int) int {
	return r.connections[char]
}

// IsConfigCorrect returns an error if reflector's connections are incorrect.
func (r *Reflector) IsConfigCorrect() error {
	return r.isGivenConfigCorrect(r.connections)
}

// isGivenConfigCorrect returns true if given connections are correct, false otherwise.
func (r *Reflector) isGivenConfigCorrect(connections [alphabetSize]int) error {
	if !helper.AreElementsIndices(connections[:]) {
		return &initError{"reflector's connections are incorrect"}
	}

	if !helper.IsSymmetric(connections[:]) {
		return &initError{"reflector's connections are not symmetric"}
	}

	return nil
}

// Connections returns reflector's connections array.
func (r *Reflector) Connections() [alphabetSize]int {
	return r.connections
}
