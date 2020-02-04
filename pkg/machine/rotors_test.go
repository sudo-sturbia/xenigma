package machine

import (
	"testing"
)

// Test rotor stepping with different steps and cycle sizes.
func TestStepRotors(t *testing.T) {
	testMachine := new(Machine)

	// Default setting
	testMachine.setNumberOfRotors(3)
	testMachine.initRotors([]int{0, 0, 0}, 1, alphabetSize)
	if takeSteps(testMachine, 1); !isPosCorrect(testMachine, []int{1, 0, 0}) {
		t.Errorf("incorrect positions after 1 step using default configurations")
	}

	if takeSteps(testMachine, alphabetSize); !isPosCorrect(testMachine, []int{1, 1, 0}) {
		t.Errorf("incorrect positions after 27 steps using default configurations")
	}

	// Different step -> 13, different cycle 2
	testMachine.setNumberOfRotors(3)
	testMachine.initRotors([]int{0, 0, 0}, 13, 2)
	if takeSteps(testMachine, 1); !isPosCorrect(testMachine, []int{13, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step of size 13,\n expected %v, got %v",
			[]int{13, 0, 0}, testMachine.RotorPositions(),
		)

	}

	if takeSteps(testMachine, alphabetSize); !isPosCorrect(testMachine, []int{13, 13, 0}) {
		t.Errorf(
			"incorrect positions after 27 steps of size 13, expected %v, got %v",
			[]int{13, 13, 0}, testMachine.RotorPositions(),
		)
	}

	// Different cycle -> 2
	testMachine.setNumberOfRotors(3)
	testMachine.initRotors([]int{0, 0, 0}, 1, 2)
	if takeSteps(testMachine, 1); !isPosCorrect(testMachine, []int{1, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step with cycle size 2, expected %v, got %v",
			[]int{1, 0, 0}, testMachine.RotorPositions(),
		)
	}

	if takeSteps(testMachine, alphabetSize); !isPosCorrect(testMachine, []int{1, 13, 6}) {
		t.Errorf(
			"incorrect positions after 27 steps with cycle size 2, expected %v, got %v",
			[]int{1, 13, 6}, testMachine.RotorPositions(),
		)
	}

	// Different step size -> 2 and cycle size -> 13
	testMachine.setNumberOfRotors(3)
	testMachine.initRotors([]int{2, 4, 0}, 2, 13)

	if takeSteps(testMachine, 20); !isPosCorrect(testMachine, []int{16, 6, 0}) {
		t.Errorf(
			"incorrect positions after 1 step with cycle size 2, expected %v, got %v",
			[]int{16, 6, 0}, testMachine.RotorPositions(),
		)
	}

	// Different number of rotors -> 12
	testMachine.setNumberOfRotors(12)
	testMachine.initRotors([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 26)

	if takeSteps(testMachine, 11881376); !isPosCorrect(testMachine, []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}) {
		t.Errorf(
			"incorrect positions after 11881376 steps for a 12 rotor machine, expected %v, got %v",
			[]int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, testMachine.RotorPositions(),
		)
	}
}

