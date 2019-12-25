// Helper functions used for validation of
// slices and their properties.
package helper

import (
	"testing"
)

// Test are elements indices method
func TestAreElementsIndices(t *testing.T) {
	correct := []int{0, 2, 1, 3, 5, 6, 4, 8, 7}
	if !AreElementsIndices(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{0, 2, 1, 3, 5, 9, 8, 7}
	if AreElementsIndices(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}

// Test is symmetric method
func TestIsSymmetric(t *testing.T) {
	correct := []int{1, 0, 3, 2, 5, 4, 7, 6}
	if !IsSymmetric(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{0, 2, 1, 3, 5, 9, 8, 7}
	if IsSymmetric(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}

// Test are elements ordered indices method
func TestAreElementsOrderedIndices(t *testing.T) {
	correct := []int{4, 5, 0, 1, 2, 3}
	if !AreElementsOrderedIndices(correct) {
		t.Errorf("Correct input: %v not accepted", correct)
	}

	wrong := []int{1, 2, 4, 3, 21, -4}
	if AreElementsOrderedIndices(wrong) {
		t.Errorf("Incorrect input: %v accepted", wrong)
	}
}
