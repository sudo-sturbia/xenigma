package machine

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestTakeSteps tests rotors' movement using different step and cycle
// sizes.
func TestTakeSteps(t *testing.T) {
	for i, test := range []struct {
		rotors   *Rotors
		steps    []int
		expected [][]int
	}{
		{
			rotors: newTestRotors(
				t,
				[]int{0, 0, 0},
				[]int{1, 1, 1},
				[]int{26, 26, 26},
			),
			steps:    []int{10, 26 * 26 * 26},
			expected: [][]int{{10, 0, 0}, {10, 0, 0}},
		},
		{
			rotors: newTestRotors(
				t,
				[]int{1, 0, 0},
				[]int{1, 1, 1},
				[]int{2, 2, 2},
			),
			steps:    []int{3, 20},
			expected: [][]int{{4, 2, 1}, {24, 12, 6}},
		},
		{
			rotors: newTestRotors(
				t,
				[]int{0, 13, 0},
				[]int{13, 13, 13},
				[]int{2, 2, 2},
			),
			steps:    []int{2, 8},
			expected: [][]int{{0, 0, 13}, {0, 0, 13}},
		},
		{
			rotors: newTestRotors(
				t,
				[]int{0, 0, 0},
				[]int{2, 2, 2},
				[]int{13, 13, 13},
			),
			steps:    []int{1, 14},
			expected: [][]int{{2, 0, 0}, {4, 2, 0}},
		},
		{
			rotors: newTestRotors(
				t,
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				[]int{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
			),
			steps: []int{5, 26 * 26 * 26 * 26},
			expected: [][]int{
				{5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			rotors: newTestRotors(
				t,
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
				[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			),
			steps: []int{15, 30},
			expected: [][]int{
				{13, 13, 13, 13, 0, 0, 0, 0, 0, 0, 0, 0},
				{13, 0, 13, 13, 0, 13, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			rotors: newTestRotors(
				t,
				[]int{0, 0, 0},
				[]int{1, 2, 1},
				[]int{1, 13, 26},
			),
			steps:    []int{8, 2},
			expected: [][]int{{8, 16, 0}, {10, 20, 0}},
		},
		{
			rotors: newTestRotors(
				t,
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
			steps: []int{30, 30},
			expected: [][]int{
				{
					0, 4, 13, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
				{
					0, 8, 0, 2, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
		},
	} {
		for j, step := range test.steps {
			for k := 0; k < step; k++ {
				test.rotors.takeStep()
			}

			if diff := cmp.Diff(test.expected[j], test.rotors.Setting()); diff != "" {
				t.Errorf("test %d: mismatch (-want +got):\n%s", i, diff)
			}
		}
	}
}

// BenchmarkTakeStep benchmarks the stepping of a 1000-rotor machine.
func BenchmarkTakeStep(b *testing.B) {
	r := GenerateRotors(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			r.takeStep()
		}
	}
}

// TestNewRotors test config validation.
func TestNewRotors(t *testing.T) {
	for i, test := range []struct {
		arr       []*Rotor
		shouldErr bool
	}{
		{
			arr: newRotorArr(
				t,
				[]int{0, 0, 0},
				[]int{1, 1, 1},
				[]int{26, 26, 26},
			),
			shouldErr: false,
		},
		{
			arr: newRotorArr(
				t,
				[]int{-1, 11, 24},
				[]int{1, 1, 1},
				[]int{26, 26, 26},
			),
			shouldErr: true,
		},
	} {
		_, err := NewRotors(test.arr)
		if test.shouldErr && err == nil {
			t.Errorf("test %d: want error, got nil", i)
		} else if !test.shouldErr && err != nil {
			t.Errorf("test %d: want nil, got %w", i, err)
		}
	}
}

// TestNewRotor test rotor validation.
func TestNewRotor(t *testing.T) {
	var pathways [alphabetSize]int
	for i := 0; i < alphabetSize; i++ {
		pathways[i] = i
	}

	for i, test := range []struct {
		position  int
		step      int
		cycle     int
		shouldErr bool
	}{
		{
			position:  0,
			step:      1,
			cycle:     26,
			shouldErr: false,
		},
		{
			position:  -121,
			step:      1,
			cycle:     26,
			shouldErr: true,
		},
		{
			position:  0,
			step:      2,
			cycle:     26,
			shouldErr: true,
		},
		{
			position:  0,
			step:      -21,
			cycle:     26,
			shouldErr: true,
		},
		{
			position:  0,
			step:      1,
			cycle:     0,
			shouldErr: true,
		},
	} {
		_, err := NewRotor(pathways, test.position, test.step, test.cycle)
		if test.shouldErr && err == nil {
			t.Errorf("test %d: want error, got nil", i)
		} else if !test.shouldErr && err != nil {
			t.Errorf("test %d: want nil, got %w", i, err)
		}
	}
}

// TestStepCycle tests step-cycle compatability.
func TestStepCycle(t *testing.T) {
	var pathways [alphabetSize]int
	for i := 0; i < alphabetSize; i++ {
		pathways[i] = i
	}

	for i, test := range []struct {
		step      int
		cycle     int
		shouldErr bool
	}{
		{
			step:      1,
			cycle:     1,
			shouldErr: false,
		},
		{
			step:      1,
			cycle:     2,
			shouldErr: false,
		},
		{
			step:      2,
			cycle:     1,
			shouldErr: false,
		},
		{
			step:      1,
			cycle:     13,
			shouldErr: false,
		},
		{
			step:      13,
			cycle:     1,
			shouldErr: false,
		},
		{
			step:      1,
			cycle:     26,
			shouldErr: false,
		},
		{
			step:      26,
			cycle:     1,
			shouldErr: false,
		},
		{
			step:      2,
			cycle:     13,
			shouldErr: false,
		},
		{
			step:      13,
			cycle:     2,
			shouldErr: false,
		},
		{
			step:      23,
			cycle:     32,
			shouldErr: true,
		},
		{
			step:      1,
			cycle:     -3,
			shouldErr: true,
		},
		{
			step:      0,
			cycle:     1,
			shouldErr: true,
		},
	} {
		err := verifyRotor(pathways, 0, test.step, test.cycle)
		if test.shouldErr && err == nil {
			t.Errorf("test %d: want error, got nil", i)
		} else if !test.shouldErr && err != nil {
			t.Errorf("test %d: want nil, got %w", i, err)
		}
	}
}

// newTestRotors creates and returns a Rotors with the given properties
// for testing, should be used only for testing as errors are not accounted for.
func newTestRotors(t *testing.T, setting []int, steps []int, cycles []int) *Rotors {
	t.Helper()
	rotors := newRotorArr(t, setting, steps, cycles)
	return &Rotors{
		rotors: rotors,
		count:  len(rotors),
	}
}

// newRotorArr creates a new Rotor array using given properties to use
// for testing.
func newRotorArr(t *testing.T, setting []int, steps []int, cycles []int) []*Rotor {
	t.Helper()
	var pathways [alphabetSize]int
	for i := 0; i < alphabetSize; i++ {
		pathways[i] = i
	}

	rotors := make([]*Rotor, len(setting))
	for i := 0; i < len(setting); i++ {
		rotors[i], _ = NewRotor(pathways, setting[i], steps[i], cycles[i])
	}
	return rotors
}
