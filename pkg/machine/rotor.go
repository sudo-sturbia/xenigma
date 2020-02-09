package machine

import (
	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Rotor represents a mechanical rotor used in xenigma.
type Rotor struct {
	pathways   [alphabetSize]int // Connections that form electric pathways.
	position   int               // Current position of rotor.
	takenSteps int               // Number of rotor's taken steps.
}

// step shifts rotor's position one step forward.
func (r *Rotor) step(step, cycle int) {
	r.position = (r.position + step) % alphabetSize
	r.takenSteps = (r.takenSteps + 1) % cycle
}

// InitRotor initializes all rotor's fields. Returns an error
// if given parameters are incorrect, nil otherwise.
func (r *Rotor) InitRotor(pathways [alphabetSize]int, position, step, cycle int) error {
	if err := r.isConfigCorrect(step); err != nil {
		return err
	}

	r.setPathways(pathways)
	r.setPosition(position, step, cycle)

	return nil
}

// isConfigCorrect returns an init error if rotor's config is
// incorrect, returns nil otherwise.
func (r *Rotor) isConfigCorrect(step int) error {
	if !helper.AreElementsIndices(r.pathways[:]) {
		return &initError{"electric pathways are incorrect"}
	}

	if (r.position)%step != 0 || r.position < 0 || r.position > alphabetSize {
		return &initError{"rotor's position is incorrect"}
	}

	return nil
}

// setPathways sets rotor's pathway connections.
func (r *Rotor) setPathways(pathways [alphabetSize]int) {
	r.pathways = pathways
}

// setPosition sets rotor's position and number of taken steps.
func (r *Rotor) setPosition(position, step, cycle int) {
	r.position = position
	r.takenSteps = (r.position / step) % cycle
}

// resetPosition resets rotor's position and taken step to 0.
func (r *Rotor) resetPosition() {
	r.position = 0
	r.takenSteps = 0
}

// Pathways returns rotor's pathway connections.
func (r *Rotor) Pathways() [alphabetSize]int {
	return r.pathways
}

// Position returns rotor's position.
func (r *Rotor) Position() int {
	return r.position
}
