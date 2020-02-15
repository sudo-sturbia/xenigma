package machine

import (
	"testing"
)

// Test rotor stepping with different steps and cycle sizes.
func TestStepRotors(t *testing.T) {
	testMachine := new(Machine)

	// 3 rotors, defaults
	testMachine.SetRotors(
		initRotorArr(
			[]int{0, 0, 0},
			[]int{1, 1, 1},
			[]int{26, 26, 26},
		),
	)

	takeSteps(testMachine, 10)
	if !compareSettings(testMachine, []int{10, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{10, 0, 0}, testMachine.Setting())
	}

	takeSteps(testMachine, 26*26*26)
	if !compareSettings(testMachine, []int{10, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{10, 0, 0}, testMachine.Setting())
	}

	// 3 rotors, for all {DefaultStep, cycle = 2}
	testMachine.SetRotors(
		initRotorArr(
			[]int{1, 0, 0},
			[]int{1, 1, 1},
			[]int{2, 2, 2},
		),
	)

	takeSteps(testMachine, 3)
	if !compareSettings(testMachine, []int{4, 2, 1}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{4, 2, 1}, testMachine.Setting())
	}

	takeSteps(testMachine, 20)
	if !compareSettings(testMachine, []int{24, 12, 6}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{24, 12, 6}, testMachine.Setting())
	}

	// 3 rotors, for all {step = 13, cycle = 2}
	testMachine.SetRotors(
		initRotorArr(
			[]int{0, 13, 0},
			[]int{13, 13, 13},
			[]int{2, 2, 2},
		),
	)

	takeSteps(testMachine, 2)
	if !compareSettings(testMachine, []int{0, 0, 13}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{0, 0, 13}, testMachine.Setting())
	}

	takeSteps(testMachine, 8)
	if !compareSettings(testMachine, []int{0, 0, 13}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{0, 0, 13}, testMachine.Setting())
	}

	// 3 rotors, for all {step = 2, cycle = 13}
	testMachine.SetRotors(
		initRotorArr(
			[]int{0, 0, 0},
			[]int{2, 2, 2},
			[]int{13, 13, 13},
		),
	)

	takeSteps(testMachine, 1)
	if !compareSettings(testMachine, []int{2, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{2, 0, 0}, testMachine.Setting())
	}

	takeSteps(testMachine, 14)
	if !compareSettings(testMachine, []int{4, 2, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{4, 2, 0}, testMachine.Setting())
	}

	// 12 rotors, defaults
	testMachine.SetRotors(
		initRotorArr(
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]int{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		),
	)

	takeSteps(testMachine, 5)
	if !compareSettings(testMachine, []int{5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, testMachine.Setting())
	}

	takeSteps(testMachine, 26*26*26*26)
	if !compareSettings(testMachine, []int{5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, testMachine.Setting())
	}

	// 12 rotors, for all {step = 13, cycle = 2}
	testMachine.SetRotors(
		initRotorArr(
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		),
	)

	takeSteps(testMachine, 15)
	if !compareSettings(testMachine, []int{13, 13, 13, 13, 0, 0, 0, 0, 0, 0, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{13, 13, 13, 13, 0, 0, 0, 0, 0, 0, 0, 0}, testMachine.Setting())
	}

	takeSteps(testMachine, 30)
	if !compareSettings(testMachine, []int{13, 0, 13, 13, 0, 13, 0, 0, 0, 0, 0, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{13, 0, 13, 13, 0, 13, 0, 0, 0, 0, 0, 0}, testMachine.Setting())
	}

	// 3 rotors, rotor 0 {step = 1, cycle = 1}
	//           rotor 1 {step = 2, cycle = 13}
	//           rotor 2 defaults
	testMachine.SetRotors(
		initRotorArr(
			[]int{0, 0, 0},
			[]int{1, 2, 1},
			[]int{1, 13, 26},
		),
	)

	takeSteps(testMachine, 8)
	if !compareSettings(testMachine, []int{8, 16, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{8, 16, 0}, testMachine.Setting())
	}

	takeSteps(testMachine, 2)
	if !compareSettings(testMachine, []int{10, 20, 0}) {
		t.Errorf("incorrect setting, expected %v, got %v",
			[]int{10, 20, 0}, testMachine.Setting())
	}

	// 30 rotors, even {step = 13, cycle = 2}
	//            odd  {step = 2, cycle = 13}
	testMachine.SetRotors(
		initRotorArr(
			[]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			[]int{
				13, 2, 13, 2, 13, 2, 13, 2, 13, 2,
				13, 2, 13, 2, 13, 2, 13, 2, 13, 2,
				13, 2, 13, 2, 13, 2, 13, 2, 13, 2,
			},
			[]int{
				2, 13, 2, 13, 2, 13, 2, 13, 2, 13,
				2, 13, 2, 13, 2, 13, 2, 13, 2, 13,
				2, 13, 2, 13, 2, 13, 2, 13, 2, 13,
			},
		),
	)

	takeSteps(testMachine, 30)
	expected := []int{
		0, 4, 13, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	if !compareSettings(testMachine, expected) {
		t.Errorf("incorrect setting, expected %v, got %v",
			expected, testMachine.Setting())
	}

	takeSteps(testMachine, 30)
	expected = []int{
		0, 8, 0, 2, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	if !compareSettings(testMachine, expected) {
		t.Errorf("incorrect setting, expected %v, got %v",
			expected, testMachine.Setting())
	}
}

// Benchmark rotor stepping 1000 steps in a 1000-rotor machine.
func BenchmarkStepRotors(b *testing.B) {
	m := Generate(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		takeSteps(m, 1000)
	}
}

// takeSteps steps the rotors given number of steps.
func takeSteps(m *Machine, steps int) {
	for i := 0; i < steps; i++ {
		m.stepRotors()
	}
}

// Test machine's rotor setter.
func TestSetRotors(t *testing.T) {
	testMachine := new(Machine)

	// Correct rotors
	err := testMachine.SetRotors(
		initRotorArr(
			[]int{0, 0, 0},
			[]int{1, 1, 1},
			[]int{26, 26, 26},
		),
	)

	if err != nil {
		t.Errorf("correct set of rotors produces error, %s", err.Error())
	}

	// Incorrect position
	pathways := [alphabetSize]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}

	r1, _ := NewRotor(pathways, -1, 1, 26)
	r2, _ := NewRotor(pathways, 11, 1, 26)
	r3, _ := NewRotor(pathways, 24, 1, 26)

	err = testMachine.SetRotors([]*Rotor{r1, r2, r3})
	if err == nil {
		t.Errorf("incorrect set of rotors doesn't produce error")
	}
}

