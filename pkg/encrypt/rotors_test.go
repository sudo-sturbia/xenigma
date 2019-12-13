// Encrypt messages using engima code
package encrypt

import (
	"testing"
)

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
