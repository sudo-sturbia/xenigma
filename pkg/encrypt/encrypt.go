// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"bytes"
	"strings"
	"unicode"
)

const (
	// NumberOfRotors is the number of rotors used in machine
	NumberOfRotors = 3

	alphabetSize = 26
)

// Encrypt encrypts a string using enigma.
// returns encrypted string and an error incase of an initialization error.
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
	for i := 0; i < NumberOfRotors; i++ {
		encryptedChar = m.pathConnections[i][m.rotors[i][encryptedChar]]
	}

	// Reflector and return through electric pathways
	encryptedChar = m.reflector[encryptedChar]
	for i := 0; i < NumberOfRotors; i++ {
		encryptedChar = m.rotors[i][m.pathConnections[i][encryptedChar]]
	}

	return m.plugOut(encryptedChar)
}

// Change byte (character) to an int (0 -> 25) based on plugboard connections
// Used when character is entered
func (m *Machine) plugIn(char byte) int {
	return int(m.plugboardConnections[int(char-'a')])
}

// Change int to a byte (character) based on plugboard connections
// Used when character is returned
func (m *Machine) plugOut(char int) byte {
	return byte(m.plugboardConnections[char]) + 'a'
}
