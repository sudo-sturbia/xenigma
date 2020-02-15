package machine

import (
	"fmt"
	"math/rand"

	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Default step and cycle sizes used for a rotor.
const (
	DefaultStep  = 1
	DefaultCycle = 26
)

// Rotor represents a mechanical rotor used in xenigma.
type Rotor struct {
	pathways   [alphabetSize]int // Connections that form electric pathways.
	position   int               // Current position of rotor.
	takenSteps int               // Number of rotor's taken steps.
	step       int               // Size of shift between rotor steps.
	cycle      int               // Number of rotor steps considered a full cycle.
}

// NewRotor returns a pointer to a new Rotor object initialized with the specified
// fields. Returns an initialization error if fields are incorrect.
func NewRotor(pathways [alphabetSize]int, position, step, cycle int) (*Rotor, error) {
	r := new(Rotor)
	if err := r.InitRotor(pathways, position, step, cycle); err != nil {
		return nil, err
	}

	return r, nil
}

// GenerateRotor generates and returns a rotor with random configurations
// and returns a pointer to it.
func GenerateRotor() *Rotor {
	r := new(Rotor)
	r.pathways = [alphabetSize]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}

	rand.Shuffle(alphabetSize, func(j, k int) {
		r.pathways[j], r.pathways[k] = r.pathways[k], r.pathways[j]
	})

	r.setPosition(rand.Intn(alphabetSize), DefaultStep, DefaultCycle)
	r.setStep(DefaultStep)
	r.setCycle(DefaultCycle)

	return r
}

// takeStep shifts rotor's position one step forward.
func (r *Rotor) takeStep() {
	r.position = (r.position + r.step) % alphabetSize
	r.takenSteps = (r.takenSteps + 1) % r.cycle
}

// InitRotor initializes all rotor's fields including pathway connections,
// current position, step size, and cycle size. Returns an error if given
// parameters are incorrect, nil otherwise.
func (r *Rotor) InitRotor(pathways [alphabetSize]int, position, step, cycle int) error {
	if err := r.isGivenConfigCorrect(pathways, position, step, cycle); err != nil {
		return err
	}

	r.setPathways(pathways)
	r.setPosition(position, step, cycle)
	r.setStep(step)
	r.setCycle(cycle)

	return nil
}

// UseDefaultProperties sets all rotor's fields, except pathways,
// to their default values. Defaults are 'a' for position, 1 for
// step size, and 26 for cycle size.
func (r *Rotor) UseDefaultProperties() {
	r.setPosition(0, DefaultStep, DefaultCycle)
	r.setStep(DefaultStep)
	r.setCycle(DefaultCycle)
}

// IsConfigCorrect verifies rotor's current configuration, returns
// an error if rotor's fields are incorrect or incompatible.
func (r *Rotor) IsConfigCorrect() error {
	return r.isGivenConfigCorrect(r.pathways, r.position, r.step, r.cycle)
}

// isGivenConfigCorrect verifies given pathway connections, position,
// step size, and cycle size, and returns an error if given values are
// incorrect or incompatible.
func (r *Rotor) isGivenConfigCorrect(pathways [alphabetSize]int, position, step, cycle int) error {
	if !helper.AreElementsIndices(pathways[:]) {
		return &initError{"electric pathways are incorrect"}
	}

	if step <= 0 {
		return &initError{fmt.Sprintf("invalid step size %d", step)}
	}

	if cycle <= 0 {
		return &initError{fmt.Sprintf("invalid cycle size %d", cycle)}
	}

	if ((alphabetSize) % (step * cycle)) != 0 {
		return &initError{"cycle size and step size are not compatible, some collisions may occur"}
	}

	if (position)%step != 0 || position < 0 || position > alphabetSize {
		return &initError{"rotor's position is incorrect"}
	}

	return nil
}

// setPathways sets rotor's pathway connections.
func (r *Rotor) setPathways(pathways [alphabetSize]int) {
	r.pathways = pathways
}

// Pathways returns rotor's pathway connections.
func (r *Rotor) Pathways() [alphabetSize]int {
	return r.pathways
}

// setPosition sets rotor's position and number of taken steps.
func (r *Rotor) setPosition(position, step, cycle int) {
	r.position = position
	r.takenSteps = (r.position / (step % alphabetSize)) % cycle
}

// resetPosition resets rotor's position and taken step to 0.
func (r *Rotor) resetPosition() {
	r.position = 0
	r.takenSteps = 0
}

// Position returns rotor's position.
func (r *Rotor) Position() int {
	return r.position
}

// setStep sets rotor's step size.
func (r *Rotor) setStep(value int) {
	r.step = value % alphabetSize
}

// Step returns rotor's step size. Step represents the number of positions
// a rotor jumps when taking one step. The default size of a step is 1.
func (r *Rotor) Step() int {
	return r.step
}

// setCycle sets size of a rotor's full cycle.
func (r *Rotor) setCycle(value int) {
	r.cycle = value
}

// Cycle returns rotor's cycle size. Cycle is the number of steps that
// represent a rotor's full cycle. When a rotor completes a full cycle
// the following rotor is shifted. The default size of a cycle is 26.
func (r *Rotor) Cycle() int {
	return r.cycle
}
