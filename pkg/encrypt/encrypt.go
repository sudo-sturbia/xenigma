// Encrypt messages using engima code
package encrypt

import ()

type machine struct {
	pathConnections      [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Connections that form electric pathways
	collector            [ALPHABET_SIZE]int                   // Collector connections, symmetric
	plugboardConnections [ALPHABET_SIZE]int                   // Plugboard connections, symmetric

	rotors [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Mechanical rotors, 1st element of each array represents rotor's current position
}

const (
	NUMBER_OF_ROTORS = 3
	ALPHABET_SIZE    = 26
)
