package machine

import (
	"fmt"
)

// Rotors is a list of rotors used as a part of a machine.
type Rotors struct {
	rotors []*Rotor
	count  int
}

// NewRotors returns a new, initialized Rotors pointer, and an error if given
// rotor list is invalid.
func NewRotors(rotors []*Rotor) (*Rotors, error) {
	if len(rotors) == 0 {
		return nil, fmt.Errorf("no rotors given")
	}

	for i, rotor := range rotors {
		if rotor == nil {
			return nil, fmt.Errorf("rotor %d doesn't exist", i)
		}
		if err := rotor.Verify(); err != nil {
			return nil, fmt.Errorf("rotor %d: %w", i, err)
		}
	}

	return &Rotors{
		rotors: rotors,
		count:  len(rotors),
	}, nil
}

// GenerateRotors returns a list of randomly generated rotors.
func GenerateRotors(count int) *Rotors {
	rotors := make([]*Rotor, count)
	for i := 0; i < count; i++ {
		rotors[i] = GenerateRotor()
	}

	return &Rotors{
		rotors: rotors,
		count:  count,
	}
}

// Rotor returns rotor at position i.
func (r *Rotors) Rotor(i int) (*Rotor, error) {
	if i < 0 || i >= r.count {
		return nil, fmt.Errorf("invalid index %d", i)
	}
	return r.rotors[i], nil
}

// takeStep moves the rotors one step forward.
func (r *Rotors) takeStep() {
	for i, rotor := range r.rotors {
		if i != 0 && (r.rotors[i-1].takenSteps != 0) { // Previous rotor didn't complete a cycle.
			break
		}
		rotor.takeStep()
	}
}

// Verify verifies that rotors' are valid, and returns an error otherwise.
func (r *Rotors) Verify() error {
	if len(r.rotors) == 0 || r.count != len(r.rotors) {
		return fmt.Errorf("invalid number of rotors")
	}

	for i, rotor := range r.rotors {
		if err := rotor.Verify(); err != nil {
			return fmt.Errorf("rotor %d: %w", i, err)
		}
	}
	return nil
}

// UseDefaults sets all fields of each rotor to, except pathways, to their default
// values.
func (r *Rotors) UseDefaults() {
	for _, rotor := range r.rotors {
		rotor.UseDefaults()
	}
}

// Setting returns rotors' current setting. A setting is a list containing the current
// position of each rotor.
func (r *Rotors) Setting() []int {
	setting := make([]int, r.count)
	for i, rotor := range r.rotors {
		setting[i] = rotor.Position()
	}

	return setting
}

// Count returns number of rotors.
func (r *Rotors) Count() int {
	return r.count
}
