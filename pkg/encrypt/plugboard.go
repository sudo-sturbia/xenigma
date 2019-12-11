// Encrypt messages using engima code
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
		// If character was mapped before
		if m.plugboardConnections[i] != 0 {
			continue
		}

		// If character is not alphabetical and lowercase
		if plugCons[byte(i)] < 97 || plugCons[byte(i)] > 122 {
			return &connectionErr{fmt.Sprintf("input character '%v' is incorrect", plugCons[byte(i)])}
		}

		m.plugboardConnections[i] = int(plugCons[byte(i)])
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
	return byte(m.plugboardConnections[int(char-'a')] + 'a')
}
