// Encrypt messages using engima code
package encrypt

import (
	"fmt"
)

// Set initial rotors' position
func (m *machine) setInitialRotors() {
	for i := 0; i < ALPHABET_SIZE; i++ {
		m.rotors[0][i] = i
		m.rotors[1][i] = i
		m.rotors[2][i] = i
	}
}

// Set position of each rotor
func (m *machine) setRotorsPosition(positions [NUMBER_OF_ROTORS]int) error {
	// Verify positions
	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		if positions[i] < 0 || positions[i] > ALPHABET_SIZE {
			return &connectionErr{fmt.Sprintf("invalid position for rotor %d", i)}
		}
	}

	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		for j := 0; j < ALPHABET_SIZE; j++ {
			m.rotors[i][j] = (i + positions[i]) % ALPHABET_SIZE
		}
	}

	return nil
}

// Set value of rotors step
func (m *machine) setStep(value int) error {
	if value <= 0 {
		return &rotorConfigErr{"invalid step value"}
	}

	m.step = value
	return nil
}

// Set size of rotor's full cycle.
// Indicates number of steps considered a full rotor cycle.
// Used to signal when a following rotor should step based
// on current rotor's step count.
func (m *machine) setFullCycle(value int) error {
	if value <= 0 {
		return &rotorConfigErr{"invalid cycle size"}
	}

	m.fullCycle = value
	return nil
}

// Turn rotors one step
func (m *machine) stepRotors() {
	previousRotorPos := m.rotors[0][0] % m.fullCycle // Old position of previous rotor (before step)

	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		tempPos := m.rotors[i][0] % m.fullCycle

		// Previous rotor resetted
		if i == 0 || (m.rotors[i-1][0] < previousRotorPos) {
			for j := 0; j < ALPHABET_SIZE; j++ {
				m.rotors[i][j] = (m.rotors[i][j] + 1) % ALPHABET_SIZE
			}
		}

		previousRotorPos = tempPos
	}
}
