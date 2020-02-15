package machine

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Example of usage of the machine package.
func Example() {
	// Generate a random configuration
	rand.Seed(time.Now().UnixNano())

	numberOfRotors := rand.Intn(100)
	m := Generate(numberOfRotors)

	// Encrypt a message
	message := "Hello, world!"
	encrypted, _ := m.Encrypt(message)

	fmt.Printf("message: %s, encryption: %s\n", message, encrypted)

	// Write configurations to a JSON file
	err := m.Write("generate/generated.json")
	if err != nil {
		panic("couldn't save configurations")
	}
}

// Test encryption of strings.
func TestEncrypt(t *testing.T) {
	m1, err := Read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	encrypted1, _ := m1.Encrypt("Hello, world!")
	if encrypted1 != "sispr, areko!" {
		t.Errorf("incorrect encryption of \"Hello, world!\",\nexpected \"sispr, areko!\", got \"%s\"", encrypted1)
	}

	m2, err := Read("../../test/data/config-2.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	encrypted2, _ := m2.Encrypt("Hello, again!")
	if encrypted2 != "lcsml, fccmb!" {
		t.Errorf("incorrect encryption of \"Hello, again!\",\nexpected \"lcsml, fccmb!\", got \"%s\"", encrypted2)
	}
}

// Benchmark encryption using a 1000-rotor machine.
func BenchmarkEncrypt(b *testing.B) {
	m := Generate(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Encrypt("Hello, world!\nThis is a benchmark.")
	}
}

// Benchmark encryption of README.md using a 1000-rotor machine.
func BenchmarkEncryptREADME(b *testing.B) {
	m := Generate(1000)
	message := helper.ReadStringFromFile("../../README.md")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Encrypt(message)
	}
}

// Test encryption and decryption of messages.
func TestEncryptDecrypt(t *testing.T) {
	encryptor, err := Read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	decryptor, err := Read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	message := "Hello, world!"
	encrypted, _ := encryptor.Encrypt(message)
	decrypted, _ := decryptor.Encrypt(encrypted)

	if decrypted != strings.ToLower(message) {
		t.Errorf("incorrect decryption of %s,\nexpected %s, got %s", encrypted, message, decrypted)
	}

	message = "This is an example of encryption using an enigma machine.\n" +
		"Encrypted messages can also be decrypted using the same machine."
	encrypted, _ = encryptor.Encrypt(message)
	decrypted, _ = decryptor.Encrypt(encrypted)

	if decrypted != strings.ToLower(message) {
		t.Errorf("incorrect decryption of %s,\nexpected %s, got %s", encrypted, message, decrypted)
	}
}

// Test reading, writing, and encryption.
func TestReadWriteEncrypt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	numberOfRotors := rand.Intn(100)
	m := Generate(numberOfRotors)

	err := m.Write("../../test/generate/generated-3.json")
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
	if encrypted1 != 'n' {
		t.Errorf("character 'r' encrypted to '%c', expected 'n'", encrypted1)
	}

	// Using configuration 2
	m2, err := Read("../../test/data/config-2.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	// s -> d
	encrypted2 := m2.encryptChar('s')
	if encrypted2 != 'r' {
		t.Errorf("character 's' encrypted to '%c', expected 'r'", encrypted2)
	}

}

// Test encryption of a list of non-alphabetical characters.
// Characters are not meant to change when encrypted.
func TestEncryptCharNonAlpha(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	numberOfRotors := rand.Intn(100)
	m := Generate(numberOfRotors)

	nonAlpha := []byte{',', ' ', '1', '\n', '[', '\t'}
	for _, char := range nonAlpha {
		encrypted := m.encryptChar(char)
		if encrypted != char {
			t.Errorf("character '%c' encrypted to '%c', expected '%c'", char, encrypted, char)
		}
	}
}
