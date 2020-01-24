package machine

import (
	"testing"
)

// Test encryption of individual alphabetical characters.
func TestEncryptCharAlpha(t *testing.T) {
	// Using configuration 1
	m1, err := read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	// r -> u
	encrypted1 := m1.encryptChar('r')
	if encrypted1 != 'u' {
		t.Errorf("character 'r' encrypted to '%c', expected 'u'", encrypted1)
	}

	// Using configuration 2
	m2, err := read("../../test/data/config-2.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	// s -> d
	encrypted2 := m2.encryptChar('s')
	if encrypted2 != 'd' {
		t.Errorf("character 's' encrypted to '%c', expected 'd'", encrypted2)
	}

}

// Test encryption of a list of non-alphabetical characters.
// Characters are not meant to change when encrypted.
func TestEncryptCharNonAlpha(t *testing.T) {
	m := Generate()

	nonAlpha := []byte{',', ' ', '1', '\n', '[', '\t'}
	for _, char := range nonAlpha {
		encrypted := m.encryptChar(char)
		if encrypted != char {
			t.Errorf("character '%c' encrypted to '%c', expected '%c'", char, encrypted, char)
		}
	}
}
