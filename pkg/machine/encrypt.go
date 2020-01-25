package machine

import (
	"bytes"
	"strings"
	"unicode"
)

// Encrypt encrypts a string using a Machine object.
// returns encrypted string and an error in case of an incorrect configuration.
// Non-alphabetical characters are returned without change, and don't affect
// rotors' movement (rotors are not shifted).
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

	encryptedChar := m.plugIn(char)
	for i := 0; i < numberOfRotors; i++ {
		encryptedChar = m.pathConnections[i][m.rotors[i][encryptedChar]]
	}

	encryptedChar = m.reflector[encryptedChar]
	for i := 0; i < numberOfRotors; i++ {
		encryptedChar = m.rotors[numberOfRotors-i-1][m.pathConnections[numberOfRotors-i-1][encryptedChar]]
	}

	m.stepRotors()

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
