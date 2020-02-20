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
	"math/rand"
	"os"
	"time"

	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

const alphabetSize = 26

// Machine represents an enigma machine with mechanical components.
// Components are electric pathways, reflector, plugboard, and rotors.
type Machine struct {
	reflector [alphabetSize]int // Reflector connections, symmetric
	plugboard [alphabetSize]int // Plugboard connections, symmetric

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

// SetComponents initializes all components of the machine.
// Returns an error if given incorrect configurations.
func (m *Machine) SetComponents(rotors []*Rotor, plugboard [alphabetSize]int, reflector [alphabetSize]int) error {
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

	m.GeneratePlugboard()
	m.GenerateReflector()

	return m
}

// IsConfigCorrect verifies that all fields of the machine are
// initialized correctly, returns an error if not.
func (m *Machine) IsConfigCorrect() error {
	switch {
	case !m.areRotorsCorrect():
		return &initError{"rotors' configurations are invalid"}
	case !m.isReflectorCorrect():
		return &initError{"reflector connections are incorrect"}
	case !m.isPlugboardCorrect():
		return &initError{"plugboard connections are incorrect"}
	default:
		return nil
	}
}

// Reflector returns machine's reflector connections.
func (m *Machine) Reflector() [alphabetSize]int {
	return m.reflector
}

// Plugboard returns machine's plugboard connections.
func (m *Machine) Plugboard() [alphabetSize]int {
	return m.plugboard
}

// SetReflector sets reflector connections. Returns an error
// if given connections are incorrect.
func (m *Machine) SetReflector(reflector [alphabetSize]int) error {
	m.reflector = reflector
	if !m.isReflectorCorrect() {
		return &initError{"given reflector connections are incorrect"}
	}

	return nil
}

// SetPlugboard sets plugboard connections. Returns an error
// if given connections are incorrect.
func (m *Machine) SetPlugboard(plugboard [alphabetSize]int) error {
	m.plugboard = plugboard
	if !m.isPlugboardCorrect() {
		return &initError{"given plugboard connections are incorrect"}
	}

	return nil
}

// GenerateReflector creates random reflector connections for the current machine.
func (m *Machine) GenerateReflector() {
	half := []int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	rand.Shuffle(alphabetSize/2, func(i, j int) {
		half[i], half[j] = half[j], half[i]
	})

	// Assign symmetric values to machine's reflector
	for i := 0; i < alphabetSize/2; i++ {
		m.reflector[i], m.reflector[half[i]] = half[i], i
	}
}

// GeneratePlugboard creates random plugboard connections for the current machine.
func (m *Machine) GeneratePlugboard() {
	half := []int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	rand.Shuffle(alphabetSize/2, func(i, j int) {
		half[i], half[j] = half[j], half[i]
	})

	// Assign symmetric values to machine's plugboard
	for i := 0; i < alphabetSize/2; i++ {
		m.plugboard[i], m.plugboard[half[i]] = half[i], i
	}
}

// isReflectorCorrect returns true if reflector is initialized correctly.
func (m *Machine) isReflectorCorrect() bool {
	return helper.AreElementsIndices(m.reflector[:]) &&
		helper.IsSymmetric(m.reflector[:])
}

// isPlugboardCorrect returns true if plugboard connections initialized correctly.
func (m *Machine) isPlugboardCorrect() bool {
	return helper.AreElementsIndices(m.plugboard[:]) &&
		helper.IsSymmetric(m.plugboard[:])
}
