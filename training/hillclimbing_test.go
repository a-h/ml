package training

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestHillClimbing(t *testing.T) {
	// Start with 1*1
	memory := []float64{1, 1}
	desiredOutput := 25.0
	// Train finding two numbers which can be multiplied together to equal 25.
	toTrain := func() (e float64, err error) {
		result := memory[0] * memory[1]
		e = math.Pow(result-desiredOutput, 2)
		return
	}
	// Train the algorithm, but step by 1 each time, so it's easy to test and know that we'll
	// get the desired factors of 5*5 in just a few steps.
	hc := NewHillClimbing(memory, 1, 1)
	for i := 1; i < 5; i++ {
		updatedMemory, err := hc.Next(toTrain)
		if err != nil {
			t.Errorf("unexpected error training simple multiplication: %v", err)
		}
		if !reflect.DeepEqual(updatedMemory, []float64{1.0 + float64(i), 1.0 + float64(i)}) {
			t.Errorf("expected the first movement to take the parameters up by 1, but got: %v", updatedMemory)
		}
	}
	// Check that training doesn't make a difference once we've got the correct value.
	updatedMemory, _ := hc.Next(toTrain)
	if !reflect.DeepEqual(updatedMemory, []float64{5.0, 5.0}) {
		t.Errorf("expected result to be 5*5, but was")
	}
	// Check that the error is now zero.
	if hc.BestError() != 0.0 {
		t.Errorf("expected the error to be zero, but got: %v", hc.BestError())
	}
	// Check that the best memory is 5*5 as expected.
	if !reflect.DeepEqual(hc.BestMemory(), []float64{5.0, 5.0}) {
		t.Errorf("expected the error to be zero, but got: %v", hc.BestError())
	}
}

func TestHillClimbingReturnsErrorsFromTheEvaluator(t *testing.T) {
	hc := NewHillClimbing([]float64{0, 1, 2}, 1, 1)
	ev := func() (e float64, err error) {
		err = errors.New("expected error")
		return
	}
	_, err := hc.Next(ev)
	if err.Error() != "expected error" {
		t.Errorf("unexpected error message while evaluating function: %v", err)
	}
}
