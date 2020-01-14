// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

const (
	numberOfRotors = 3
	alphabetSize   = 26
)

// Machine represents an enigma machine's components
type Machine struct {
	pathConnections      [numberOfRotors][alphabetSize]int // Connections that form electric pathways
	reflector            [alphabetSize]int                 // Reflector connections, symmetric
	plugboardConnections [alphabetSize]int                 // Plugboard connections, symmetric

	rotors     [numberOfRotors][alphabetSize]int // Mechanical rotors, 1st element represents rotor's current position
	takenSteps [numberOfRotors - 1]int           // Number of steps taken by each rotor except the last
	step       int                               // Size of shift between rotor steps (move)
	cycle      int                               // Number of steps considered a full cycle, considered by following rotor when stepping
}

// PathConnections returns electric pathway connections
func (m *Machine) PathConnections() [numberOfRotors][alphabetSize]int {
	return m.pathConnections
}

// SetPathConnections sets path connections array in Machine.
func (m *Machine) SetPathConnections(paths [numberOfRotors][alphabetSize]int) {
	m.pathConnections = paths
}

// Reflector returns reflector connections array
func (m *Machine) Reflector() [alphabetSize]int {
	return m.reflector
}

// SetReflector sets reflector connections.
func (m *Machine) SetReflector(reflector [alphabetSize]int) {
	m.reflector = reflector
}

// PlugboardConnections returns plugboard connections
func (m *Machine) PlugboardConnections() [alphabetSize]int {
	return m.plugboardConnections
}

// SetPlugboard sets reflector connections.
func (m *Machine) SetPlugboard(plugboard [alphabetSize]int) {
	m.plugboardConnections = plugboard
}
