package machine

import (
	"bytes"
	"strings"
	"unicode"
)

// Encrypt encrypts a string using a Machine object.
// Returns encrypted string and an error in case of an incorrect configuration.
// Uppercase and lowercase letters are treated similarly and produce the same
// results. Non-alphabetical characters are returned without change, and don't
// affect rotors' movement (rotors are not shifted).
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
	for i := 0; i < m.numberOfRotors; i++ {
		index := (encryptedChar + m.rotors[i]) % alphabetSize
		encryptedChar = m.pathConnections[i][index]
	}

	encryptedChar = m.reflector[encryptedChar]

	flipped := m.flippedConnections()
	for i := m.numberOfRotors - 1; i >= 0; i-- {
		encryptedChar = (flipped[i][encryptedChar] - m.rotors[i] + alphabetSize) % alphabetSize
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

// flippedConnections returns a slice of flipped pathway connections
// to be used in encryption cycle after reflecting.
func (m *Machine) flippedConnections() [][alphabetSize]int {
	flipped := make([][alphabetSize]int, m.numberOfRotors)
	for i, slice := range m.pathConnections {
		for j, val := range slice {
			flipped[i][val] = j
		}
	}

	return flipped
}
