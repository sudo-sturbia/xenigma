package machine

import (
	"testing"
)

func TestZeroToNSlice(t *testing.T) {
	correct := []int{0, 2, 1, 3, 5, 6, 4, 8, 7}
	if !zeroToNSlice(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{0, 2, 1, 3, 5, 9, 8, 7}
	if zeroToNSlice(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}

func TestZeroToN(t *testing.T) {
	correct := map[int]int{
		0: 1,
		1: 2,
		2: 3,
		3: 4,
		4: 0,
	}
	if !zeroToN(correct, len(correct)) {
		t.Errorf("correct input %v not accepted", correct)
	}

	wrong := map[int]int{
		0: 1,
		1: 2,
		2: 3,
		3: 4,
		4: 5,
	}
	if zeroToN(wrong, len(wrong)) {
		t.Errorf("incorrect input %v accepted", wrong)
	}
}

func TestIsSymmetric(t *testing.T) {
	correct := map[int]int{
		0: 1,
		1: 0,
		2: 3,
		3: 2,
		4: 4,
	}
	if !isSymmetric(correct) {
		t.Errorf("correct input: %v not accepted", correct)
	}

	wrong := map[int]int{
		0: 1,
		1: 0,
		2: 3,
		3: 4,
		4: 2,
	}
	if isSymmetric(wrong) {
		t.Errorf("incorrect input: %v accepted", wrong)
	}
}
