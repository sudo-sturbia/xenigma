package machine

import (
	"github.com/sudo-sturbia/xenigma/pkg/helper"
)

// Rotor represents a mechanical rotor used in xenigma.
type Rotor struct {
	pathways  [alphabetSize]int // Connections that form electric pathways.
	position  int               // Current position of rotor.
	takeSteps int               // Number of rotor's taken steps.
}

// initRotor initializes all rotor's fields. Returns an error
// if given parameters are incorrect, nil otherwise.
func (r *Rotor) initRotor(pathways [alphabetSize]int, position int, step int, cycle int) error {
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
func (r *Rotor) setPosition(position int, step int, cycle int) {
	r.position = position
	r.takeSteps = (r.position / step) % cycle
}

// resetPosition resets rotor's position and taken step to 0.
func (r *Rotor) resetPosition() {
	r.position = 0
	r.takeSteps = 0
}
