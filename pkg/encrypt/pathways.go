// Encrypt messages using engima code
package encrypt

import ()

const (
	NUMBER_OF_ROTORS = 3
	ALPHABET_SIZE    = 26
)

// Specify electric path connections.
//		config: specifies which configuration to use, default value: 0.
func (m *machine) createPathConnections(config int) {
	if config < 0 || config > 9 {
		config = 0
	}

	switch config {
	case 0:
		m.pathConnections[0] = [ALPHABET_SIZE]int{23, 18, 2, 11, 25, 9, 20, 5, 12, 10, 0, 13, 8, 14, 17, 3, 1, 24, 6, 15, 19, 22, 16, 4, 21, 7}
		m.pathConnections[1] = [ALPHABET_SIZE]int{20, 9, 3, 19, 2, 13, 12, 25, 24, 4, 8, 17, 22, 16, 11, 21, 1, 5, 14, 10, 18, 6, 15, 7, 23, 0}
		m.pathConnections[2] = [ALPHABET_SIZE]int{16, 0, 21, 5, 3, 1, 25, 9, 15, 13, 7, 11, 10, 14, 20, 24, 19, 6, 17, 2, 22, 23, 4, 18, 8, 12}
	case 1:
		m.pathConnections[0] = [ALPHABET_SIZE]int{14, 25, 23, 3, 4, 8, 24, 21, 13, 22, 9, 10, 19, 16, 12, 11, 20, 17, 1, 7, 0, 2, 15, 6, 5, 18}
		m.pathConnections[1] = [ALPHABET_SIZE]int{4, 2, 0, 7, 1, 23, 9, 6, 10, 13, 22, 11, 20, 3, 25, 5, 17, 24, 18, 8, 14, 16, 15, 19, 21, 12}
		m.pathConnections[2] = [ALPHABET_SIZE]int{19, 13, 15, 23, 7, 16, 12, 25, 14, 24, 21, 2, 5, 9, 11, 18, 4, 1, 3, 6, 10, 0, 22, 8, 17, 20}
	case 2:
		m.pathConnections[0] = [ALPHABET_SIZE]int{23, 24, 22, 4, 5, 1, 13, 21, 7, 11, 18, 10, 17, 2, 6, 16, 14, 8, 0, 12, 19, 3, 15, 20, 9, 25}
		m.pathConnections[1] = [ALPHABET_SIZE]int{7, 6, 19, 23, 0, 3, 18, 25, 4, 10, 15, 17, 20, 21, 14, 11, 5, 13, 1, 9, 2, 12, 24, 22, 8, 16}
		m.pathConnections[2] = [ALPHABET_SIZE]int{3, 8, 17, 22, 6, 19, 13, 4, 0, 21, 24, 5, 12, 10, 2, 1, 7, 15, 14, 16, 23, 20, 18, 11, 9, 25}
	case 3:
		m.pathConnections[0] = [ALPHABET_SIZE]int{19, 11, 12, 3, 8, 16, 1, 17, 23, 13, 24, 6, 15, 4, 22, 20, 25, 18, 10, 2, 0, 7, 14, 5, 21, 9}
		m.pathConnections[1] = [ALPHABET_SIZE]int{6, 0, 8, 24, 21, 17, 20, 25, 12, 1, 9, 16, 19, 23, 15, 5, 7, 18, 11, 4, 10, 3, 13, 14, 2, 22}
		m.pathConnections[2] = [ALPHABET_SIZE]int{23, 19, 10, 6, 9, 11, 15, 17, 12, 14, 5, 24, 0, 2, 4, 16, 21, 3, 20, 22, 1, 18, 13, 8, 25, 7}
	case 4:
		m.pathConnections[0] = [ALPHABET_SIZE]int{1, 16, 21, 4, 2, 24, 9, 8, 0, 20, 7, 13, 22, 14, 5, 25, 10, 11, 12, 3, 15, 19, 23, 18, 17, 6}
		m.pathConnections[1] = [ALPHABET_SIZE]int{9, 8, 13, 23, 25, 12, 0, 21, 7, 5, 19, 22, 16, 10, 20, 3, 15, 1, 2, 4, 18, 17, 14, 11, 6, 24}
		m.pathConnections[2] = [ALPHABET_SIZE]int{22, 10, 25, 13, 8, 14, 20, 21, 24, 9, 12, 4, 0, 23, 15, 17, 2, 3, 1, 19, 18, 5, 7, 11, 6, 16}
	case 5:
		m.pathConnections[0] = [ALPHABET_SIZE]int{0, 19, 24, 25, 16, 1, 20, 13, 18, 9, 6, 14, 3, 5, 4, 12, 23, 8, 15, 22, 17, 2, 21, 7, 10, 11}
		m.pathConnections[1] = [ALPHABET_SIZE]int{25, 23, 22, 1, 18, 11, 19, 12, 5, 17, 7, 13, 24, 20, 8, 9, 2, 14, 16, 21, 6, 4, 0, 10, 3, 15}
		m.pathConnections[2] = [ALPHABET_SIZE]int{23, 11, 9, 0, 16, 14, 5, 21, 24, 10, 12, 6, 2, 18, 17, 4, 1, 19, 8, 22, 7, 20, 3, 15, 13, 25}
	case 6:
		m.pathConnections[0] = [ALPHABET_SIZE]int{1, 3, 24, 22, 17, 16, 10, 21, 13, 14, 19, 0, 23, 15, 12, 5, 18, 7, 2, 25, 8, 9, 4, 11, 20, 6}
		m.pathConnections[1] = [ALPHABET_SIZE]int{25, 21, 8, 14, 1, 6, 18, 15, 19, 20, 12, 10, 3, 23, 7, 4, 13, 5, 0, 9, 16, 2, 22, 24, 11, 17}
		m.pathConnections[2] = [ALPHABET_SIZE]int{8, 1, 25, 3, 21, 15, 22, 7, 19, 14, 18, 0, 16, 6, 10, 11, 24, 17, 23, 4, 2, 12, 20, 5, 13, 9}
	case 7:
		m.pathConnections[0] = [ALPHABET_SIZE]int{16, 0, 15, 7, 9, 18, 24, 14, 17, 5, 13, 6, 22, 12, 20, 21, 23, 2, 25, 10, 4, 1, 3, 19, 8, 11}
		m.pathConnections[1] = [ALPHABET_SIZE]int{1, 18, 5, 14, 10, 15, 6, 25, 0, 13, 23, 11, 9, 22, 21, 3, 8, 19, 20, 12, 4, 17, 7, 24, 2, 16}
		m.pathConnections[2] = [ALPHABET_SIZE]int{16, 2, 18, 23, 7, 6, 25, 22, 24, 21, 13, 11, 1, 20, 12, 10, 8, 19, 0, 3, 5, 4, 14, 9, 15, 17}
	case 8:
		m.pathConnections[0] = [ALPHABET_SIZE]int{11, 14, 4, 15, 16, 18, 8, 7, 17, 5, 12, 6, 13, 0, 22, 9, 19, 3, 25, 23, 24, 2, 20, 1, 10, 21}
		m.pathConnections[1] = [ALPHABET_SIZE]int{8, 15, 13, 23, 14, 6, 9, 21, 16, 5, 19, 17, 10, 2, 11, 1, 20, 18, 22, 4, 3, 12, 25, 24, 0, 7}
		m.pathConnections[2] = [ALPHABET_SIZE]int{18, 6, 0, 21, 3, 12, 8, 11, 19, 4, 7, 17, 14, 20, 25, 24, 23, 1, 9, 5, 10, 15, 22, 16, 13, 2}
	case 9:
		m.pathConnections[0] = [ALPHABET_SIZE]int{20, 13, 17, 7, 8, 5, 25, 4, 6, 15, 22, 3, 9, 11, 24, 10, 18, 19, 23, 12, 1, 2, 21, 14, 16, 0}
		m.pathConnections[1] = [ALPHABET_SIZE]int{13, 16, 8, 5, 21, 9, 19, 2, 20, 24, 7, 22, 6, 23, 15, 10, 11, 12, 25, 1, 3, 17, 14, 18, 0, 4}
		m.pathConnections[2] = [ALPHABET_SIZE]int{19, 20, 14, 0, 7, 24, 23, 12, 15, 21, 8, 5, 2, 11, 16, 18, 13, 4, 10, 22, 25, 1, 17, 3, 6, 9}
	}
}
