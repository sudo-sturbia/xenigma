/*
Package machine represents an xenigma machine which is a modified version
of enigma used for encryption and decryption of text messages.

The machine consists of Rotors, plugboard, and reflector. machine package
contains two structs machine.Machine, and machine.Rotor.

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

		"reflector": ["a", "b", "c", ...],
		"plugboard": ["a", "b", "c", ...]
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

Reflector is a connections array similar to pathways with a condition
that it must be symmetric, meaning that if "a" is connected to "b", then
"b" must also be connected to "a".

Plugboard

Plugboard is also a connections array exactly the same as a reflector.
Note that the plugboard is required to have 26 elements, so characters
not connected to anything should be connected to themselves (so that
they wouldn't be transformed.)

Licensed under MIT license @github.com/sudo-sturbia
*/
package machine

import (
	"math/rand"
	"os"
	"time"
)

const alphabetSize = 26

// Machine represents an enigma machine with mechanical components.
// Components are electric pathways, reflector, plugboard, and rotors.
type Machine struct {
	reflector *Reflector // Machine's reflector
	plugboard *Plugboard // Machine's plugboard

	rotors         []*Rotor // Machine's mechanical rotors
	numberOfRotors int      // Number of rotors used in the machine
}

// Initializion error
type initError struct {
	message string
}

func (err *initError) Error() string {
	return "incorrect init, " + err.message
}

// New creates and returns a new machine object with the given configurations.
// Returns an error if given configurations are incorrect.
func New(rotors []*Rotor, plugboard *Plugboard, reflector *Reflector) (*Machine, error) {
	m := new(Machine)
	if err := m.Set(rotors, plugboard, reflector); err != nil {
		return nil, err
	}

	return m, nil
}

// Load returns a fully initialized Machine object. Configurations of
// Machine's fields are read from config file $HOME/.config/xenigma.json
// If the file contains correct configurations, a machine object is
// initialized and returned with error being nil.
// Otherwise overwrite parameter is checked. If overwrite is true, random
// configs are generated using the specified number of rotors and written
// to file, a machine object with the same configs is returned and error
// is nil. Otherwise an initialization error is returned and Machine is nil.
func Load(numberOfRotors int, overwrite bool) (*Machine, error) {
	machine, err := Read(os.Getenv("HOME") + "/.config/xenigma.json")

	if err != nil {
		if overwrite {
			machine = Generate(numberOfRotors)
			if err = machine.Write(os.Getenv("HOME") + "/.config/xenigma.json"); err != nil {
				return machine, &initError{err.Error()}
			}

			return machine, nil
		}

		return nil, &initError{err.Error()}
	}

	return machine, nil
}

// Generate creates a machine object with a specified number of rotors
// containing randomly generated component configurations.
func Generate(numberOfRotors int) *Machine {
	rand.Seed(time.Now().UnixNano())

	m := new(Machine)

	rotors := make([]*Rotor, numberOfRotors)
	for i := 0; i < numberOfRotors; i++ {
		rotors[i] = GenerateRotor()
	}
	m.SetRotors(rotors)

	m.plugboard = GeneratePlugboard()
	m.reflector = GenerateReflector()

	return m
}

// Set initializes all components of the machine.
// Returns an error if given incorrect configurations.
func (m *Machine) Set(rotors []*Rotor, plugboard *Plugboard, reflector *Reflector) error {
	if err := m.SetRotors(rotors); err != nil {
		return err
	}

	if err := m.SetPlugboard(plugboard); err != nil {
		return err
	}

	if err := m.SetReflector(reflector); err != nil {
		return err
	}

	return m.IsConfigCorrect()
}

// IsConfigCorrect verifies that all fields of the machine are
// initialized correctly, returns an error if not.
func (m *Machine) IsConfigCorrect() error {
	if err := m.AreRotorsCorrect(); err != nil {
		return err
	}

	if err := m.reflector.IsConfigCorrect(); err != nil {
		return err
	}

	if err := m.plugboard.IsConfigCorrect(); err != nil {
		return err
	}

	return nil
}

// SetPlugboard sets machine's plugboard. Returns an error
// if given connections are incorrect.
func (m *Machine) SetPlugboard(plugboard *Plugboard) error {
	m.plugboard = plugboard
	if err := plugboard.IsConfigCorrect(); err != nil {
		return err
	}

	return nil
}

// SetReflector sets machine's reflector. Returns an error
// if given connections are incorrect.
func (m *Machine) SetReflector(reflector *Reflector) error {
	m.reflector = reflector
	if err := reflector.IsConfigCorrect(); err != nil {
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
