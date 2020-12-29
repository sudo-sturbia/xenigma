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
	if err := m.Verify(); err != nil {
		return "", err
	}

	reversed := m.reverseConnections()

	buffer := new(bytes.Buffer)
	for _, char := range strings.ToLower(message) {
		buffer.WriteByte(m.encryptChar(byte(char), reversed))
	}
	return buffer.String(), nil
}

// encryptChar encrypts a character using machine.
// Takes character to encrypt and flipped pathways for usage in the
// second half of the encryption cycle.
func (m *Machine) encryptChar(char byte, reversed [][alphabetSize]int) byte {
	if !unicode.IsLetter(rune(char)) {
		return char
	}

	encrypted := m.plugboard.PlugIn(char)
	for i := 0; i < m.rotors.count; i++ {
		index := (encrypted + m.rotors.rotors[i].position) % alphabetSize
		encrypted = m.rotors.rotors[i].pathways[index]
	}

	encrypted = m.reflector.Reflect(encrypted)
	for i := m.rotors.count - 1; i >= 0; i-- {
		encrypted = (reversed[i][encrypted] - m.rotors.rotors[i].position + alphabetSize) % alphabetSize
	}
	m.rotors.takeStep()

	return m.plugboard.PlugOut(encrypted)
}

// flippedConnections returns a slice of flipped pathway connections
// to be used in encryption cycle after reflecting.

func (m *Machine) reverseConnections() [][alphabetSize]int {
	reversed := make([][alphabetSize]int, m.rotors.count)
	for i, rotor := range m.rotors.rotors {
		for j, val := range rotor.pathways {
			reversed[i][val] = j
		}
	}
	return reversed
}
