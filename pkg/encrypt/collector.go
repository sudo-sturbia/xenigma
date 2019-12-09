// Encrypt messages using engima code
package encrypt

import ()

var collector [ALPHABET_SIZE]int

// Create collector connections
func createCollectorConnections() {
	halfConnections := [13]int{15, 19, 17, 25, 24, 14, 23, 18, 16, 22, 20, 21, 13}

	for i := 0; i < len(halfConnections); i++ {
		collector[i] = halfConnections[i]
		collector[halfConnections[i]] = i
	}
}
