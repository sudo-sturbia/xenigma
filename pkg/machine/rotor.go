package machine

import (
	"fmt"
	"math/rand"
)

// Default values for rotor properties.
const (
	DefaultPosition = 0
	DefaultStep     = 1
	DefaultCycle    = 26
)

// Rotor represents a mechanical rotor used in xenigma. A rotor contains connections
// used to make electric pathways and generate a path through the machine.
type Rotor struct {
	pathways   [alphabetSize]int // Connections that form electric pathways.
	position   int               // Current position.
	takenSteps int               // Number of taken steps.
	step       int               // Size of shift between steps, in characters.
	cycle      int               // Number of steps considered a full cycle.
}

// NewRotor returns a pointer to a new, initialized Rotor, and an error if
// given fields are invalid.
func NewRotor(pathways [alphabetSize]int, position, step, cycle int) (*Rotor, error) {
	if err := verifyRotor(pathways, position, step, cycle); err != nil {
		return nil, err
	}

	return &Rotor{
		pathways:   pathways,
		position:   position,
		takenSteps: (position / (step % alphabetSize)) % cycle,
		step:       step % alphabetSize,
		cycle:      cycle,
	}, nil
}

// GenerateRotor generates and returns a rotor with random config.
func GenerateRotor() *Rotor {
	var pathways [alphabetSize]int
	for i := 0; i < alphabetSize; i++ {
		pathways[i] = i
	}

	rand.Shuffle(
		alphabetSize,
		func(j, k int) {
			pathways[j], pathways[k] = pathways[k], pathways[j]
		},
	)

	position := rand.Intn(alphabetSize)
	return &Rotor{
		pathways:   pathways,
		position:   position,
		takenSteps: (position / (DefaultStep % alphabetSize)) % DefaultCycle,
		step:       DefaultStep,
		cycle:      DefaultCycle,
	}
}

// Verify verifies rotor's current configuration, returns an error if rotor's
// fields are incorrect or incompatible.
func (r *Rotor) Verify() error {
	return verifyRotor(r.pathways, r.position, r.step, r.cycle)
}

// verifyRotor verifies given pathway connections, position, step size, and
// cycle size, and returns an error if given values are incorrect or incompatible.
func verifyRotor(pathways [alphabetSize]int, position, step, cycle int) (err error) {
	switch {
	case !areElementsIndices(pathways[:]):
		err = fmt.Errorf("electric pathways are incorrect")
	case step <= 0:
		err = fmt.Errorf("invalid step: %d", step)
	case cycle <= 0:
		err = fmt.Errorf("invalid cycle: %d", cycle)
	case (position)%step != 0 || position < 0 || position > alphabetSize:
		err = fmt.Errorf("invalid position: %d", position)
	case ((alphabetSize) % (step * cycle)) != 0:
		err = fmt.Errorf("cycle and step are incompatible, some collisions may occur")
	}
	return err
}

// takeStep moves rotor one step forward.
func (r *Rotor) takeStep() {
	r.position = (r.position + r.step) % alphabetSize
	r.takenSteps = (r.takenSteps + 1) % r.cycle
}

// UseDefaults sets all rotor's fields, except pathways, to their default
// values. Defaults are 'a' for position, 1 for step, and 26 for cycle.
func (r *Rotor) UseDefaults() {
	r.position = DefaultPosition
	r.takenSteps = (DefaultPosition / (DefaultStep % alphabetSize)) % DefaultCycle
	r.step = DefaultStep
	r.cycle = DefaultCycle
}

// Pathways returns rotor's pathway connections.
func (r *Rotor) Pathways() [alphabetSize]int {
	return r.pathways
}

// Position returns rotor's current position.
func (r *Rotor) Position() int {
	return r.position
}

// Step returns rotor's step size. Step represents the number of positions
// a rotor jumps when moving one step forward, and defaults to 1.
func (r *Rotor) Step() int {
	return r.step
}

// Cycle returns rotor's cycle size. Cycle is the number of steps that
// represent a rotor's full cycle. When one rotor in a machine completes a
// full cycle the following rotor is shifted.
func (r *Rotor) Cycle() int {
	return r.cycle
}
