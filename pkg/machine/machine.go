/*
Package machine represents an enigma machine used for encryption and
decryption of text messages.

The machine consists of the following components: electric pathways,
reflector, plugboard, and a variable number rotors. All of which can
be configured through JSON or a collective setter.

A machine can be created using machine.Load, machine.Generate, machine.Read,
or by simply creating a Machine object and calling SetComponents method.

machine.Load is meant for usage in an independent program. The other
three options are more suitable for usage when the package is imported.

A machine can have any number of rotors given that it's > 0. Other
properties related to rotors such as step and cycle sizes can also be
configured.

Encrypt method is used to encrypt (or decrypt) a given message as
the following example
    m := machine.Generate(3)
    encrypted := m.Encrypt("message")

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
// Machine's fields are read from config file $HOME/.config/enigma.json
// If the file contains correct configurations, a machine object is
// initialized and returned with error being nil.
// Otherwise overwrite parameter is checked. If overwrite is true, random
// configs are generated using the specified number of rotors and written
// to file, a machine object with the same configs is returned and error
// is nil. Otherwise an initialization error is returned and Machine is nil.
func Load(numberOfRotors int, overwrite bool) (*Machine, error) {
	machine, err := Read(os.Getenv("HOME") + "/.config/enigma.json")

	if err != nil {
		if overwrite {
			machine = Generate(numberOfRotors)
			if err = machine.Write(os.Getenv("HOME") + "/.config/enigma.json"); err != nil {
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

	return m.isInit()
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