// Test creation of a new Rotor object.
func TestInitRotor(t *testing.T) {
	// Correct init
	pathways := [alphabetSize]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}

	if _, err := NewRotor(pathways, 0, 1, 26); err != nil {
		t.Errorf("correct configuration produces error, %s", err.Error())
	}

	// Incorrect init - invalid position
	if _, err := NewRotor(pathways, -121, 1, 26); err == nil {
		t.Errorf("incorrect configuration doesn't produce error")
	}

	// Incorrect init - invalid step-cycle combination
	if _, err := NewRotor(pathways, 0, 2, 26); err == nil {
		t.Errorf("incorrect configuration doesn't produce error")
	}

	// Incorrect init - invalid step
	if _, err := NewRotor(pathways, 0, -21, 26); err == nil {
		t.Errorf("incorrect configuration doesn't produce error")
	}

	// Incorrect init - invalid cycle
	if _, err := NewRotor(pathways, 0, 1, 0); err == nil {
		t.Errorf("incorrect configuration doesn't produce error")
	}
}

// Return an array of rotors initialized with the given properties
// and with pathway connections ["a", "b", "c", ...] for all.
func initRotorArr(rotorHeads []int, steps []int, cycles []int) []*Rotor {
	pathways := [alphabetSize]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}

	rotorCount := len(rotorHeads)
	rotors := make([]*Rotor, rotorCount)
	for i := 0; i < rotorCount; i++ {
		rotors[i], _ = NewRotor(pathways, rotorHeads[i], steps[i], cycles[i])
	}

	return rotors
}

// compareSettings compares machine's setting and given positions.
func compareSettings(m *Machine, setting []int) bool {
	mSetting := m.Setting()
	for i := 0; i < m.numberOfRotors; i++ {
		if mSetting[i] != setting[i] {
			return false
		}
	}

	return true
}

// Test validation of step and cycle sizes.
func TestStepCycle(t *testing.T) {
	testRotor := new(Rotor)

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

	tempPos := [alphabetSize]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}

	for _, comb := range correct {
		if err := testRotor.isGivenConfigCorrect(tempPos, 0, comb[0], comb[1]); err != nil {
			t.Errorf("correct step-cycle produce err")
		}
	}

	if err := testRotor.isGivenConfigCorrect(tempPos, 0, 23, 32); err == nil {
		t.Errorf("incorrect step-cycle doesn't produce err")
	}

	if err := testRotor.isGivenConfigCorrect(tempPos, 0, 1, -3); err == nil {
		t.Errorf("incorrect step-cycle doesn't produce err")
	}

	if err := testRotor.isGivenConfigCorrect(tempPos, 0, 0, 1); err == nil {
		t.Errorf("incorrect step-cycle doesn't produce err")
	}
}
