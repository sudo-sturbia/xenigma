package machine

import (
	"fmt"
)

// initRotors initializes all components related to rotors.
// If incorrect values are given fields are set to default
// and an error is returned.
func (m *Machine) initRotors(positions []int, stepSize int, cycleSize int) (err error) {
	if tempErr := m.setStepAndCycle(stepSize, cycleSize); tempErr != nil {
		err = tempErr
	}

	if tempErr := m.setRotorsPosition(positions); tempErr != nil {
		err = tempErr
	}

	if tempErr := m.setTakenSteps(positions); tempErr != nil {
		m.resetRotors()

		err = tempErr
	}

	return err
}

// setRotorsPosition sets current position of each rotor.
// If given positions are invalid all positions are set to
// default value (starting at 0).
func (m *Machine) setRotorsPosition(positions []int) error {
	m.rotors = make([][alphabetSize]int, m.numberOfRotors)

	if len(positions) != m.numberOfRotors {
		m.resetRotors()
		return &initError{"number of rotors =/= number of given positions"}
	}

	// Verify positions
	for i := 0; i < m.numberOfRotors; i++ {
		if positions[i] < 0 || positions[i] > alphabetSize {
			m.resetRotors()
			return &initError{fmt.Sprintf("invalid position for rotor %d", i)}
		}
	}

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
// be reached using specified step and cycle values (takenSteps can't
// be calculated) and taken steps are set to 0 for all rotors.
func (m *Machine) setTakenSteps(position []int) error {
	m.takenSteps = make([]int, m.numberOfRotors-1)
	for i := 0; i < m.numberOfRotors; i++ {
		if (position[i] % m.step) != 0 {
			m.resetTakenSteps()
			return &initError{"given position of rotors is incorrect"}
		}
	}

	for i := 0; i < m.numberOfRotors-1; i++ {
		m.takenSteps[i] = (position[i] / m.step) % m.cycle
	}

	return nil
}

// resetTakenSteps sets taken steps for all rotors to 0.
func (m *Machine) resetTakenSteps() {
	for i := 0; i < m.numberOfRotors-1; i++ {
		m.takenSteps[i] = 0
	}
}

// setStepAndCycle verifies and sets both step size and cycle size.
// If given values are incorrect or incompatibe both fields are set
// to default. Step = 1, Cycle = 26.
func (m *Machine) setStepAndCycle(stepSize int, cycleSize int) error {
	if err := m.verifyStepCycle(stepSize, cycleSize); err != nil {
		m.setStep(DefaultStep)
		m.setCycle(DefaultCycle)
		return err
	}

	err1 := m.setStep(stepSize)
	err2 := m.setCycle(cycleSize)

	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return err2
	}

	return nil
}

// verifyStepCycle verifies that a full machine cycle can be achieved
// using given step (full machine cycle breaks into a whole number of steps).
func (m *Machine) verifyStepCycle(stepSize int, cycleSize int) error {
	if ((alphabetSize) % (stepSize * cycleSize)) != 0 {
		return &initError{"cycle size and step size are not compatible, some collisions may occur"}
	}

	return nil
}

// setStep sets rotors' step size.
// If given value is invalid step size is set to 1.
func (m *Machine) setStep(value int) error {
	if value <= 0 {
		m.step = DefaultStep
		return &initError{"invalid step size"}
	}

	m.step = value % alphabetSize
	return nil
}

// Step returns rotors' step size. Step represents the number of positions
// a rotor jumps when taking one step. The default size of a step is 1.
func (m *Machine) Step() int {
	return m.step
}

// setCycle sets size of rotor's full cycle.
// If given value is invalid cycle size is set to 26.
func (m *Machine) setCycle(value int) error {
	if value <= 0 {
		m.cycle = DefaultCycle
		return &initError{"invalid cycle size"}
	}

	m.cycle = value
	return nil
}

// Cycle returns rotors' cycle size. Cycle is the number of steps that
// represent a rotor's full cycle. Cycle is used to indicate when
// rotor should step (move) based on the number of steps taken by preceding
// rotor. The default size of a cycle is 26 (alphabet size).
func (m *Machine) Cycle() int {
	return m.cycle
}

// setNumberOfRotors sets number of used rotors.
func (m *Machine) setNumberOfRotors(value int) error {
	if value <= 0 {
		return &initError{"invalid number of rotors"}
	}

	m.numberOfRotors = value
	return nil
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
