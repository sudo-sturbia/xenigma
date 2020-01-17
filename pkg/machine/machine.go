/*
	Package machine represents an enigma machine used for encryption and
	decryption of text messages.

	The machine consists of the following components: electric pathways,
	reflector, plugboard, and 3 rotors. Components are configurable through
	JSON file $HOME/.config/enigma.json. See documentation of machine.New
	for more details.

	enigma.json should be of the following form.
		{
			"pathways": [
				["a", "b", "c", "d", ... "z"],
				["a", "b", "c", "d", ... "z"],
				["a", "b", "c", "d", ... "z"]
			],
			"reflector": ["a", "b", "c", "d", ... "z"],
			"plugboard": ["a", "b", "c", "d", ... "z"],
			"rotorPositions": ["a", "b", "c"]
		}

	All arrays, except for "rotorPositions", should have 26 elements.
	Both reflector and plugboard should be symmetric (every two elements
	are connected to each other).

	Licensed under MIT license @github.com/sudo-sturbia
*/
package machine

import (
	"fmt"
	"os"
)

const (
	numberOfRotors = 3
	alphabetSize   = 26
)

// Machine represents an enigma machine with mechanical components.
// Components are electric pathways, reflector, plugboard, and rotors.
type Machine struct {
	pathConnections      [numberOfRotors][alphabetSize]int // Connections that form electric pathways
	reflector            [alphabetSize]int                 // Reflector connections, symmetric
	plugboardConnections [alphabetSize]int                 // Plugboard connections, symmetric

	rotors     [numberOfRotors][alphabetSize]int // Mechanical rotors, 1st element is rotor's position
	takenSteps [numberOfRotors - 1]int           // Number of steps taken by each rotor except the last
	step       int                               // Size of shift between rotor steps (move)
	cycle      int                               // Number of rotor steps considered a full cycle
}

// New returns a newly created, fully initialized Machine object.
// Machine's fields are read from config file $HOME/.config/enigma.json
// If the file contains correct configurations, a machine object is
// initialized and returned with error being nil.
// Otherwise overwrite parameter is checked. If overwrite is true, random
// configs are generated and written to file, a machine object with the
// same configs is returned and error is nil. Otherwise an initialization
// error is returned and Machine is nil.
func New(overwrite bool) (*Machine, error) {
	machine, err := read(os.Getenv("HOME") + "/.config/enigma.json")

	if err != nil {
		if overwrite {
			machine = randMachine()
			if err = write(machine, os.Getenv("HOME")+"/.config/enigma.json"); err != nil {
				return machine, &initError{err.Error()}
			}

			return machine, nil
		}

		return nil, &initError{err.Error()}
	}

	return machine, nil
}

// PathConnections returns electric pathway connections
func (m *Machine) PathConnections() [numberOfRotors][alphabetSize]int {
	return m.pathConnections
}

// setPathConnections sets path connections array in Machine.
func (m *Machine) setPathConnections(paths [numberOfRotors][alphabetSize]int) {
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
