// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"bytes"
	"strings"
	"unicode"
)

const (
	NUMBER_OF_ROTORS = 3
	ALPHABET_SIZE    = 26
)

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

// Encrypt a full message using enigma
// returns encrypted message and an error
// indicating an initialization error.
func (m *Machine) Encrypt(message string) (string, error) {
	if !m.isInit() {
		return "", &initError{"Enigma machine is not initialized correctly"}
	}

	// Create a buffer to add encrypted characters to
	message = strings.ToLower(message)
	encryptedBuffer := new(bytes.Buffer)

	for _, char := range message {
		encryptedBuffer.WriteByte(m.encryptChar(byte(char)))
	}

	return encryptedBuffer.String(), nil
}

// Encrypt a character using engima
func (m *Machine) encryptChar(char byte) byte {
	if !unicode.IsLetter(rune(char)) {
		return char
	}

	// Plugboard
	encryptedChar := m.plugIn(char)

	// Rotors and electric pathways
	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		encryptedChar = m.pathConnections[i][m.rotors[i][encryptedChar]]
	}

	// Reflector and return through electric pathways
	encryptedChar = m.reflector[encryptedChar]
	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		encryptedChar = m.rotors[i][m.pathConnections[i][encryptedChar]]
	}

	return m.plugOut(encryptedChar)
}
