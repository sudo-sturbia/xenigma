// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"fmt"
)

// Validate and create plugboard connections
func (m *Machine) createPlugboardConnections(plugCons map[byte]byte) error {
	// Validate length
	if len(plugCons) > 13 {
		return &connectionErr{"number of connections is invalid"}
	}

	for i := 0; i < alphabetSize; i++ {
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
