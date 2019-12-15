// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"fmt"
)

// Validate and create plugboard connections
func (m *machine) createPlugboardConnections(plugCons map[byte]byte) error {
	// Validate length
	if len(plugCons) > 13 {
		return &connectionErr{"number of connections is invalid"}
	}

	for i := 0; i < ALPHABET_SIZE; i++ {
		m.plugboardConnections[i] = i
	}

	for key, value := range plugCons {
		if m.plugboardConnections[int(key-'a')] != int(key-'a') || m.plugboardConnections[int(value-'a')] != int(value-'a') {
			return &connectionErr{fmt.Sprintf("characters '%c' and '%c' are mapped several times", key, value)}
		}

		if key < 97 || key > 122 || value < 97 || value > 122 {
			return &connectionErr{"mappings contain non alphabetical characters"}
		}

		m.plugboardConnections[int(key-'a')] = int(value - 'a')
		m.plugboardConnections[int(value-'a')] = int(key - 'a')
	}

	return nil
}

// Change byte (character) to an int (0 -> 25) based on plugboard connections
// Used when character is entered
func (m *machine) plugIn(char byte) int {
	return int(m.plugboardConnections[int(char-'a')])
}

// Change int to a byte (character) based on plugboard connections
// Used when character is returned
func (m *machine) plugOut(char int) byte {
	return byte(m.plugboardConnections[char]) + 'a'
}
