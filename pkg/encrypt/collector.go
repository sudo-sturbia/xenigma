// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import ()

// Create collector connections
//		config: specifies which configuration to use, default value: 0.
func (m *machine) createCollectorConnections(config int) {
	if config < 0 || config > 9 {
		config = 0
	}

	var halfConnections [13]int
	switch config {
	case 0:
		halfConnections = [13]int{15, 19, 17, 25, 24, 14, 23, 18, 16, 22, 20, 21, 13}
	case 1:
		halfConnections = [13]int{14, 21, 18, 19, 24, 23, 22, 20, 15, 16, 25, 17, 13}
	case 2:
		halfConnections = [13]int{17, 18, 14, 24, 23, 22, 15, 19, 25, 20, 21, 16, 13}
	case 3:
		halfConnections = [13]int{13, 17, 14, 24, 23, 21, 20, 15, 18, 25, 19, 22, 16}
	case 4:
		halfConnections = [13]int{22, 15, 24, 23, 13, 21, 20, 19, 17, 25, 18, 14, 16}
	case 5:
		halfConnections = [13]int{21, 16, 17, 23, 13, 19, 15, 25, 20, 24, 18, 14, 22}
	case 6:
		halfConnections = [13]int{19, 20, 14, 23, 13, 15, 24, 16, 18, 25, 22, 17, 21}
	case 7:
		halfConnections = [13]int{14, 25, 22, 20, 17, 24, 13, 19, 15, 21, 23, 18, 16}
	case 8:
		halfConnections = [13]int{16, 22, 25, 18, 14, 21, 23, 24, 13, 15, 20, 19, 17}
	case 9:
		halfConnections = [13]int{17, 21, 25, 19, 18, 20, 14, 13, 24, 16, 22, 23, 15}
	}

	for i := 0; i < len(halfConnections); i++ {
		m.collector[i] = halfConnections[i]
		m.collector[halfConnections[i]] = i
	}
}
