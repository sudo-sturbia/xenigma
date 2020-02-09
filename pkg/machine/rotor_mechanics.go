package machine

import (
	"fmt"
)

// stepRotors turns fast rotor one step forward. Other rotors are turned
// accordingy (based on number of previously taken steps).
func (m *Machine) stepRotors() {
	for i, rotor := range m.rotors {
		if i == 0 || (m.rotors[i-1].takenSteps == 0) { // Previous rotor completed a full cycle
			rotor.step(m.step, m.cycle)
		} else { // Rotor didn't shift, so next rotors won't shift
			break
		}
	}
}

// setRotorsFields verifies and sets machine's rotors, step, and cycle.
// Returns an init error if rotors are incorrectly initialized.
func (m *Machine) setRotorsFields(rotors []Rotor, step, cycle int) error {
	if rotors == nil || len(rotors) == 0 {
		return &initError{"no rotors given"}
	}

	m.setNumberOfRotors(len(rotors))

	if err := m.setStepAndCycle(step, cycle); err != nil {
		return err
	}

	if err := m.setRotors(rotors, step, cycle); err != nil {
		return err
	}

	return nil
}

// DefaultRotorProperties sets rotor's step, cycle, and setting to their
// default values. Defaults are "a"'s for rotors' setting, 1 for step size,
// and 26 (size of the alphabet) for cycle size.
func (m *Machine) DefaultRotorProperties() {
	m.resetRotorsPosition()

	m.setStep(DefaultStep)
	m.setCycle(DefaultCycle)
}

// setRotors verifies and sets machine's rotors' slice.
// Returns an error if given rotors are incorrect.
func (m *Machine) setRotors(rotors []Rotor, step, cycle int) error {
	m.rotors = make([]Rotor, m.numberOfRotors)
	for i, rotor := range rotors {
		err := m.rotors[i].InitRotor(rotor.pathways, rotor.position, step, cycle)
		if err != nil {
			return fmt.Errorf("rotor %d: %s", i, err.Error())
		}
	}

	return nil
}

// resetRotorsPosition sets position and number of taken steps of
// machine's rotors to 0.
func (m *Machine) resetRotorsPosition() {
	for _, rotor := range m.rotors {
		rotor.resetPosition()
	}
}

// Setting returns current rotor setting (position of machine's rotors).
func (m *Machine) Setting() {
	setting := make([]int, m.numberOfRotors)
	for i, rotor := range m.rotors {
		setting[i] = rotor.Position()
	}
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
