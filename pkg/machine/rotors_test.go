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
		t.Errorf("incorrect positions after 1 step")
	}

	if takeSteps(testMachine, alphabetSize); !isPosCorrect(testMachine, []int{1, 1, 0}) {
		t.Errorf("incorrect positions after 27 steps")
	}

	// Different step -> 5
	testMachine.setNumberOfRotors(3)
	testMachine.initRotors([]int{0, 0, 0}, 5, alphabetSize)
	if takeSteps(testMachine, 1); !isPosCorrect(testMachine, []int{5, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step of size 5, expected %v, got %v",
			[]int{5, 0, 0}, testMachine.CurrentRotors(),
		)

	}

	if takeSteps(testMachine, alphabetSize); !isPosCorrect(testMachine, []int{5, 5, 0}) {
		t.Errorf(
			"incorrect positions after 27 steps of with step size 5, expected %v, got %v",
			[]int{5, 5, 0}, testMachine.CurrentRotors(),
		)
	}

	// Different cycle -> 3
	testMachine.setNumberOfRotors(3)
	testMachine.initRotors([]int{0, 0, 0}, 1, 3)
	if takeSteps(testMachine, 1); !isPosCorrect(testMachine, []int{1, 0, 0}) {
		t.Errorf(
			"incorrect positions after 1 step with cycle size 3, expected %v, got %v",
			[]int{1, 0, 0}, testMachine.CurrentRotors(),
		)
	}

	if takeSteps(testMachine, alphabetSize); !isPosCorrect(testMachine, []int{1, 9, 3}) {
		t.Errorf(
			"incorrect positions after 27 steps with cycle size 3, expected %v, got %v",
			[]int{1, 9, 3}, testMachine.CurrentRotors(),
		)
	}

	// TODO
	// Different step size and cycle size

	// TODO
	// Variable number of rotors
}

// Test initialization of rotors.
func TestInitRotors(t *testing.T) {
	testMachine := new(Machine)

	// Correct initialization
	testMachine.setNumberOfRotors(3)
	if err := testMachine.initRotors([]int{4, 2, 0}, 2, 13); err != nil {
		t.Errorf("correct init produces err, message %s", err.Error())
	}

	if !arePropertiesCorrect(testMachine, []int{4, 2, 0}, 2, 13) {
		t.Errorf("machine properties are incorrect")
	}

	// Incorrect initialization
	testMachine.setNumberOfRotors(3)
	if err := testMachine.initRotors([]int{3, 1, 0}, 2, 26); err == nil {
		t.Errorf("incorrect init doesn't produce err")
	}

	if !arePropertiesCorrect(testMachine, []int{3, 1, 0}, 2, 26) {
		t.Errorf("machine properties are incorrect")
	}
}

// Test calculation of taken steps.
func TestSetTakenSteps(t *testing.T) {
	testMachine := new(Machine)

	testMachine.setNumberOfRotors(3)
	testMachine.setStep(2)
	testMachine.setCycle(13)

	// Correct position
	if err := testMachine.setTakenSteps([]int{2, 4, 3}); err != nil {
		t.Errorf("correct position produces err")
	} else {
		steps := []int{1, 0, 0}
		for i, step := range testMachine.takenSteps {
			if steps[i] != step {
				t.Errorf("incorrect taken step\n expected %v, found %v", steps[i], step)
			}
		}
	}

	// Incorrect position
	if err := testMachine.setTakenSteps([]int{1, 2, 0}); err == nil {
		t.Errorf("incorrect position produces err")
	} else {
		steps := []int{0, 0, 0}
		for i, step := range testMachine.takenSteps {
			if steps[i] != step {
				t.Errorf("incorrect taken step\n expected %v, found %v", steps[i], step)
			}
		}
	}

	testMachine.setNumberOfRotors(7)
	testMachine.setStep(13)
	testMachine.setCycle(2)

	// Correct position
	if err := testMachine.setTakenSteps([]int{13, 13, 0, 13, 13, 13, 0}); err != nil {
		t.Errorf("correct position produces err")
	} else {
		steps := []int{1, 1, 0, 1, 1, 1, 0}
		for i, step := range testMachine.takenSteps {
			if steps[i] != step {
				t.Errorf("incorrect taken step\n expected %v, found %v", steps[i], step)
			}
		}
	}

	// Incorrect position
	if err := testMachine.setTakenSteps([]int{14, 12, 1, 0, 0, 0, 0}); err == nil {
		t.Errorf("incorrect position produces err")
	} else {
		steps := []int{0, 0, 0, 0, 0, 0, 0}
		for i, step := range testMachine.takenSteps {
			if steps[i] != step {
				t.Errorf("incorrect taken step\n expected %v, found %v", steps[i], step)
			}
		}
	}
}

// Test position setter.
func TestSetPosition(t *testing.T) {
	testMachine := new(Machine)
	testMachine.setNumberOfRotors(3)

	// Test wrong positions
	if err := testMachine.setRotorsPosition([]int{3, -22, 6}); err == nil {
		t.Errorf("incorrect rotor positions accepted")
	}

	// Test correct positions
	if err := testMachine.setRotorsPosition([]int{3, 22, 6}); err != nil {
		t.Errorf("correct positions not accepted")
	}
}

// Test step setter.
func TestSetStep(t *testing.T) {
	testMachine := new(Machine)

	if err := testMachine.setStep(23); err != nil {
		t.Errorf("correct step size \"%d\" not accepted", 23)
	}

	if err := testMachine.setStep(0); err == nil {
		t.Errorf("incorrect step size \"%d\" accepted", 0)
	}
}

// Test cycle setter.
func TestCycleSetter(t *testing.T) {
	testMachine := new(Machine)

	if err := testMachine.setCycle(5); err != nil {
		t.Errorf("correct cycle size \"%d\" not accepted", 5)
	}

	if err := testMachine.setCycle(0); err == nil {
		t.Errorf("incorrect cycle size \"%d\" accepted", 0)
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
		for j := 0; j < alphabetSize; j++ {
			if testMachine.rotors[i][j] != (j+positions[i])%alphabetSize {
				return false
			}
		}
	}

	return true
}
