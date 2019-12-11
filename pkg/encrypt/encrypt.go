// Encrypt messages using engima code
package encrypt

import ()

type machine struct {
	pathConnections [NUMBER_OF_ROTORS][ALPHABET_SIZE]int // Connections that form electric pathways
	collector       [ALPHABET_SIZE]int                   // Collector connections, symmetric

	plugboardConnections map[byte]byte // Plugboard connections, symmetric
}
