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
	"os"
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
	pathConnections      [][alphabetSize]int // Connections that form electric pathways
	reflector            [alphabetSize]int   // Reflector connections, symmetric
	plugboardConnections [alphabetSize]int   // Plugboard connections, symmetric

	numberOfRotors int   // Number of rotors used in the machine
	rotors         []int // Mechanical rotors' heads
	takenSteps     []int // Number of steps taken by each rotor except the last
	step           int   // Size of shift between rotor steps (move)
	cycle          int   // Number of rotor steps considered a full cycle
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
func (m *Machine) SetComponents(
	pathways [][alphabetSize]int,
	plugboard [alphabetSize]int,
	reflector [alphabetSize]int,
	rotorsPositions []int, step int, cycle int) error {

	if len(pathways) != len(rotorsPositions) {
		return &initError{"rotors and electric pathways are of different sizes"}
	}

	m.setNumberOfRotors(len(pathways))
	m.setPathConnections(pathways)
	m.setPlugboard(plugboard)
	m.setReflector(reflector)
	m.initRotors(rotorsPositions, step, cycle)

	return m.IsConfigCorrect()
}

// PathConnections returns electric pathway connections
func (m *Machine) PathConnections() [][alphabetSize]int {
	return m.pathConnections
}

// setPathConnections sets path connections array in Machine.
func (m *Machine) setPathConnections(paths [][alphabetSize]int) {
	m.pathConnections = paths
}

// Reflector returns reflector connections.
func (m *Machine) Reflector() [alphabetSize]int {
	return m.reflector
}

// setReflector sets reflector connections.
func (m *Machine) setReflector(reflector [alphabetSize]int) {
	m.reflector = reflector
}

// PlugboardConnections returns plugboard connections.
func (m *Machine) PlugboardConnections() [alphabetSize]int {
	return m.plugboardConnections
}

// setPlugboard sets reflector connections.
func (m *Machine) setPlugboard(plugboard [alphabetSize]int) {
	m.plugboardConnections = plugboard
}
