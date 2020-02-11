/*
Package machine represents an xenigma machine which is a modified version
of the enigma machine used for encryption and decryption of text messages.

The machine consists of the following components: electric pathways,
reflector, plugboard, and a variable number of rotors.

A machine object can be created using machine.Load, machine.Generate,
machine.Read, or by simply creating a empty Machine object and calling
SetComponents method.

machine.Load is meant for usage in an independent program. The other
three options are more suitable for usage when the package is imported.

Encrypt is the method used for encryption (or decryption) of strings
and is simply used as the following.
    m := machine.Generate(3)
    encrypted := m.Encrypt("message")

Configuration

All components of the machine can be configured through JSON or using the
collective setter SetComponents. An example of a JSON config file is the
following
	{
		"pathways": [
			["a", "b", "c", ...],
			["a", "b", "c", ...],
			["a", "b", "c", ...]
		],
		"reflector": ["a", "b", "c", ...],
		"plugboard": ["a", "b", "c", ...],
		"rotorPositions": ["a", "b", "c"],
		"rotorStep": 1,
		"rotorCycle": 26
	}

Connections: pathways, reflector, and plugboard, are all specified through
JSON arrays where elements' indices represent a character's position in
the alphabet. Meaning that if, in reflector array, element at index 0 is
"b" then "a" is connected to "b".

A machine can have any number of rotors given that it's > 0. A machine's
step and cycle sizes can also be configured.

Step represents the size of the shift between rotor steps. For example if
step size is 2 then rotors jump 2 positions when shifting. A rotor at position
"a" jumps to "c".

Cycle represents the number of steps considered a full cycle, after which
the following rotor is shifted. For example in a 3-rotor machine if cycle
size is 13 then the second rotor is shifted once every time the first rotor
completes 13 steps, the third rotor operates similarly but depends on second
rotor's movement, etc.

To avoid collisions and guarantee that a rotor configuration can be reached
using only one step sequence (step x cycle) must divide 26 (alphabet size).
Combinations that don't satisfy that relation are considered wrong.

Licensed under MIT license @github.com/sudo-sturbia
*/
package machine

import (
	"math/rand"
	"os"
	"time"
)

const (
	alphabetSize = 26

	// DefaultStep represents a Machine's default step size.
	DefaultStep = 1

	// DefaultCycle represents a Machine's default cycle size.
	DefaultCycle = 26
)

// Machine represents an enigma machine with mechanical components.
// Components are electric pathways, reflector, plugboard, and rotors.
type Machine struct {
	reflector [alphabetSize]int // Reflector connections, symmetric
	plugboard [alphabetSize]int // Plugboard connections, symmetric

	rotors         []Rotor // Machine's mechanical rotors
	numberOfRotors int     // Number of rotors used in the machine
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
func (m *Machine) SetComponents(rotors []Rotor, plugboard [alphabetSize]int, reflector [alphabetSize]int) error {
	if err := m.SetRotors(rotors); err != nil {
		return err
	}

	m.SetPlugboard(plugboard)
	m.SetReflector(reflector)

	return m.IsConfigCorrect()
}

// Generate creates a machine object with a specified number of rotors
// containing randomly generated component configurations.
func Generate(numberOfRotors int) *Machine {
	rand.Seed(time.Now().UnixNano())

	m := new(Machine)

	rotors := make([]Rotor, numberOfRotors)
	for i := 0; i < numberOfRotors; i++ {
		rotors[i] = *GenerateRotor()
	}
	m.SetRotors(rotors)

	m.GeneratePlugboard()
	m.GenerateReflector()

	return m
}

// Reflector returns reflector connections.
func (m *Machine) Reflector() [alphabetSize]int {
	return m.reflector
}

// SetReflector sets reflector connections.
func (m *Machine) SetReflector(reflector [alphabetSize]int) {
	m.reflector = reflector
}

// GenerateReflector creates random reflector connections.
func (m *Machine) GenerateReflector() {
	// Create half of the connections
	halfConnections := [alphabetSize / 2]int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	rand.Shuffle(alphabetSize/2, func(i, j int) {
		halfConnections[i], halfConnections[j] =
			halfConnections[j], halfConnections[i]
	})

	for i := 0; i < len(halfConnections); i++ {
		m.reflector[i] = halfConnections[i]
		m.reflector[halfConnections[i]] = i
	}
}

// PlugboardConnections returns plugboard connections.
func (m *Machine) PlugboardConnections() [alphabetSize]int {
	return m.plugboard
}

// SetPlugboard sets plugboard connections.
func (m *Machine) SetPlugboard(plugboard [alphabetSize]int) {
	m.plugboard = plugboard
}

// GeneratePlugboard creates random plugboard connections.
func (m *Machine) GeneratePlugboard() {
	// Create half of the connections
	halfConnections := [alphabetSize / 2]int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	rand.Shuffle(alphabetSize/2, func(i, j int) {
		halfConnections[i], halfConnections[j] =
			halfConnections[j], halfConnections[i]
	})

	for i := 0; i < len(halfConnections); i++ {
		m.plugboard[i] = halfConnections[i]
		m.plugboard[halfConnections[i]] = i
	}
}
