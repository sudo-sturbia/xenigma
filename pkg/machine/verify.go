package machine

import (
	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Initializion error
type initError struct {
	message string
}

func (err *initError) Error() string {
	return "incorrect init, " + err.message
}

// IsConfigCorrect verifies that all fields of the machine are
// initialized correctly, returns an error if not.
func (m *Machine) IsConfigCorrect() error {
	switch {
	case !m.areRotorsCorrect():
		return &initError{"rotors configurations are invalid"}
	case !m.isReflectorCorrect():
		return &initError{"reflector connections are incorrect"}
	case !m.isPlugboardCorrect():
		return &initError{"plugboard connections are incorrect"}
	default:
		return nil
	}
}

// areRotorsCorrect returns true if rotors are initialized correctly and
// verifies all rotor related values.
func (m *Machine) areRotorsCorrect() bool {
	if m.rotors == nil || len(m.rotors) == 0 || m.numberOfRotors != len(m.rotors) {
		return false
	}

	for _, rotor := range m.rotors {
		if err := rotor.IsConfigCorrect(); err != nil {
			return false
		}
	}

	return true
}

// isReflectorCorrect returns true if reflector is initialized correctly.
func (m *Machine) isReflectorCorrect() bool {
	return helper.AreElementsIndices(m.reflector[:]) &&
		helper.IsSymmetric(m.reflector[:])
}

// isPlugboardCorrect returns true if plugboard connections initialized correctly.
func (m *Machine) isPlugboardCorrect() bool {
	return helper.AreElementsIndices(m.plugboard[:]) &&
		helper.IsSymmetric(m.plugboard[:])
}
