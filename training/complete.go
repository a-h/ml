package training

import (
	"context"
	"fmt"

	"github.com/a-h/ml/distance"
)

// Complete executes a training session against the trainee t, using data d, algorithm a and
// the distance function dist to calculate the distance between the expected and received values.
// Stoppers can be provided to limit the training to a time period, maximum number of iterations,
// error etc.
func Complete(t Trainee, d []Data, a Algorithm, dist distance.Function, stoppers ...Stopper) (iterations int, err error) {
	for {
		e := func() (e float64, err error) {
			return evaluateTrainee(t, d, dist)
		}
		updatedMemory, err := a.Next(e)
		if err != nil {
			return iterations, fmt.Errorf("training.Complete: error at iteration %v: %v", iterations, err)
		}
		t.SetMemory(updatedMemory)

		iterations++
		if shouldStop(iterations, a.BestError(), stoppers) {
			break
		}
	}
	return
}

func evaluateTrainee(t Trainee, d []Data, dist distance.Function) (e float64, err error) {
	for _, td := range d {
		actual, err := t.Calculate(td.Input)
		if err != nil {
			return e, fmt.Errorf("error training data: %v", err)
		}
		de, err := dist(actual, td.Expected)
		if err != nil {
			return e, fmt.Errorf("error calculating distance: %v", err)
		}
		e += de
	}
	e = e / float64(len(d))
	return
}

func shouldStop(iterations int, e float64, stoppers []Stopper) bool {
	for _, s := range stoppers {
		if s(iterations, e) {
			return true
		}
	}
	return false
}

// A Stopper provides a way to stop the training.
type Stopper func(iterations int, ee float64) bool

// StopWhenContextCancelled stops training if the context is cancelled.
// signal.Notify(c, os.Interrupt)
func StopWhenContextCancelled(ctx context.Context) Stopper {
	return func(iterations int, e float64) bool {
		select {
		case <-ctx.Done():
			return true
		default:
			return false
		}
	}
}

// StopWhenChannelReceives stops training if a value is received on the channel returned by the function.
// For example, the returned channel could be connected to OS signals for graceful shutdown:
// signal.Notify(c, os.Interrupt)
func StopWhenChannelReceives() (Stopper, chan interface{}) {
	c := make(chan interface{}, 1)
	return func(iterations int, e float64) bool {
		select {
		case <-c:
			return true
		default:
			return false
		}
	}, c
}

// StopAfterXIterations stops training after the specified iteration count has completed.
func StopAfterXIterations(x int) Stopper {
	return func(iterations int, e float64) bool {
		if iterations >= x {
			return true
		}
		return false
	}
}

// StopWhenErrorIsLessThan stops training when the error is below the value provided.
func StopWhenErrorIsLessThan(e float64) Stopper {
	return func(iterations int, ee float64) bool {
		if ee < e {
			return true
		}
		return false
	}
}

// StopWhenErrorIsGreaterThan stops training when the error is greater than the value provided.
func StopWhenErrorIsGreaterThan(e float64) Stopper {
	return func(iterations int, ee float64) bool {
		if ee > e {
			return true
		}
		return false
	}
}
