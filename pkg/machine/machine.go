/*
Package machine represents an xenigma machine, a modified version of enigma
used for encryption and decryption of text messages.

The machine consists of Rotors, Plugboard, and Reflector. machine package
contains four structs machine.Machine, machine.Plugboard, machine.Reflector,
and machine.Rotor. Each of which can be created with pre specified settings
or generated.

Encryption/decryption is done using Machine.Encrypt, which can be used simply
as the following.
    m := machine.Generate(10)
    encrypted := m.Encrypt("Hello, world!")

Components

Machine's components can be generated or set at creation/using JSON.

Rotors

Rotors are represented by two structs machine.Rotors and machine.Rotor. A
machine can have any number of rotors.

Rotor's fields are pathways, position, step, and cycle.

Pathways are the electric connections between characters. Pathways are
represented using a 26 element array where indices represent characters and
array elements represent the character they are connected to. For example
if element at index 0 is 1, then "a" (0) is connected to "b" (1).

Position is the current position of the rotor, which must be reachable from
the starting position 0 or "a".

Step is the number of positions a rotor shifts when stepping once (the size
of rotor's jump.) For example if a rotor at position "a", has step 3, then
a jump will change rotor's position to "d". The default step is 1.

Cycle is the number of rotor steps considered a full cycle, after which the
following rotor is shifted. For example, if a rotor has cycle 13, then it
needs to complete 13 steps for the following rotor to step once. The default
cycle size is 26.

To avoid position collisions and guarantee that rotor's current setting can
be reached using only one sequence of steps, (step*cycle % 26 == 0) must be
true. Combinations that don't satisfy this relation are considered wrong.

Reflector

Reflector consists of a connections array similar to pathways with a condition
that it must be symmetric, meaning that if "a" is connected to "b", then "b"
must also be connected to "a".

Plugboard

Plugboard, also, consists of a connections array with the same condition as a
reflector.

Plugboard's connections are required to have 26 elements. In order to keep a
character without a connection, connect it to itself.
*/
package machine

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	alphabetSize = 26
	configPath   = "/.config/xenigma.conf"
)

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

// Load returns a fully initialized Machine object. Configurations of
// Machine's fields are read from config file $HOME/.config/xenigma.conf
// xenigma.conf is parsed as a normal JSON file. If the file contains
// correct configurations, a machine object is initialized and returned
// with error being nil. Otherwise overwrite parameter is checked.
// If overwrite is true, random configs are generated using the specified
// number of rotors and written to file, a machine object with the same
// configs is returned and error is nil. Otherwise an initialization
// error is returned and Machine is nil.
func Load(numberOfRotors int, overwrite bool) (*Machine, error) {
	machine, err := Read(os.Getenv("HOME") + configPath)
	if err != nil {
		if overwrite {
			machine = Generate(numberOfRotors)
			if err = Write(machine, os.Getenv("HOME")+configPath); err != nil {
				return machine, fmt.Errorf("failed to initialize: %w", err)
			}
			return machine, nil
		}
		return nil, fmt.Errorf("failed to load: %w", err)
	}
	return machine, nil
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

// Plugboard returns machine's plugboard.
func (m *Machine) Plugboard() *Plugboard {
	return m.plugboard
}

// Reflector returns machine's reflector.
func (m *Machine) Reflector() *Reflector {
	return m.reflector
}
