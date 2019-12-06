// Encrypt messages using engima code
package encrypt

// Connection Error
type connectionErr struct {
	message string
}

func (err *connectionErr) Error() string {
	return err.message
}