// Test initialization of rotors.
func TestInitRotors(t *testing.T) {
	testMachine := new(Machine)

	// 3 rotors - correct initialization
	testMachine.setNumberOfRotors(3)
	if err := testMachine.initRotors([]int{4, 2, 0}, 2, 13); err != nil {
		t.Errorf("correct init produces err, message %s", err.Error())
	}

	if !arePropertiesCorrect(testMachine, []int{4, 2, 0}, 2, 13) {
		t.Errorf("machine properties are incorrect")
	}

	// 3 rotors - incorrect initialization
	testMachine.setNumberOfRotors(3)
	if err := testMachine.initRotors([]int{3, 1, 0}, 2, 13); err == nil {
		t.Errorf("incorrect init doesn't produce err")
	}

	// 15 rotors - correct initialization
	testMachine.setNumberOfRotors(15)
	if err := testMachine.initRotors([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 26); err != nil {
		t.Errorf("correct init produces err, message %s", err.Error())
	}

	if !arePropertiesCorrect(testMachine, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 26) {
		t.Errorf("machine properties are incorrect")
	}

	// 15 rotors - incorrect initialization
	testMachine.setNumberOfRotors(15)
	if err := testMachine.initRotors([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, 2, 13); err == nil {
		t.Errorf("incorrect init doesn't produce err")
	}

	// 15 rotors - incorrect initialization
	testMachine.setNumberOfRotors(7)
	if err := testMachine.initRotors([]int{2, 4, 2, 4, 2, 4, 2}, 23, 42); err == nil {
		t.Errorf("incorrect init doesn't produce err")
	}
}

// Test position and taken steps setter.
func TestSetPosition(t *testing.T) {
	testMachine := new(Machine)

	// 1st configuration
	testMachine.setNumberOfRotors(3)
	testMachine.setStep(2)
	testMachine.setCycle(13)

	if err := testMachine.setRotorsPosition([]int{2, 4, 4}); err != nil {
		t.Errorf("correct position produces err")
	} else {
		// Verify position
		positions := []int{2, 4, 4}
		for i, pos := range testMachine.RotorPositions() {
			if positions[i] != pos {
				t.Errorf("incorrect position\n expected %v, found %v", positions[i], pos)
			}
		}

		// Verify taken steps
		steps := []int{1, 2, 2}
		for i, step := range testMachine.takenSteps {
			if steps[i] != step {
				t.Errorf("incorrect taken step\n expected %v, found %v", steps[i], step)
			}
		}
	}

	if err := testMachine.setRotorsPosition([]int{1, 2, 0}); err == nil {
		t.Errorf("incorrect position doesn't produce err")
	}

	if err := testMachine.setRotorsPosition([]int{3, -22, 6}); err == nil {
		t.Errorf("incorrect position doesn't produce err")
	}

	// 2nd configuration
	testMachine.setNumberOfRotors(7)
	testMachine.setStep(13)
	testMachine.setCycle(2)

	if err := testMachine.setRotorsPosition([]int{13, 13, 0, 13, 13, 13, 0}); err != nil {
		t.Errorf("correct position produces err")
	} else {
		// Verify position
		positions := []int{13, 13, 0, 13, 13, 13, 0}
		for i, pos := range testMachine.RotorPositions() {
			if positions[i] != pos {
				t.Errorf("incorrect position\n expected %v, found %v", positions[i], pos)
			}
		}

		// Verify taken steps
		steps := []int{1, 1, 0, 1, 1, 1, 0}
		for i, step := range testMachine.takenSteps {
			if steps[i] != step {
				t.Errorf("incorrect taken step\n expected %v, found %v", steps[i], step)
			}
		}
	}

	if err := testMachine.setRotorsPosition([]int{14, 12, 1, 0, 0, 0, 0}); err == nil {
		t.Errorf("incorrect position doesn't produce err")
	}
}

// Test validation of step and cycle sizes.
func TestAreStepCycleValid(t *testing.T) {
	testMachine := new(Machine)

	correct := [][2]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{1, 13},
		{13, 1},
		{1, 26},
		{26, 1},
		{2, 13},
		{13, 2},
	}

	for _, comb := range correct {
		if err := testMachine.areStepCycleValid(comb[0], comb[1]); err != nil {
			t.Errorf("correct step-cycle produces err")
		}
	}

	if err := testMachine.areStepCycleValid(23, 32); err == nil {
		t.Errorf("incorrect step-cycle doesn't produce err")
	}

	if err := testMachine.areStepCycleValid(1, -3); err == nil {
		t.Errorf("incorrect step-cycle doesn't produce err")
	}

	if err := testMachine.areStepCycleValid(0, 1); err == nil {
		t.Errorf("incorrect step-cycle doesn't produce err")
	}
}

// takeSteps steps the rotors given number of steps.
func takeSteps(m *Machine, steps int) {
	for i := 0; i < steps; i++ {
		m.stepRotors()
	}
}

// arePropertiesCorrect checks if values in machine were initialized correctly based on given input.
func arePropertiesCorrect(testMachine *Machine, positions []int, step int, cycleSize int) bool {
	if step != testMachine.step {
		return false
	}

	if cycleSize != testMachine.cycle {
		return false
	}

	return isPosCorrect(testMachine, positions)
}

// isPosCorrect verifies current position of rotors.
func isPosCorrect(testMachine *Machine, positions []int) bool {
	for i := 0; i < testMachine.numberOfRotors; i++ {
		if testMachine.rotors[i] != positions[i]%alphabetSize {
			return false
		}
	}

	return true
}
