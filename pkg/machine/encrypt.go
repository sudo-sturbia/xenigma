package machine

import (
	"bytes"
	"strings"
	"unicode"
)

// Encrypt encrypts a string using a Machine object.
// returns encrypted string and an error incase of an initialization error.
func (m *Machine) Encrypt(message string) (string, error) {
	if err := m.isInit(); err != nil {
		return "", err
	}

	encryptedBuffer := new(bytes.Buffer)

	message = strings.ToLower(message)
	for _, char := range message {
		encryptedBuffer.WriteByte(m.encryptChar(byte(char)))
	}

	return encryptedBuffer.String(), nil
}

// encryptChar encrypts a character using machine.
func (m *Machine) encryptChar(char byte) byte {
	if !unicode.IsLetter(rune(char)) {
		return char
	}

	// Plugboard
	encryptedChar := m.plugIn(char)

	// Rotors and electric pathways
	for i := 0; i < numberOfRotors; i++ {
		encryptedChar = m.pathConnections[i][m.rotors[i][encryptedChar]]
	}

	// Reflector and return through electric pathways
	encryptedChar = m.reflector[encryptedChar]
	for i := 0; i < numberOfRotors; i++ {
		encryptedChar = m.rotors[i][m.pathConnections[i][encryptedChar]]
	}

	return m.plugOut(encryptedChar)
}

// plugIn changes a byte (character) to an int (0 -> 25) based on
// plugboard connections. Used when character is entered.
func (m *Machine) plugIn(char byte) int {
	return m.plugboardConnections[int(char-'a')]
}

// plugOut changes an int to a byte (character) based on
// plugboard connections. Used when character is returned.
func (m *Machine) plugOut(char int) byte {
	return byte(m.plugboardConnections[char] + 'a')
}
