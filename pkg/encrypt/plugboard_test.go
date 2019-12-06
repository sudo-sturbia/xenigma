// Encrypt messages using engima code
package encrypt

import (
	"testing"
)

func TestSetConnections(t *testing.T) {
	correct := [2][]byte{
		{'a', 'b', 'c', 'd', 'e', 'f'},
		{'j', 'k', 'l', 'y', 'u', 'z'},
	}

	wrong1 := [2][]byte{
		{'a', 'b', 'k', 'h'},
		{'h', 'g', 'l', 'd'},
	}

	wrong2 := [2][]byte{
		{'a'},
		{'a'},
	}

	wrong3 := [2][]byte{
		{'h', 'l', 't', 'y', 'j', 'd', 'v', 'i'},
		{'g', 'k', 'f', 'g'},
	}

	wrong4 := [2][]byte{
		{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'o'},
		{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'o'},
	}

	err1 := setConnections(correct)
	err2 := setConnections(wrong1)
	err3 := setConnections(wrong2)
	err4 := setConnections(wrong3)
	err5 := setConnections(wrong4)

	if err1 != nil {
		t.Errorf("Correct connection fails")
	}

	if err2 == nil {
		t.Errorf("Wrong connection passed, same character mapped twice")
	}

	if err3 == nil {
		t.Errorf("Wrong connection passed, character mapped to itself")
	}

	if err4 == nil {
		t.Errorf("Wrong connection passed, invalid number of connections")
	}

	if err5 == nil {
		t.Errorf("Wrong connection passed, more than 13 characters in a row")
	}
}

func TestChangeChar(t *testing.T) {
	connection := [2][]byte{
		{'a', 'b', 'c', 'd', 'e', 'f'},
		{'j', 'k', 'l', 'y', 'u', 'z'},
	}

	setConnections(connection)

	for i := 0; i < len(connection[0]); i++ {
		if changeChar(connection[0][i]) != connection[1][i] || changeChar(connection[1][i]) != connection[0][i] {
			t.Errorf("Mapping incorrect, expected %v -> %v, got %v -> %v", connection[0][i], connection[1][i], connection[0][i], changeChar(connection[0][i]))
		}
	}
}
