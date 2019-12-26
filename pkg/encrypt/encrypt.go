// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import ()

const (
	NUMBER_OF_ROTORS = 3
	ALPHABET_SIZE    = 26
)

// Represents an Enigma machine's components
type Machine struct {
	pathConnections      [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Connections that form electric pathways
	collector            [ALPHABET_SIZE]int                   // Collector connections, symmetric
	plugboardConnections [ALPHABET_SIZE]int                   // Plugboard connections, symmetric

	rotors     [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Mechanical rotors, 1st element represents rotor's current position
	takenSteps [NUMBER_OF_ROTORS - 1]int            // Number of steps taken by each rotor except the last
	step       int                                  // Size of shift between rotor steps (move)
	cycle      int                                  // Number of steps considered a full cycle, considered by following rotor when stepping
}

// Encrypt a character using engima
// Returns encrypted character and an error indicating a machine configuration problem.
func (m *Machine) encryptChar(char byte) (byte, error) {
	if !m.isInit() {
		return ' ', &initError{"Enigma machine is not initialized correctly"}
	}

	// ...
	return ' ', nil
}
