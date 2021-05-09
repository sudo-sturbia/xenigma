package machine

// Contains Plugboard and Reflector struct and their needed functions. This file
// is named mappers simply because both reflector, and plugboard are used to create
// mapping of one character to another.

import (
	"fmt"
	"math/rand"
)

// Plugboard is a set of connections that maps different characters to each
// other. Plugboard is used as an initial step in xenigma.
type Plugboard struct {
	connections map[int]int
}

// Reflector is a set of connections that maps two characters to each other.
// Reflector is used as a middle step in xenigma.
type Reflector struct {
	connections map[int]int
}

// NewPlugboard creates and returns a new plugboard, and an error if given
// connections are incorrect.
func NewPlugboard(connections map[int]int) (*Plugboard, error) {
	if err := verifyConnections(connections); err != nil {
		return nil, err
	}

	return &Plugboard{
		connections: connections,
	}, nil
}

// NewReflector creates and returns a new reflector. An error is returned if
// given connections are incorrect.
func NewReflector(connections map[int]int) (*Reflector, error) {
	if err := verifyConnections(connections); err != nil {
		return nil, err
	}

	return &Reflector{
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

// GenerateReflector generates a reflector with random configurations and
// returns a pointer to it.
func GenerateReflector() *Reflector {
	return &Reflector{
		connections: generateConnections(),
	}
}

// generateConnections generates a random map of symmetric connections populated
// with elements 0 through n-1. Symmetric means that if slice[n] = m, then
// slice[m] = n.
func generateConnections() map[int]int {
	var ordered [alphabetSize]int
	for i := 0; i < alphabetSize; i++ {
		ordered[i] = i
	}

	rand.Shuffle(
		alphabetSize,
		func(i, j int) {
			ordered[i], ordered[j] = ordered[j], ordered[i]
		},
	)

	connections := make(map[int]int)
	for i := 0; i < alphabetSize/2; i++ {
		connections[ordered[i]], connections[ordered[i+13]] = ordered[i+13], ordered[i]
	}
	return connections
}

// Connections returns plugboard's connections map.
func (p *Plugboard) Connections() map[int]int {
	return p.connections
}

// Connections returns reflector's connections map.
func (r *Reflector) Connections() map[int]int {
	return r.connections
}

// PlugIn returns the int mapped to char based on plugboard's
// connections. Should be used when a character is entered.
func (p *Plugboard) PlugIn(char byte) int {
	return p.connections[int(char-'a')]
}

// PlugOut returns the byte mapped to char based on plugboard's
// connections. Should be used when a character is returned.
func (p *Plugboard) PlugOut(char int) byte {
	return byte(p.connections[char] + 'a')
}

// Reflect returns the reflection of the given character using reflector's
// connections array.
func (r *Reflector) Reflect(char int) int {
	return r.connections[char]
}

// Verify returns an error if plugboard's connections are incorrect.
func (p *Plugboard) Verify() error {
	return verifyConnections(p.connections)
}

// Verify returns an error if Reflector's connections are incorrect.
func (r *Reflector) Verify() error {
	return verifyConnections(r.connections)
}

// verifyConnections verifies that given connections are valid, and returns an
// error if not.
func verifyConnections(connections map[int]int) error {
	if !zeroToN(connections, alphabetSize) {
		return fmt.Errorf("connections are invalid")
	}

	if !isSymmetric(connections) {
		return fmt.Errorf("connections are not symmetric")
	}

	return nil
}
