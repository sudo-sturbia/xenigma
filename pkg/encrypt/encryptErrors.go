// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

// Connection error
type connectionErr struct {
	message string
}

func (err *connectionErr) Error() string {
	return err.message
}

// Rotor configuration error
type rotorConfigErr struct {
	message string
}

func (err *rotorConfigErr) Error() string {
	return err.message
}

// Initializion error
type initError struct {
	message string
}

func (err *initError) Error() string {
	return err.message
}
