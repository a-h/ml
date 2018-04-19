package rbf

import (
	"fmt"
	"math"
	"testing"

	"github.com/a-h/ml/random"

	"github.com/a-h/ml/distance"
)

func TestNetwork(t *testing.T) {
	ideal := []TrainingData{
		{
			Input:    []float64{0, 0},
			Expected: []float64{0},
		},
		{
			Input:    []float64{0, 1},
			Expected: []float64{1},
		},
		{
			Input:    []float64{1, 0},
			Expected: []float64{1},
		},
		{
			Input:    []float64{1, 1},
			Expected: []float64{0},
		},
	}

	network, err := NewNetwork(
		NewNode(2, 1),
		NewNode(2, 1),
		NewNode(2, 1),
		NewBias(1),
	)
	if err != nil {
		t.Fatal("Error creating network:", err)
	}

	finalError := math.MaxFloat64
	for i := 0; i < 1000000; i++ {
		// Run all of the training data through the network and calculate how good it is.
		var e float64
		for _, td := range ideal {
			// Execute the network against the training data.
			actual, err := network.Calculate(td.Input)
			if err != nil {
				t.Fatal("Error calculating network output:", err)
			}
			// Calculate the distance away from the expected training result.
			dist, err := distance.Euclidean(actual, td.Expected)
			if err != nil {
				t.Fatal("Error calculating distance:", err)
			}
			e += dist
			//fmt.Printf("%d: %v -> %v expected: %v error: %v\n", i, td.Input, actual, td.Expected, dist)
		}
		// If the error we have is lower than the current error, then keep the new network values.
		if e < finalError {
			finalError = e
			fmt.Println("Improvement. ", e)
		} else {
			// Try a different set of values.
			for _, nd := range network {
				if trainable, ok := nd.(TrainableNode); ok {
					memory := trainable.GetMemory()
					trainable.SetMemory(random.Float64Vector(-100, 100, len(memory)))
				}
			}
		}
	}
	t.Error("Output error:", finalError)
	fmt.Println(network)
}

type TrainingData struct {
	Input    []float64
	Expected []float64
}
