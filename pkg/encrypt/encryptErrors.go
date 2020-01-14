// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

// Connection error
type connectionErr struct {
	message string
}

func (err *connectionErr) Error() string {
	return "connections error: " + err.message
}

// Rotor configuration error
type rotorConfigErr struct {
	message string
}

func (err *rotorConfigErr) Error() string {
	return "rotor configuration error: " + err.message
}

// Initializion error
type initError struct {
	message string
}

func (err *initError) Error() string {
	return "initialization error: " + err.message
}
