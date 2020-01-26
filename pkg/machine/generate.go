package machine

import (
	"math/rand"
	"time"
)

// Generate creates a machine object with a specified number of rotors
// containing randomly generated component configurations.
func Generate(numberOfRotors int) *Machine {
	rand.Seed(time.Now().UnixNano())

	m := new(Machine)
	m.setNumberOfRotors(numberOfRotors)

	m.randPathways()
	m.randPlugboard()
	m.randReflector()
	m.randRotors()

	return m
}

// randPathways creates random pathway connections.
func (m *Machine) randPathways() {
	m.pathConnections = make([][alphabetSize]int, m.numberOfRotors)
	for i := range m.pathConnections {
		m.pathConnections[i] = [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
			14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

		// Shuffle array elements
		rand.Shuffle(alphabetSize, func(j, k int) {
			m.pathConnections[i][j], m.pathConnections[i][k] =
				m.pathConnections[i][k], m.pathConnections[i][j]
		})
	}
}

// randPlugboard creates random plugboard connections.
func (m *Machine) randPlugboard() {
	// Create half of the connections
	halfConnections := [alphabetSize / 2]int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	rand.Shuffle(alphabetSize/2, func(i, j int) {
		halfConnections[i], halfConnections[j] =
			halfConnections[j], halfConnections[i]
	})

	for i := 0; i < len(halfConnections); i++ {
		m.plugboardConnections[i] = halfConnections[i]
		m.plugboardConnections[halfConnections[i]] = i
	}
}

// randReflector creates random reflector connections.
func (m *Machine) randReflector() {
	// Create half of the connections
	halfConnections := [alphabetSize / 2]int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	rand.Shuffle(alphabetSize/2, func(i, j int) {
		halfConnections[i], halfConnections[j] =
			halfConnections[j], halfConnections[i]
	})

	for i := 0; i < len(halfConnections); i++ {
		m.reflector[i] = halfConnections[i]
		m.reflector[halfConnections[i]] = i
	}
}

// randRotors creates random rotor positions.
func (m *Machine) randRotors() {
	rotorsPositions := make([]int, m.numberOfRotors)
	for i := range rotorsPositions {
		rotorsPositions[i] = rand.Intn(alphabetSize)
	}

	m.initRotors(rotorsPositions, 1, alphabetSize)
}
