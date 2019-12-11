// Encrypt messages using engima code
package encrypt

import ()

// Validate and set plugboard connections
func (m *machine) createPlugboardConnections(connectionsArr [2][]byte) error {
	m.plugboardConnections = make(map[byte]byte)

	// Validate length
	if len(connectionsArr[0]) != len(connectionsArr[1]) || len(connectionsArr[0]) > 13 || len(connectionsArr[1]) > 13 {
		return &connectionErr{"Incorrect number of connections"}
	}

	// Validate character connections
	isConnected := make(map[byte]bool)

	for i := 0; i < len(connectionsArr[0]); i++ {
		if (connectionsArr[0][i] != connectionsArr[1][i]) && !isConnected[connectionsArr[0][i]] && !isConnected[connectionsArr[1][i]] {
			isConnected[connectionsArr[0][i]] = true
			isConnected[connectionsArr[1][i]] = true

			m.plugboardConnections[connectionsArr[0][i]] = connectionsArr[1][i]
			m.plugboardConnections[connectionsArr[1][i]] = connectionsArr[0][i]
		} else {
			return &connectionErr{"Incorrect number of connections for a character"}
		}
	}

	return nil
}

// Change character based on plugboard connections
func (m *machine) changeChar(char byte) byte {
	if m.plugboardConnections[char] != 0 {
		return m.plugboardConnections[char]
	} else {
		return char
	}
}
