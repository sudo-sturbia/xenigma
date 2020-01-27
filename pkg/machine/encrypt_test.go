package machine

import (
	"fmt"
	"testing"
)

// Example of usage of the enigma package.
func Example() {
	// Generate a random configuration
	m := Generate()

	// Encrypt a message
	message := "Hello, world!"
	encrypted, _ := m.Encrypt(message)

	fmt.Printf("message: %s, encryption: %s\n", message, encrypted)

	// Write configurations to a JSON file
	err := Write(m, "generate/generated.json")
	if err != nil {
		panic("couldn't save configuratoins")
	}
}

// Test encryption of strings.
func TestEncrypt(t *testing.T) {
	m1, err := Read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	encrypted1, _ := m1.Encrypt("Hello, world!")
	if encrypted1 != "suelb, dpkqr!" {
		t.Errorf("incorrect encryption of \"Hello, world!\",\nexpected \"suelb, dpkqr!\", got \"%s\"", encrypted1)
	}

	m2, err := Read("../../test/data/config-2.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	encrypted2, _ := m2.Encrypt("Hello, again!")
	if encrypted2 != "rbhxx, zihgu!" {
		t.Errorf("incorrect encryption of \"Hello, again!\",\nexpected \"rbhxx, zihgu!\", got \"%s\"", encrypted2)
	}
}

func TestReadWriteEncrypt(t *testing.T) {
	m := Generate()

	err := Write(m, "../../test/generate/generated-3.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	loaded, err := Read("../../test/generate/generated-3.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	message := "Hello, world!"
	originalEnc, err := m.Encrypt(message)
	if err != nil {
		t.Errorf(err.Error())
	}

	loadedEnc, err := loaded.Encrypt(message)
	if err != nil {
		t.Errorf(err.Error())
	}

	if originalEnc != loadedEnc {
		t.Errorf("same message encrypted differently,\noriginal machine: \"%s\",\nloaded machine: \"%s\"", originalEnc, loadedEnc)
	}
}

// Test encryption of individual alphabetical characters.
func TestEncryptCharAlpha(t *testing.T) {
	// Using configuration 1
	m1, err := Read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	// r -> u
	encrypted1 := m1.encryptChar('r')
	if encrypted1 != 'u' {
		t.Errorf("character 'r' encrypted to '%c', expected 'u'", encrypted1)
	}

	// Using configuration 2
	m2, err := Read("../../test/data/config-2.json")
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
