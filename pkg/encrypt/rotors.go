// Encrypt messages using engima code
package encrypt

import ()

// Set initial rotors' position
func (m *machine) setInitialRotors() {
	for i := 0; i < ALPHABET_SIZE; i++ {
		m.rotors[0][i] = i
		m.rotors[1][i] = i
		m.rotors[2][i] = i
	}
}
