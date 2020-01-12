// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

// Represents an Enigma machine's components
type Machine struct {
	pathConnections      [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Connections that form electric pathways
	reflector            [ALPHABET_SIZE]int                   // Reflector connections, symmetric
	plugboardConnections [ALPHABET_SIZE]int                   // Plugboard connections, symmetric

	rotors     [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Mechanical rotors, 1st element represents rotor's current position
	takenSteps [NUMBER_OF_ROTORS - 1]int            // Number of steps taken by each rotor except the last
	step       int                                  // Size of shift between rotor steps (move)
	cycle      int                                  // Number of steps considered a full cycle, considered by following rotor when stepping
}

// Get electric pathway connections
func (m *Machine) PathConnections() [NUMBER_OF_ROTORS][ALPHABET_SIZE]int {
	return m.pathConnections
}

// SetPathConnections sets path connections array in Machine.
func (m *Machine) SetPathConnections(paths [NUMBER_OF_ROTORS][ALPHABET_SIZE]int) {
	m.pathConnections = paths
}

// Get reflector connections array
func (m *Machine) Reflector() [ALPHABET_SIZE]int {
	return m.reflector
}

// SetReflector sets reflector connections.
func (m *Machine) SetReflector(reflector [ALPHABET_SIZE]int) {
	m.reflector = reflector
}

// Get plugboard connections
func (m *Machine) PlugboardConnections() [ALPHABET_SIZE]int {
	return m.plugboardConnections
}

// SetPlugboard sets reflector connections.
func (m *Machine) SetPlugboard(plugboard [ALPHABET_SIZE]int) {
	m.plugboardConnections = plugboard
}
