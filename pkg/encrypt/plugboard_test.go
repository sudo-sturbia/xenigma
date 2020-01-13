// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"testing"
)

func TestCreatePlugboardConnections(t *testing.T) {
	correct := map[byte]byte{'a': 'j', 'b': 'k', 'c': 'l', 'd': 'y', 'e': 'u', 'f': 'z'}

	wrong1 := map[byte]byte{'a': 'h', 'b': 'g', 'k': 'l', 'h': 'd'}
	wrong2 := map[byte]byte{'a': 'a', 'b': 'a', 'c': 'a', 'd': 'a', 'e': 'a', 'f': 'a', 'g': 'a', 'h': 'a', 'j': 'a', 'k': 'a', 'l': 'a', 'm': 'a', 'n': 'a', 'o': 'a'}

	testMachine := new(Machine)
	err1 := testMachine.createPlugboardConnections(correct)
	err2 := testMachine.createPlugboardConnections(wrong1)
	err3 := testMachine.createPlugboardConnections(wrong2)

	if err1 != nil {
		t.Errorf("Correct connection fails, %s", err1.Error())
	}

	if err2 == nil {
		t.Errorf("Wrong connection passed, same character mapped twice")
	}

	if err3 == nil {
		t.Errorf("Wrong connection passed, invalid number of connections")
	}
}

func TestPlug(t *testing.T) {
	connection := map[byte]byte{'a': 'j', 'b': 'k', 'c': 'l', 'd': 'y', 'e': 'u', 'f': 'z'}

	testMachine := new(Machine)
	testMachine.createPlugboardConnections(connection)

	for key, value := range connection {
		if testMachine.plugIn(key) != int(value-'a') {
			t.Errorf("Incorrect mapping plugIn 1, expected %c -> %c, got %c -> %c", key, value, key, byte(testMachine.plugIn(key))+'a')
		}

		if testMachine.plugIn(value) != int(key-'a') {
			t.Errorf("Incorrect mapping plugIn 2, expected %c -> %c, got %c -> %c", value, key, value, byte(testMachine.plugIn(value))+'a')
		}

		if testMachine.plugOut(int(key-'a')) != value {
			t.Errorf("Incorrect mapping plugOut 1, expected %c -> %c, got %c -> %c", key, value, key, testMachine.plugOut(int(key-'a')))
		}

		if testMachine.plugOut(int(value-'a')) != key {
			t.Errorf("Incorrect mapping plugOut 2, expected %c -> %c, got %c -> %c", value, key, value, testMachine.plugOut(int(value-'a')))
		}
	}
}
