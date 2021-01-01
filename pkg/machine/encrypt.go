package machine

import (
	"bytes"
	"strings"
	"unicode"
)

// Encrypt encrypts a string message, and return the encrypted string and an
// error if the machine's fields are invalid. When encrypting uppercase and
// lowercase letters produce the same results. Non-alphabetical characters are
// returned without change, and don't affect rotors' movement (rotors are not
// shifted).
func (m *Machine) Encrypt(message string) (string, error) {
	if err := m.Verify(); err != nil {
		return "", err
	}

	reversed := reverseConnections(m)
	buffer := new(bytes.Buffer)
	for _, char := range strings.ToLower(message) {
		buffer.WriteByte(m.encryptChar(byte(char), reversed))
	}
	return buffer.String(), nil
}

// encryptChar encrypts one byte using Machine m. Arguments are the byte to
// encrypt and the reversed connections to use in the reverse cycle.
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

// reverseConnections returns a reversed list of the machine's pathway
// connections.
func reverseConnections(m *Machine) [][alphabetSize]int {
	reversed := make([][alphabetSize]int, m.rotors.count)
	for i, rotor := range m.rotors.rotors {
		for j, val := range rotor.pathways {
			reversed[i][val] = j
		}
	}
	return reversed
}
