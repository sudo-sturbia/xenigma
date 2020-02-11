package machine

import (
	"fmt"
)

// stepRotors turns fast rotor one step forward. Other rotors are turned
// accordingy (based on number of previously taken steps).
func (m *Machine) stepRotors() {
	for i, rotor := range m.rotors {
		if i == 0 || (m.rotors[i-1].takenSteps == 0) { // Previous rotor completed a full cycle
			rotor.takeStep()
		} else { // Rotor didn't shift, so next rotors won't shift
			break
		}
	}
}

// SetRotors verifies and sets machine's rotors. Returns an error
// if configurations of given rotors are incorrect, nil otherwise.
func (m *Machine) SetRotors(rotors []Rotor) error {
	if rotors == nil || len(rotors) <= 0 {
		return &initError{"no rotors given"}
	}

	m.setNumberOfRotors(len(rotors))

	m.rotors = make([]Rotor, m.numberOfRotors)
	for i, rotor := range rotors {
		if err := rotor.IsConfigCorrect(); err != nil {
			return fmt.Errorf("rotor %d: %s", i, err.Error())
		}

		m.rotors[i] = rotor
	}

	return nil
}

// Setting returns current rotor setting (position of machine's rotors).
func (m *Machine) Setting() {
	setting := make([]int, m.numberOfRotors)
	for i, rotor := range m.rotors {
		setting[i] = rotor.Position()
	}
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
