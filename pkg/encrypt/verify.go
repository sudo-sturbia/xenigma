// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"github.com/sudo-sturbia/enigma/pkg/helper"
)

// Check if all fields of a machine were initialized correctly
func (m *Machine) isInit() bool {
	isInit := m.areRotorsInit() && m.arePathwaysInit() && m.isCollectorInit() &&
		m.isPlugboardInit() && (m.step > 0) && (m.cycle > 0)

	return isInit
}

// Return true if rotors are initialized correctly
func (m *Machine) areRotorsInit() bool {
	for _, rotor := range m.rotors {
		if !helper.AreElementsOrderedIndices(rotor[:]) {
			return false
		}
	}

	return true
}

// Return true if pathways are initialized correctly
func (m *Machine) arePathwaysInit() bool {
	for _, pathwaysArr := range m.pathConnections {
		if !helper.AreElementsIndices(pathwaysArr[:]) {
			return false
		}
	}

	return true
}

// Return true if collector is initialized correctly
func (m *Machine) isCollectorInit() bool {
	return helper.AreElementsIndices(m.collector[:]) && helper.IsSymmetric(m.collector[:])
}

// Return true if plugboard connections initialized correctly
func (m *Machine) isPlugboardInit() bool {
	return helper.AreElementsIndices(m.plugboardConnections[:]) && helper.IsSymmetric(m.plugboardConnections[:])
}
