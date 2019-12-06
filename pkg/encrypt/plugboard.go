// Encrypt messages using engima code
package encrypt

import ()

var connections map[byte]byte // Map with max size of 26

// Validate and set plugboard connections
func setConnections(connectionsArr [2][]byte) error {
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

			connections[connectionsArr[0][i]] = connectionsArr[1][i]
			connections[connectionsArr[1][i]] = connectionsArr[0][i]
		} else {
			return &connectionErr{"Incorrect number of connections for a character"}
		}
	}

	return nil
}

// Change character based on plugboard connections
func changeChar(char byte) byte {
	if connections[char] != nil {
		return connections[char]
	} else {
		return char
	}
}
