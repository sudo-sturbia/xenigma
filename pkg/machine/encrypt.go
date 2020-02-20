package machine

import (
	"bytes"
	"strings"
	"unicode"
)

// Encrypt encrypts a string using a Machine object.
// Returns encrypted string and an error in case of an incorrect configuration.
// When encrypting uppercase and lowercase letters are treated similarly and
// produce the same results. Non-alphabetical characters are returned without
// change, and don't affect rotors' movement (rotors are not shifted).
func (m *Machine) Encrypt(message string) (string, error) {
	if err := m.IsConfigCorrect(); err != nil {
		return "", err
	}

	encryptedBuffer := new(bytes.Buffer)

	flipped := m.flippedConnections()

	message = strings.ToLower(message)
	for _, char := range message {
		encryptedBuffer.WriteByte(m.encryptChar(byte(char), flipped))
	}

	return encryptedBuffer.String(), nil
}

// encryptChar encrypts a character using machine.
// Takes character to encrypt and flipped pathways for usage in the
// second half of the encryption cycle.
func (m *Machine) encryptChar(char byte, flipped [][alphabetSize]int) byte {
	if !unicode.IsLetter(rune(char)) {
		return char
	}

	encrypted := m.plugboard.plugIn(char)

	for i := 0; i < m.numberOfRotors; i++ {
		index := (encrypted + m.rotors[i].position) % alphabetSize
		encrypted = m.rotors[i].pathways[index]
	}

	encrypted = m.reflector.reflect(encrypted)

	for i := m.numberOfRotors - 1; i >= 0; i-- {
		encrypted = (flipped[i][encrypted] - m.rotors[i].position + alphabetSize) % alphabetSize
	}

	m.stepRotors()

	return m.plugboard.plugOut(encrypted)
}

// flippedConnections returns a slice of flipped pathway connections
// to be used in encryption cycle after reflecting.
func (m *Machine) flippedConnections() [][alphabetSize]int {
	flipped := make([][alphabetSize]int, m.numberOfRotors)

	for i, rotor := range m.rotors {
		for j, val := range rotor.pathways {
			flipped[i][val] = j
		}
	}

	return flipped
}
