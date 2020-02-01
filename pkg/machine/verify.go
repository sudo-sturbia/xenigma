package machine

import (
	"github.com/sudo-sturbia/enigma/pkg/helper"
)

// Initializion error
type initError struct {
	message string
}

func (err *initError) Error() string {
	return "incorrect init, " + err.message
}

// isInit verifies that all fields of a machine were initialized
// correctly. If not an error is returned.
func (m *Machine) isInit() error {
	switch {
	case len(m.pathConnections) != m.numberOfRotors:
		return &initError{"invalid number of pathway layers"}
	case len(m.rotors) != m.numberOfRotors:
		return &initError{"invalid number of rotors"}
	case !m.areRotorsInit():
		return &initError{"invalid rotor configs"}
	case !m.arePathwaysInit():
		return &initError{"invalid pathway connections"}
	case !m.isReflectorInit():
		return &initError{"invalid reflector connections"}
	case !m.isPlugboardInit():
		return &initError{"invalid plugboard connections"}
	case m.step <= 0:
		return &initError{"invalid step size"}
	case m.cycle <= 0:
		return &initError{"invalid cycle size"}
	default:
		return nil
	}
}

// areRotorsInit returns true if rotors are initialized correctly and
// verifies all rotor related values.
func (m *Machine) areRotorsInit() bool {
	if m.rotors == nil || m.takenSteps == nil {
		return false
	}

	for _, rotorHead := range m.rotors {
		if rotorHead < 0 || rotorHead >= alphabetSize {
			return false
		}
	}

	if err := m.areStepCycleValid(m.step, m.cycle); err != nil {
		return false
	}

	if err := m.arePositionsValid(m.RotorPositions()); err != nil {
		return false
	}

	return true
}

// arePathwaysInit returns true if pathways are initialized correctly.
func (m *Machine) arePathwaysInit() bool {
	if m.pathConnections == nil {
		return false
	}

	for _, pathwaysArr := range m.pathConnections {
		if !helper.AreElementsIndices(pathwaysArr[:]) {
			return false
		}
	}

	return true
}

// isReflectorInit returns true if reflector is initialized correctly.
func (m *Machine) isReflectorInit() bool {
	return helper.AreElementsIndices(m.reflector[:]) &&
		helper.IsSymmetric(m.reflector[:])
}

// isPlugboardInit returns true if plugboard connections initialized correctly.
func (m *Machine) isPlugboardInit() bool {
	return helper.AreElementsIndices(m.plugboardConnections[:]) &&
		helper.IsSymmetric(m.plugboardConnections[:])
}
