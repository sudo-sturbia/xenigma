// Encrypt messages using engima code
package encrypt

import ()

var plugboardConnections map[byte]byte // Map with max size of 26

// Validate and set plugboard connections
func createPlugboardConnections(connectionsArr [2][]byte) error {
	plugboardConnections = make(map[byte]byte)

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

			plugboardConnections[connectionsArr[0][i]] = connectionsArr[1][i]
			plugboardConnections[connectionsArr[1][i]] = connectionsArr[0][i]
		} else {
			return &connectionErr{"Incorrect number of connections for a character"}
		}
	}

	return nil
}

// Change character based on plugboard connections
func changeChar(char byte) byte {
	if plugboardConnections[char] != 0 {
		return plugboardConnections[char]
	} else {
		return char
	}
}
