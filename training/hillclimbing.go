package training

import (
	"math"
)

// NewHillClimbing uses random parameters from the Min to Max value to train, attempting to minimise
// error.
func NewHillClimbing(memory []float64, velocity, acceleration float64) *HillClimbing {
	min, max := -10.0, 10.0
	hc := &HillClimbing{
		Min:      min,
		Max:      max,
		current:  memory,
		e:        math.MaxFloat64,
		velocity: velocity,
	}
	hc.movements = []float64{
		-acceleration,     // Go back
		-1 / acceleration, // Go back slower
		0,                 // Stop
		1 / acceleration,  // Go forward slower
		acceleration,      // Go forward
	}
	return hc
}

// HillClimbing is a training algorithm which uses random parameters from the Min to Max value to train,
// attempting to minimise error.
type HillClimbing struct {
	current []float64
	// best memory and error recorded during training.
	best []float64
	e    float64
	// Min and Max values for memory values.
	Min, Max float64
	// The amount to move in each direction.
	velocity float64
	// movements (forward, back, stay still)
	movements []float64
}

// Next records the error from the previous set of values, and returns a new set of values to try.
func (hc *HillClimbing) Next(ev Evaluator) ([]float64, error) {
	for i := range hc.current {
		// Calculate the best move for this dimension.
		bestMovement := 0
		bestError := math.MaxFloat64
		current := hc.current[i]
		for mi, mv := range hc.movements {
			hc.current[i] = current + (hc.velocity * mv)
			e, err := ev()
			if err != nil {
				return hc.current, err
			}
			if e < bestError {
				bestError = e
				bestMovement = mi
			}
		}
		// Stick with the best move for this dimension.
		hc.current[i] = current + (hc.velocity * hc.movements[bestMovement])
	}
	// Calculate the error across all dimensions.
	e, err := ev()
	hc.e = e
	return hc.current, err
}

// BestError returns the best (lowest) error discovered by training.
// If no training has happened, it will math.MaxFloat64.
func (hc *HillClimbing) BestError() (e float64) {
	return hc.e
}

// BestMemory returns the best set of parameters discovered by the algorithm during training.
// If no training has happened, it will return nil.
func (hc *HillClimbing) BestMemory() (memory []float64) {
	return hc.current
}
