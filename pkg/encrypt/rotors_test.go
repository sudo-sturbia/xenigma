// Components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"testing"
)

/*
 * Initialization methods
 */

// Test full initialization
func TestInitRotors(t *testing.T) {
	var (
		testMachine   *machine
		testPositions [NUMBER_OF_ROTORS]int
		testStep      int
		testCycleSize int

		err error
	)

	testMachine = new(machine)

	// Correct initialization
	testPositions = [NUMBER_OF_ROTORS]int{2, 3, 4}
	testStep = 1
	testCycleSize = ALPHABET_SIZE

	err = testMachine.initRotors(testPositions, testStep, testCycleSize)
	if err != nil {
		t.Errorf("correct initialization not accepted, error message %s", err.Error())
	}

	if !verifyValues(testMachine, testPositions, testStep, testCycleSize) {
		t.Errorf("values incorrectly initialized")
	}

	// Incorrect initialization 1
	testPositions = [NUMBER_OF_ROTORS]int{-1, 20, 28}

	err = testMachine.initRotors(testPositions, testStep, testCycleSize)
	if err == nil {
		t.Errorf("incorrect positions accepted")
	}

	if !verifyValues(testMachine, [NUMBER_OF_ROTORS]int{0, 0, 0}, testStep, testCycleSize) {
		t.Errorf("values incorrectly initialized")
	}

	// Incorrect initialization 2
	testPositions = [NUMBER_OF_ROTORS]int{2, 3, 4}
	testStep = -223

	err = testMachine.initRotors(testPositions, testStep, testCycleSize)
	if err == nil {
		t.Errorf("incorrect step accepted")
	}

	if !verifyValues(testMachine, testPositions, 1, testCycleSize) {
		t.Errorf("values incorrectly initialized")
	}

	// Incorrect initialization 3
	testStep = 2
	testCycleSize = -4

	err = testMachine.initRotors(testPositions, testStep, testCycleSize)
	if err == nil {
		t.Errorf("incorrect cycle size accepted")
	}

	if !verifyValues(testMachine, testPositions, testStep, ALPHABET_SIZE) {
		t.Errorf("values incorrectly initialized")
	}
}

// Test position setter
func TestSetPosition(t *testing.T) {
	testMachine := new(machine)

	// Test wrong positions
	wrongPositions := [NUMBER_OF_ROTORS]int{3, -22, 6}
	err1 := testMachine.setRotorsPosition(wrongPositions)

	if err1 == nil {
		t.Errorf("incorrect positions accepted")
	}

	// Test correct positions
	correctPositions := [NUMBER_OF_ROTORS]int{3, 22, 6}
	err2 := testMachine.setRotorsPosition(correctPositions)

	if err2 != nil {
		t.Errorf("correct positions not accepted")
	}
}

// Test step setter
func TestSetStep(t *testing.T) {
	var err error

	testMachine := new(machine)

	err = testMachine.setStep(23)
	if err != nil {
		t.Errorf("correct step size \"%d\" not accepted", 23)
	}

	err = testMachine.setStep(0)
	if err == nil {
		t.Errorf("incorrect step size \"%d\" accepted", 0)
	}
}

// Test cycle setter
func TestCycleSetter(t *testing.T) {
	var err error

	testMachine := new(machine)

	err = testMachine.setFullCycle(5)
	if err != nil {
		t.Errorf("correct cycle size \"%d\" not accepted", 5)
	}

	err = testMachine.setFullCycle(0)
	if err == nil {
		t.Errorf("incorrect cycle size \"%d\" accepted", 0)
	}
}

// Check if values in machine were initialized
// correctly based on given input
func verifyValues(testMachine *machine, positions [NUMBER_OF_ROTORS]int, step int, cycleSize int) bool {
	if step != testMachine.step {
		return false
	}

	if cycleSize != testMachine.fullCycle {
		return false
	}

	for i := 0; i < NUMBER_OF_ROTORS; i++ {
		for j := 0; j < ALPHABET_SIZE; j++ {
			if testMachine.rotors[i][j] != (j+positions[i])%ALPHABET_SIZE {
				return false
			}
		}
	}

	return true
}
