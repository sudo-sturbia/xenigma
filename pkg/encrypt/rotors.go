// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"fmt"
)

// Initialize all components related to rotors.
// If incorrect values are given fields are set
// to default and an error is returned.
func (m *Machine) initRotors(positions [NUMBER_OF_ROTORS]int, stepSize int, cycleSize int) (err error) {
	var tempErr error

	tempErr = m.setStep(stepSize)
	if tempErr != nil {
		m.setStep(1)
		err = tempErr
	}

	tempErr = m.setCycle(cycleSize)
	if tempErr != nil {
		m.setCycle(ALPHABET_SIZE)
		err = tempErr
	}

	tempErr = m.setRotorsPosition(positions)
	if tempErr != nil {
		m.resetRotors()
		err = tempErr
	}

	for i := 0; i < NUMBER_OF_ROTORS-1; i++ {
		m.takenSteps[i] = 0
	}

	return err
}

// Set position of each rotor
func (m *Machine) setRotorsPosition(positions [NUMBER_OF_ROTORS]int) error {
	// Verify positions
	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		if positions[i] < 0 || positions[i] > ALPHABET_SIZE {
			return &connectionErr{fmt.Sprintf("invalid position for rotor %d", i)}
		}
	}

	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		for j := 0; j < ALPHABET_SIZE; j++ {
			m.rotors[i][j] = (j + positions[i]) % ALPHABET_SIZE
		}
	}

	return nil
}

// Get current position of rotors
func (m *Machine) CurrentRotors() [NUMBER_OF_ROTORS]int {
	return [NUMBER_OF_ROTORS]int{
		m.rotors[0][0], m.rotors[1][0], m.rotors[2][0],
	}
}

// Reset all rotors (set positions to zero)
func (m *Machine) resetRotors() {
	for i := 0; i < ALPHABET_SIZE; i++ {
		m.rotors[0][i] = i
		m.rotors[1][i] = i
		m.rotors[2][i] = i
	}
}

// Set value of rotors step
func (m *Machine) setStep(value int) error {
	if value <= 0 {
		return &rotorConfigErr{"invalid step value"}
	}

	m.step = value % ALPHABET_SIZE
	return nil
}

// Get step value
func (m *Machine) Step() int {
	return m.step
}

// Set size of rotor's full cycle.
// Indicates number of steps considered a full rotor cycle.
// Used to signal when a following rotor should step based
// on current rotor's step count.
func (m *Machine) setCycle(value int) error {
	if value <= 0 {
		return &rotorConfigErr{"invalid cycle size"}
	}

	m.cycle = value
	return nil
}

// Get cycle value
func (m *Machine) Cycle() int {
	return m.cycle
}

// Turn rotors one step
func (m *Machine) stepRotors() {
	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		// If previous rotor completed a full cycle
		if i == 0 || (m.takenSteps[i-1] == m.cycle) {
			for j := 0; j < ALPHABET_SIZE; j++ {
				m.rotors[i][j] = (m.rotors[i][j] + m.step) % ALPHABET_SIZE
			}

			if i != NUMBER_OF_ROTORS-1 {
				m.takenSteps[i] += 1
			}
		}

		if i != 0 {
			m.takenSteps[i-1] %= m.cycle
		}
	}
}
