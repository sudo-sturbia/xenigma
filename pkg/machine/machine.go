/*
Package machine represents an xenigma machine, a modified version of enigma
used for encryption and decryption of text messages.

The machine consists of Rotors, Plugboard, and Reflector. Each of which has
can be created with pre-specified settings, or generated.

Encryption/decryption is done using Machine.Encrypt, which can be used as the
following.
    m := machine.Generate(10)
    encrypted := m.Encrypt("Hello, world!")

Components

Machine's components can be generated or specified at creation, or read as
JSON.

Rotors

A machine can have any number of rotors, the number of rotors is the size
of the given rotor array. Rotor's fields are pathways, position, step, and
cycle.

Pathways are the electric connections between characters. They are represented
using a map-like 26 element array where an index and a character represent a
map pair. Key and value pairs are translated into their position in the english
alphabet. For example, if pathways[0]=2, then a is mapped to c. Arrays are chosen
over maps for pathways because ordering matters.

Position is the current position of the rotor, which must be reachable from
the starting position 0 or "a".

Step is the number of positions a rotor jumps when moving one step forward.
For example, if a rotor with position="a" and step="3" jumps once, the position
will change to "d". The default step is 1.

Cycle is the number of steps needed to complete a full cycle, after which the
following rotor is shifted. For example, if a rotor with cycle=13, then it
needs to complete 13 steps for the next rotor to move one step. The default
cycle is 26.

To avoid position collisions and guarantee that any of the rotor's settings can
be reached using only one sequence of steps, (step*cycle)%26 must equal zero.
Combinations that don't satisfy this relation are considered invalid.

Reflector

Reflector is connections map, which must contain all characters in the english
alphabet, and must be symmetric. Symmetry means that if "a" is connected to "b",
then "b" must also be connected to "a".

Plugboard

Plugboard is also a connections map similar to reflector. To keep a character
unconnected/unplugged, connect it to itself.
*/
package machine

import (
	"fmt"
	"math/rand"
	"time"
)

const alphabetSize = 26

// Machine represents a xenigma encryption machine. Machine's components are
// electric pathways, reflector, plugboard, and rotors.
type Machine struct {
	rotors    *Rotors
	plugboard *Plugboard
	reflector *Reflector
}

// New creates and returns a new, initialized Machine, and an error if any of
// the given fields is invalid.
func New(rotors *Rotors, plugboard *Plugboard, reflector *Reflector) (*Machine, error) {
	if err := verifyMachine(rotors, plugboard, reflector); err != nil {
		return nil, err
	}

	return &Machine{
		rotors:    rotors,
		plugboard: plugboard,
		reflector: reflector,
	}, nil
}

// Generate generates a machine with the specified number of rotors containing
// randomly generated component configurations.
func Generate(numberOfRotors int) *Machine {
	rand.Seed(time.Now().UnixNano())
	return &Machine{
		plugboard: GeneratePlugboard(),
		reflector: GenerateReflector(),
		rotors:    GenerateRotors(numberOfRotors),
	}
}

// Verify verifies that all components of the machine are initialized
// correctly, and returns an error if not.
func (m *Machine) Verify() error {
	return verifyMachine(m.rotors, m.plugboard, m.reflector)
}

func verifyMachine(rotors *Rotors, plugboard *Plugboard, reflector *Reflector) error {
	if rotors == nil {
		return fmt.Errorf("no rotors given")
	}
	if err := rotors.Verify(); err != nil {
		return err
	}

	if reflector == nil {
		return fmt.Errorf("no reflector given")
	}
	if err := reflector.Verify(); err != nil {
		return err
	}

	if plugboard == nil {
		return fmt.Errorf("no plugboard given")
	}
	if err := plugboard.Verify(); err != nil {
		return err
	}

	return nil
}

// Rotors returns machine's rotors.
func (m *Machine) Rotors() *Rotors {
	return m.rotors
}

// Plugboard returns machine's plugboard.
func (m *Machine) Plugboard() *Plugboard {
	return m.plugboard
}

// Reflector returns machine's reflector.
func (m *Machine) Reflector() *Reflector {
	return m.reflector
}
