// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"testing"
)

// Initialization methods

// Test full initialization
func TestInitRotors(t *testing.T) {
	var (
		testMachine   *Machine
		testPositions [NumberOfRotors]int
		testStep      int
		testCycleSize int

		err error
	)

	testMachine = new(Machine)

	// Correct initialization
	testPositions = [NumberOfRotors]int{2, 3, 4}
	testStep = 1
	testCycleSize = alphabetSize

	err = testMachine.initRotors(testPositions, testStep, testCycleSize)
	if err != nil {
		t.Errorf("correct initialization not accepted, error message %s", err.Error())
	}

	if !verifyValues(testMachine, testPositions, testStep, testCycleSize) {
		t.Errorf("values incorrectly initialized")
	}

	// Incorrect initialization 1
	testPositions = [NumberOfRotors]int{-1, 20, 28}

	err = testMachine.initRotors(testPositions, testStep, testCycleSize)
	if err == nil {
		t.Errorf("incorrect positions accepted")
	}

	if !verifyValues(testMachine, [NumberOfRotors]int{0, 0, 0}, testStep, testCycleSize) {
		t.Errorf("values incorrectly initialized")
	}

	// Incorrect initialization 2
	testPositions = [NumberOfRotors]int{2, 3, 4}
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

	if !verifyValues(testMachine, testPositions, testStep, alphabetSize) {
		t.Errorf("values incorrectly initialized")
	}
}

// Test position setter
func TestSetPosition(t *testing.T) {
	testMachine := new(Machine)

	// Test wrong positions
	wrongPositions := [NumberOfRotors]int{3, -22, 6}
	err1 := testMachine.setRotorsPosition(wrongPositions)

	if err1 == nil {
		t.Errorf("incorrect positions accepted")
	}

	// Test correct positions
	correctPositions := [NumberOfRotors]int{3, 22, 6}
	err2 := testMachine.setRotorsPosition(correctPositions)

	if err2 != nil {
		t.Errorf("correct positions not accepted")
	}
}

// Test step setter
func TestSetStep(t *testing.T) {
	var err error

	testMachine := new(Machine)

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

	testMachine := new(Machine)

	err = testMachine.setCycle(5)
	if err != nil {
		t.Errorf("correct cycle size \"%d\" not accepted", 5)
	}

	err = testMachine.setCycle(0)
	if err == nil {
		t.Errorf("incorrect cycle size \"%d\" accepted", 0)
	}
}

// Stepping

// Test rotor stepping
// with different steps and cycle sizes.
func TestStepRotors(t *testing.T) {
	testMachine := new(Machine)

	// Default setting
	testMachine.initRotors([NumberOfRotors]int{0, 0, 0}, 1, alphabetSize)

	testMachine.stepRotors()
	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{1, 0, 0}) {
		t.Errorf("incorrect positions after 1 step")
	}

	for i := 0; i < alphabetSize; i++ {
		testMachine.stepRotors()
	}

	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{1, 1, 0}) {
		t.Errorf("incorrect positions after full cycle")
	}

	// Different step -> 5
	testMachine.initRotors([NumberOfRotors]int{0, 0, 0}, 5, alphabetSize)

	testMachine.stepRotors()
	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{5, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step of size 5, expected %v, got %v",
			[NumberOfRotors]int{5, 0, 0},
			[NumberOfRotors]int{testMachine.rotors[0][0], testMachine.rotors[1][0], testMachine.rotors[2][0]},
		)
	}

	for i := 0; i < alphabetSize; i++ {
		testMachine.stepRotors()
	}

	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{5, 5, 0}) {
		t.Errorf(
			"incorrect positions after 26 steps of with step size 5, expected %v, got %v",
			[NumberOfRotors]int{5, 5, 0},
			[NumberOfRotors]int{testMachine.rotors[0][0], testMachine.rotors[1][0], testMachine.rotors[2][0]},
		)
	}

	// Different cycle -> 3
	testMachine.initRotors([NumberOfRotors]int{0, 0, 0}, 1, 3)

	testMachine.stepRotors()
	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{1, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step with cycle size 3, expected %v, got %v",
			[NumberOfRotors]int{1, 0, 0},
			[NumberOfRotors]int{testMachine.rotors[0][0], testMachine.rotors[1][0], testMachine.rotors[2][0]},
		)
	}

	for i := 0; i < alphabetSize; i++ {
		testMachine.stepRotors()
	}

	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{1, 9, 3}) {
		t.Errorf(
			"incorrect positions after 26 steps with cycle size 3, expected %v, got %v",
			[NumberOfRotors]int{1, 9, 3},
			[NumberOfRotors]int{testMachine.rotors[0][0], testMachine.rotors[1][0], testMachine.rotors[2][0]},
		)
	}

	// Different cycle size -> 4, step size -> 7
	testMachine.initRotors([NumberOfRotors]int{0, 0, 0}, 7, 4)

	testMachine.stepRotors()
	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{7, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step with step 7, and cycle 4, expected %v, got %v",
			[NumberOfRotors]int{7, 0, 0},
			[NumberOfRotors]int{testMachine.rotors[0][0], testMachine.rotors[1][0], testMachine.rotors[2][0]},
		)
	}

	for i := 0; i < alphabetSize; i++ {
		testMachine.stepRotors()
	}

	if !verifyRotorsPos(testMachine, [NumberOfRotors]int{7, 16, 7}) {
		t.Errorf(
			"incorrect positions after 26 steps with step 7, and cycle 4, expected %v, got %v",
			[NumberOfRotors]int{7, 16, 7},
			[NumberOfRotors]int{testMachine.rotors[0][0], testMachine.rotors[1][0], testMachine.rotors[2][0]},
		)
	}
}

// Verification methods

// Check if values in machine were initialized
// correctly based on given input
func verifyValues(testMachine *Machine, positions [NumberOfRotors]int, step int, cycleSize int) bool {
	if step != testMachine.step {
		return false
	}

	if cycleSize != testMachine.cycle {
		return false
	}

	return verifyRotorsPos(testMachine, positions)
}

// Verify current position of rotors
func verifyRotorsPos(testMachine *Machine, positions [NumberOfRotors]int) bool {
	for i := 0; i < NumberOfRotors; i++ {
		for j := 0; j < alphabetSize; j++ {
			if testMachine.rotors[i][j] != (j+positions[i])%alphabetSize {
				return false
			}
		}
	}

	return true
}
