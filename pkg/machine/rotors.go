package machine

import (
	"fmt"
)

// initRotors initializes all components related to rotors.
// If incorrect values are given fields are set
// to default values and an error is returned.
func (m *Machine) initRotors(positions []int, stepSize int, cycleSize int) (err error) {
	var tempErr error

	tempErr = m.setStep(stepSize)
	if tempErr != nil {
		m.setStep(1)
		err = tempErr
	}

	tempErr = m.setCycle(cycleSize)
	if tempErr != nil {
		m.setCycle(alphabetSize)
		err = tempErr
	}

	tempErr = m.setRotorsPosition(positions)
	if tempErr != nil {
		m.resetRotors()
		err = tempErr
	}

	m.takenSteps = make([]int, m.numberOfRotors-1)
	for i := 0; i < m.numberOfRotors-1; i++ {
		m.takenSteps[i] = 0
	}

	return err
}

// setNumberOfRotors sets number of used rotors.
func (m *Machine) setNumberOfRotors(number int) {
	m.numberOfRotors = number
}

// NumberOfRotors returns number of Machine's rotors.
func (m *Machine) NumberOfRotors() int {
	return m.numberOfRotors
}

// setRotorsPosition sets current position of each rotor.
func (m *Machine) setRotorsPosition(positions []int) error {
	if len(positions) != m.numberOfRotors {
		return &initError{"number of rotors =/= number of given positions"}
	}

	// Verify positions
	for i := 0; i < m.numberOfRotors; i++ {
		if positions[i] < 0 || positions[i] > alphabetSize {
			return &initError{fmt.Sprintf("invalid position for rotor %d", i)}
		}
	}

	m.rotors = make([][alphabetSize]int, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
		for j := 0; j < alphabetSize; j++ {
			m.rotors[i][j] = (j + positions[i]) % alphabetSize
		}
	}

	return nil
}

// CurrentRotors returns current position of each rotor.
func (m *Machine) CurrentRotors() []int {
	current := make([]int, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
		current[i] = m.rotors[i][0]
	}

	return current
}

// resetRotors sets current position of each rotor to 0.
func (m *Machine) resetRotors() {
	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < m.numberOfRotors; j++ {
			m.rotors[j][i] = i
		}
	}
}

// setStep sets rotors' step size.
func (m *Machine) setStep(value int) error {
	if value <= 0 {
		return &initError{"invalid step size"}
	}

	m.step = value % alphabetSize
	return nil
}

// Step returns rotors' step size.
func (m *Machine) Step() int {
	return m.step
}

// setCycle sets size of rotor's full cycle.
// Cycle is used to signal when a following rotor should step based
// on current rotor's step count.
func (m *Machine) setCycle(value int) error {
	if value <= 0 {
		return &initError{"invalid cycle size"}
	}

	m.cycle = value
	return nil
}

// Cycle returns rotors' cycle size.
func (m *Machine) Cycle() int {
	return m.cycle
}

// stepRotors turns first rotor one step.
func (m *Machine) stepRotors() {
	for i := 0; i < m.numberOfRotors; i++ {
		// If previous rotor completed a full cycle
		if i == 0 || (m.takenSteps[i-1] == m.cycle) {
			for j := 0; j < alphabetSize; j++ {
				m.rotors[i][j] = (m.rotors[i][j] + m.step) % alphabetSize
			}

			if i != m.numberOfRotors-1 {
				m.takenSteps[i]++
			}
		}

		if i != 0 {
			m.takenSteps[i-1] %= m.cycle
		}
	}
}
