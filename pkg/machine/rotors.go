package machine

import (
	"fmt"
)

// initRotors initializes all components related to rotors.
// If incorrect values are given an error is returned.
func (m *Machine) initRotors(positions []int, stepSize int, cycleSize int) error {
	if err := m.setStepAndCycle(stepSize, cycleSize); err != nil {
		return err
	}

	if err := m.setRotorsPosition(positions); err != nil {
		return err
	}

	return nil
}

// UseRotorDefaults use default value for rotor related components.
// Defaults are a's for rotors' positions, 1 for step size, and 26
// (size of the alphabet) for cycle size.
func (m *Machine) UseRotorDefaults() {
	m.resetRotorsPosition()

	m.setStep(DefaultStep)
	m.setCycle(DefaultCycle)
}

// setRotorsPosition sets current position of each rotor and calculates
// number of taken steps. If given positions are invalid an error is returned.
func (m *Machine) setRotorsPosition(positions []int) error {
	if err := m.arePositionsValid(positions); err != nil {
		return err
	}

	m.rotors = make([][alphabetSize]int, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
		for j := 0; j < alphabetSize; j++ {
			m.rotors[i][j] = (j + positions[i]) % alphabetSize
		}
	}

	m.setTakenSteps(positions)

	return nil
}

// resetRotorsPosition sets current position of each rotor and number
// of taken steps to 0.
func (m *Machine) resetRotorsPosition() {
	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < m.numberOfRotors; j++ {
			m.rotors[j][i] = i
		}
	}

	m.resetTakenSteps()
}

// setTakenSteps calculates the number of taken steps for each
// of the rotors except the last and populates the takenSteps array
// with the calculated values.
func (m *Machine) setTakenSteps(positions []int) {
	m.takenSteps = make([]int, m.numberOfRotors-1)
	for i := 0; i < m.numberOfRotors-1; i++ {
		m.takenSteps[i] = (positions[i] / m.step) % m.cycle
	}
}

// resetTakenSteps sets taken steps for all rotors to 0.
func (m *Machine) resetTakenSteps() {
	for i := 0; i < m.numberOfRotors-1; i++ {
		m.takenSteps[i] = 0
	}
}

// arePositionsValid verifies given rotor positions. Returns an error
// if given values are invalid, nil otherwise.
func (m *Machine) arePositionsValid(positions []int) error {
	if len(positions) != m.numberOfRotors {
		return &initError{"number of rotors =/= number of given positions"}
	}

	for i := 0; i < m.numberOfRotors; i++ {
		if positions[i] < 0 || positions[i] > alphabetSize {
			return &initError{fmt.Sprintf("invalid position for rotor %d", i)}
		}

		if (positions[i] % m.step) != 0 {
			return &initError{"given position of rotors is incorrect"}
		}
	}

	return nil
}

// CurrentRotors returns a slice containing current position of each rotor.
func (m *Machine) CurrentRotors() []int {
	current := make([]int, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
		current[i] = m.rotors[i][0]
	}

	return current
}

// setStepAndCycle verifies and sets both step size and cycle size.
// If given values are incorrect or incompatibe an error is returned.
func (m *Machine) setStepAndCycle(stepSize int, cycleSize int) error {
	if err := m.areStepCycleValid(stepSize, cycleSize); err != nil {
		return err
	}

	m.setStep(stepSize)
	m.setCycle(cycleSize)

	return nil
}

// areStepCycleValid verifies that both step and cycle sizes are positive
// and that the given step-cycle combination produces no position collisions
// when used (a position can not be achieved using several step sequences).
func (m *Machine) areStepCycleValid(stepSize int, cycleSize int) error {
	if stepSize <= 0 {
		return &initError{"invalid step size"}
	} else if cycleSize <= 0 {
		return &initError{"invalid cycle size"}
	}

	if ((alphabetSize) % (stepSize * cycleSize)) != 0 {
		return &initError{"cycle size and step size are not compatible, some collisions may occur"}
	}

	return nil
}

// setStep sets rotors' step size.
func (m *Machine) setStep(value int) {
	m.step = value % alphabetSize
}

// Step returns rotors' step size. Step represents the number of positions
// a rotor jumps when taking one step. The default size of a step is 1.
func (m *Machine) Step() int {
	return m.step
}

// setCycle sets size of a rotor's full cycle.
func (m *Machine) setCycle(value int) {
	m.cycle = value
}

// Cycle returns rotors' cycle size. Cycle is the number of steps that
// represent a rotor's full cycle. Cycle is used to indicate when
// rotor should step (move) based on the number of steps taken by preceding
// rotor. The default size of a cycle is 26 (size of the alphabet).
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

// stepRotors turns first rotor one step forward. Other rotors are turned
// accordingy (based on number of previously taken steps).
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
