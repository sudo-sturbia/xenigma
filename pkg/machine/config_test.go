package machine

import (
	"testing"
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
}
