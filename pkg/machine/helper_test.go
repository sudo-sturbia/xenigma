package machine

import (
	"testing"
)

func TestAreElementsIndices(t *testing.T) {
	correct := []int{0, 2, 1, 3, 5, 6, 4, 8, 7}
	if !areElementsIndices(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{0, 2, 1, 3, 5, 9, 8, 7}
	if areElementsIndices(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}

func TestIsSymmetric(t *testing.T) {
	correct := []int{1, 0, 3, 2, 5, 4, 7, 6}
	if !isSymmetric(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{0, 2, 1, 3, 5, 9, 8, 7}
	if isSymmetric(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}

func TestAreElementsOrderedIndices(t *testing.T) {
	correct := []int{4, 5, 0, 1, 2, 3}
	if !areElementsOrderedIndices(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{1, 2, 4, 3, 21, -4}
	if areElementsOrderedIndices(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}
