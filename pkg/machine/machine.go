/*
Package machine represents an enigma machine used for encryption and
decryption of text messages.

The machine consists of the following components: electric pathways,
reflector, plugboard, and 3 rotors.

A machine can be created using machine.Load, machine.Generate, or by
simply creating a Machine object and calling SetComponents method.

machine.Load is meant for usage in an independent program. The other
two options are more suitable for usage when the package is imported.

Encrypt method is used to encrypt (or decrypt) a given message as
the following example
	m := machine.Generate()
	encrypted := m.Encrypt("message")

Licensed under MIT license @github.com/sudo-sturbia
*/
package machine

import (
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

// Load returns a fully initialized Machine object. Configurations of
// Machine's fields are read from config file $HOME/.config/enigma.json
// If the file contains correct configurations, a machine object is
// initialized and returned with error being nil.
// Otherwise overwrite parameter is checked. If overwrite is true, random
// configs are generated and written to file, a machine object with the
// same configs is returned and error is nil. Otherwise an initialization
// error is returned and Machine is nil.
func Load(overwrite bool) (*Machine, error) {
	machine, err := read(os.Getenv("HOME") + "/.config/enigma.json")

	if err != nil {
		if overwrite {
			machine = Generate()
			if err = write(machine, os.Getenv("HOME")+"/.config/enigma.json"); err != nil {
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
	pathways [numberOfRotors][alphabetSize]int,
	plugboard [alphabetSize]int,
	reflector [alphabetSize]int,
	rotorsPositions [numberOfRotors]int, step int, cycle int) error {

	m.setPathConnections(pathways)
	m.setPlugboard(plugboard)
	m.setReflector(reflector)
	m.initRotors(rotorsPositions, step, cycle)

	return m.isInit()
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
