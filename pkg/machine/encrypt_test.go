package machine

import (
	"testing"
)

// Test character encryption using config 1.
func TestEncryptCharConfig1(t *testing.T) {
	m, err := read("../../test/data/config-1.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	// r -> u
	encrypted := m.encryptChar('r')
	if encrypted != 'u' {
		t.Errorf("character 'r' encrypted to '%c', expected 'u'", encrypted)
	}

	// , -> ,
	encrypted = m.encryptChar(',')
	if encrypted != ',' {
		t.Errorf("character ',' encrypted to '%c', expected ','", encrypted)
	}

	// ' ' -> ' '
	encrypted = m.encryptChar(' ')
	if encrypted != ' ' {
		t.Errorf("character ' ' encrypted to '%c', expected ' '", encrypted)
	}

	// 1 -> 1
	encrypted = m.encryptChar('1')
	if encrypted != '1' {
		t.Errorf("character '1' encrypted to '%c', expected '1'", encrypted)
	}
}

// Test character encryption using config 2.
func TestEncryptCharConfig2(t *testing.T) {
	m, err := read("../../test/data/config-2.json")
	if err != nil {
		t.Errorf("could not read configurations\n%s", err.Error())
	}

	// s -> d
	encrypted := m.encryptChar('s')
	if encrypted != 'd' {
		t.Errorf("character 's' encrypted to '%c', expected 'd'", encrypted)
	}
}
