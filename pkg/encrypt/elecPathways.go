// Encrypt messages using engima code
package encrypt

import ()

const NUMBER_OF_ROTORS = 3;
const ALPHABET_SIZE = 26;

var elecPathConnections [NUMBER_OF_ROTORS][ALPHABET_SIZE]byte // Connections that form electric pathways

// Specify electric path connections
func createPathConnections() {
	elecPathConnections[0] = {23, 18, 2, 11, 25, 9, 20, 5, 12, 10, 0, 13, 8, 14, 17, 3, 1, 24, 6, 15, 19, 22, 16, 4, 21, 7}
	elecPathConnections[1] = {20, 9, 3, 19, 2, 13, 12, 25, 24, 4, 8, 17, 22, 16, 11, 21, 1, 5, 14, 10, 18, 6, 15, 7, 23, 0}
	elecPathConnections[2] = {16, 0, 21, 5, 3, 1, 25, 9, 15, 13, 7, 11, 10, 14, 20, 24, 19, 6, 17, 2, 22, 23, 4, 18, 8, 12}
}

