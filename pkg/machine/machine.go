/*
Package machine represents an xenigma machine, a modified version of enigma
used for encryption and decryption of text messages.

The machine consists of Rotors, Plugboard, and Reflector. machine package
contains four structs machine.Machine, machine.Plugboard, machine.Reflector,
and machine.Rotor.

Each of the structs has at least two functions and two methods, named
according to the following
	machine.New()       // Creates a struct with the given fields.
	machine.Generate()  // Generates a struct with random configs.

	m.Set()             // Set the struct's fields.
	m.IsConfigCorrect() // Return an error if struct is incorrect.

machine.New returns a pointer to a new machine.Machine, for other structs
New should be followed by the struct's name, i.e. machine.NewRotor, etc.
and the same goes for machine.Generate

A machine object can be created using machine.Load, machine.Generate,
machine.Read, or by creating a machine pointer and using available setters
to specify machine's components.

machine.Load is meant for usage in an independent program, the other
three options are more suitable for usage when the package is imported.

Encryption / decryption is done using Machine.Encrypt, which can be used
simply as the following.
    m := machine.Generate(10)
    encrypted := m.Encrypt("Hello, world!")

Configuration

Every component of xenigma can be configured through JSON, or by using
available setters. An example of a JSON config file is the following
	{
		"rotors": [
			{
				"pathways": ["a", "b", "c", ...],
				"position": "a",
				"step": 1,
				"cycle": 26
			},
			{
				"pathways": ["a", "b", "c", ...],
				"position": "b",
				"step": 1,
				"cycle": 26
			},
			{
				"pathways": ["a", "b", "c", ...],
				"position": "c",
				"step": 1,
				"cycle": 26
			}
		],

		"reflector": {
			"connections": ["a", "b", "c", ...]
		},

		"plugboard": {
			"connections": ["a", "b", "c", ...]
		}
	}

Rotors

Rotors are represented using machine.Rotor struct. A xenigma machine
can have any number of rotors, the number of rotors is size the of "rotors"
array in JSON or the size of the slice given to the setter.

Rotor's fields are: pathways, position, step, and cycle.

Pathways are the electric connections between characters. Pathways
are represented using a 26 element array where indices represent characters
and array elements represent the character they are connected to. For
example if element at index 0 is "b", then "a" (character 0) is connected
to "b".

Position is the current position of the rotor, which must be reachable
from the starting position "a".

Step is the number of positions a rotor shifts when stepping once
(the size of rotor's jump.) For example if a rotor at position "a", with
step = 3 steps once, then rotor's position changes to "d". The default step
size is 1.

Cycle is the number of rotor steps considered a full cycle, after which
the following rotor steps (is shifted.) For example, if a rotor has
a cycle = 13, then the rotor needs to complete 13 steps in order for the
following rotor to step once. The default cycle size is 26.

To avoid position collisions and guarantee that a rotor configuration
can be reached using only one sequence of steps, (step*cycle) must divide
26 -> (step*cycle % 26 == 0). Combinations that don't satisfy that relation
are considered wrong.

Reflector

Reflector consists of a connections array similar to pathways with a condition
that it must be symmetric, meaning that if "a" is connected to "b", then
"b" must also be connected to "a".

Plugboard

Plugboard, also, consists of a connections array exactly the same as a
reflector.
Plugboard's connections are required to have 26 elements, so characters
not connected to anything should be connected to themselves (in order not
to be transformed.)

Licensed under MIT license @github.com/sudo-sturbia
*/
package machine

import (
	"fmt"
	"math/rand"
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
