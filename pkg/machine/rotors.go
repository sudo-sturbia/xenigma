package machine

import (
	"fmt"
	"math"
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

// setTakenSteps calculates the number of taken steps for each
// of the rotors except the last and populates the takenSteps array
// with the calculated values.
// An error is returned if the given position of the rotors can't
// be reached (takenSteps can't be calculated).
func (m *Machine) setTakenSteps(position []int) error {
	for i := 0; i < m.numberOfRotors-1; i++ {
		steps := 0
		for j := i; j < m.numberOfRotors; j++ {
			steps += position[j] * int(math.Pow(float64(alphabetSize), float64(j)))
		}

		if (steps % m.step) != 0 {
			return &initError{"given rotor position is incorrect"}
		}

		m.takenSteps[i] = (steps / m.step) % m.cycle
	}

	return nil
}

// verifyStepCycle verifies that given properties of the rotors are
// correct. Verifications are that a full machine cycle can be achieved
// using given step (full machine cycle breaks into a whole number of steps).
func (m *Machine) verifyStepCycle(stepSize int, cycleSize int) error {
	if (int(math.Pow(float64(alphabetSize), float64(m.numberOfRotors))) % (stepSize * cycleSize)) != 0 {
		return &initError{"a full machine cycle can not be broken into a whole number of steps."}
	}

	return nil
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

// setNumberOfRotors sets number of used rotors.
func (m *Machine) setNumberOfRotors(number int) {
	m.numberOfRotors = number
}

// NumberOfRotors returns number of Machine's rotors.
func (m *Machine) NumberOfRotors() int {
	return m.numberOfRotors
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
