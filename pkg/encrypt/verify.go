// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"github.com/sudo-sturbia/enigma/pkg/helper"
)

// Check if all fields of a machine were initialized correctly
func (m *machine) isInitialized() bool {
	// ...

	return false
}

// Return true if pathways are initialized correctly
func (m *machine) isPathwayInit() bool {
	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		if !helper.AreElementsIndices(m.pathConnections[i][:]) {
			return false
		}
	}

	return true
}

// Return true if collector is initialized correctly
func (m *machine) isCollectorInit() bool {
	if !helper.AreElementsIndices(m.collector[:]) || !helper.IsSymmetric(m.collector[:]) {
		return false
	}

	return true
}

// Return true if plugboard connections initialized correctly
func (m *machine) isPlugboardInit() bool {
	if !helper.AreElementsIndices(m.plugboardConnections[:]) || !helper.IsSymmetric(m.plugboardConnections[:]) {
		return false
	}

	return true
}
