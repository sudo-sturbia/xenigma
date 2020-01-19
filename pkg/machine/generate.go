package machine

import (
	"math/rand"
	"time"
)

// Generate returns a Machine object initialized using randomly generated
// configurations for components.
func Generate() *Machine {
	rand.Seed(time.Now().UnixNano())

	var m *Machine
	m.randPathways()
	m.randPlugboard()
	m.randReflector()
	m.randRotors()

	return m
}

// randPathways creates random pathway connections.
func (m *Machine) randPathways() {
	m.pathConnections[0] = [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	m.pathConnections[1] = [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	m.pathConnections[2] = [alphabetSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	for i := 0; i < numberOfRotors; i++ {
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
	var rotorsPositions [numberOfRotors]int
	for i := 0; i < numberOfRotors; i++ {
		rotorsPositions[i] = rand.Intn(alphabetSize)
	}

	m.initRotors(rotorsPositions, 1, alphabetSize)
}
