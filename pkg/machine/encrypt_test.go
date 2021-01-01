package machine

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// Example of usage of the machine package.
func Example() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random machine.
	m := Generate(rand.Intn(100) + 3)

	// Encrypt a message.
	message := "Hello, world!"
	encrypted, _ := m.Encrypt(message)

	fmt.Printf("message: %s, encryption: %s\n", message, encrypted)

	// Write the machine to a JSON file.
	err := Write(m, "generate/machine.json")
	if err != nil {
		panic("failed to write machine.json")
	}
}

// Test encryption of strings.
func TestEncrypt(t *testing.T) {
	for i, test := range []struct {
		path    string
		message string
		want    string
	}{
		{
			path:    "../../test/data/config-1.json",
			message: "Hello, world!",
			want:    "sispr, areko!",
		},
		{
			path:    "../../test/data/config-2.json",
			message: "Hello, again!",
			want:    "lcsml, fccmb!",
		},
	} {
		m, err := Read(test.path)
		if err != nil {
			t.Errorf("test %d: failed to read machine: %w", i, err)
		}

		got, err := m.Encrypt(test.message)
		if err != nil {
			t.Errorf("test %d: failed to encrypt: %w", i, err)
		}

		if got != test.want {
			t.Errorf("test %d: incorrect encryption, want: %s, got: %s", i, test.want, got)
		}
	}
}

// BenchmarkEncrypt benchmarks encryption of a small message using a
// 1000-rotor machine.
func BenchmarkEncrypt(b *testing.B) {
	m := Generate(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Encrypt("Hello, world!\nThis is a benchmark.")
	}
}

// BenchmarkEncryptREADME benchmarks encryption of README.md using a
// 1000-rotor machine.
func BenchmarkEncryptREADME(b *testing.B) {
	m := Generate(1000)
	contents, err := ioutil.ReadFile("../../README.md")
	if err != nil {
		b.Fatalf("failed to read contents of readme: %s", err.Error())
	}
	readme := string(contents)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Encrypt(readme)
	}
}

// TestEncryptDecrypt compares a message with its decryption.
func TestEncryptDecrypt(t *testing.T) {
	path := "../../test/data/config-1.json"

	encryptor, err := Read(path)
	if err != nil {
		t.Fatalf("failed to read machine: %s", err.Error())
	}
	decryptor, err := Read(path)
	if err != nil {
		t.Fatalf("failed to read machine: %s", err.Error())
	}

	for _, message := range []string{
		"Hello, world!",
		"This is an encryption example using a xenigma machine.\n" +
			"Encrypted messages can also be decrypted using the same machine.",
	} {
		encrypted, err := encryptor.Encrypt(message)
		if err != nil {
			t.Errorf("failed to encrypt message: %w", err)
		}

		decrypted, err := decryptor.Encrypt(encrypted)
		if err != nil {
			t.Errorf("failed to encrypt message: %w", err)
		}

		if decrypted != strings.ToLower(message) {
			t.Errorf("failed to decrypt: want %s, got %s", message, decrypted)
		}
	}
}

// TestReadWriteEncrypt generates a machine, writes it to a file, rereads
// it, and compares the encryption of the original and read machines.
func TestReadWriteEncrypt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	m := Generate(rand.Intn(100) + 3)
	err := Write(m, "../../test/generate/generated.json")
	if err != nil {
		t.Errorf("failed to write machine: %w", err)
	}

	r, err := Read("../../test/generate/generated.json")
	if err != nil {
		t.Errorf("failed to read machine: %w", err)
	}

	for _, message := range []string{
		"Hello, World!",
		"Another message,",
		"yet another message.",
		"One final message that's a bit longer, but not too long.",
	} {
		original, err := m.Encrypt(message)
		if err != nil {
			t.Errorf("failed to encrypt: %w", err)
		}
		read, err := r.Encrypt(message)
		if err != nil {
			t.Errorf("failed to encrypt: %w", err)
		}

		if original != read {
			t.Errorf("different encryption for: %s, original: %s, read: %s", message, original, read)
		}
	}
}

func TestEncryptCharNonAlpha(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	m := Generate(rand.Intn(100) + 3)
	reversed := reverseConnections(m)
	for _, c := range []byte{
		',',
		' ',
		'1',
		'\n',
		'[',
		'\t',
	} {
		enc := m.encryptChar(c, reversed)
		if enc != c {
			t.Errorf("failed to encrypt '%c', want '%c', got '%c'", c, c, enc)
		}
	}
}
