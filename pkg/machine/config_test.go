package machine

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

// Test reading of correct configs.
func TestReadCorrectConfig(t *testing.T) {
	_, err1 := read("../../test/data/config-1.json")
	if err1 != nil {
		t.Errorf("error occured while reading correct configuration 1\n%s", err1.Error())
	}

	_, err2 := read("../../test/data/config-2.json")
	if err2 != nil {
		t.Errorf("error occured while reading correct configuration 2\n%s", err2.Error())
	}

	_, err3 := read("../../test/data/config-3.json")
	if err3 != nil {
		t.Errorf("error occured while reading correct configuration 3\n%s", err3.Error())
	}
}

// Test reading of wrong configs.
func TestReadIncorrectConfig(t *testing.T) {
	_, err1 := read("../../test/data/wrong-config-1.json")
	if err1 == nil {
		t.Errorf("error not detected in incorrect configuration 1")
	}

	_, err2 := read("../../test/data/wrong-config-2.json")
	if err2 == nil {
		t.Errorf("error not detected in incorrect configuration 2")
	}

	_, err3 := read("../../test/data/wrong-config-3.json")
	if err3 == nil {
		t.Errorf("error not detected in incorrect configuration 3")
	}

	_, err4 := read("../../test/data/wrong-config-4.json")
	if err4 == nil {
		t.Errorf("error not detected in incorrect configuration 4")
	}

	_, err5 := read("../../test/data/wrong-config-5.json")
	if err5 == nil {
		t.Errorf("error not detected in incorrect configuration 5")
	}
}

// Test loading of a generated machine.
func TestWrite(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	numberOfRotors := rand.Intn(100)
	m := Generate(numberOfRotors)

	os.MkdirAll("../../test/generate", os.ModePerm)
	err := write(m, "../../test/generate/generated-1.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = os.Open("../../test/generate/generated-1.json")
	if err != nil {
		t.Errorf("generated machine was not written correctly")
	}
}

// Test saving and loading of a machine.
func TestReadAndWrite(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	numberOfRotors := rand.Intn(100)
	m := Generate(numberOfRotors)

	os.MkdirAll("../../test/generate", os.ModePerm)
	err := write(m, "../../test/generate/generated-2.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	loaded, err := read("../../test/generate/generated-2.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	// Compare machines
	if !areSimilar(m, loaded) {
		t.Errorf("incorrect saving and reloading of a machine")
	}
}

// Compare two machines.
// Returns true if machines are similar, false otherwise.
func areSimilar(m1 *Machine, m2 *Machine) bool {
	if m1 == nil && m2 == nil {
		return true
	} else if m2 == nil || m1 == nil {
		return false
	}

	if m1.numberOfRotors != m2.numberOfRotors {
		return false
	}

	numberOfRotors := m1.numberOfRotors

	// Pathways
	for i := 0; i < numberOfRotors; i++ {
		for j := 0; j < alphabetSize; j++ {
			if m1.pathConnections[i][j] != m2.pathConnections[i][j] {
				return false
			}
		}
	}

	// Reflector
	for i := 0; i < alphabetSize; i++ {
		if m1.reflector[i] != m2.reflector[i] {
			return false
		}
	}

	// Plugboard
	for i := 0; i < alphabetSize; i++ {
		if m1.plugboardConnections[i] != m2.plugboardConnections[i] {
			return false
		}
	}

	// Rotors
	for i := 0; i < numberOfRotors; i++ {
		if m1.rotors[i][0] != m2.rotors[i][0] {
			return false
		}
	}

	return true
}
