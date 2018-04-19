package main

import (
	"fmt"
	"math"
	"os"

	"github.com/a-h/ml/distance"
	"github.com/a-h/ml/random"
	"github.com/a-h/ml/rbf"
)

// TrainingData is data used to train the RBF network.
type TrainingData struct {
	Input    []float64
	Expected []float64
}

func main() {
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

	network, err := rbf.NewNetwork(
		rbf.NewNode(2, 1),
		rbf.NewNode(2, 1),
		rbf.NewNode(2, 1),
		rbf.NewBias(1),
	)
	if err != nil {
		fmt.Println("Error creating network:", err)
		os.Exit(-1)
	}

	finalError := math.MaxFloat64
	for i := 0; i < 1000000; i++ {
		// Run all of the training data through the network and calculate how good it is.
		var e float64
		for _, td := range ideal {
			// Execute the network against the training data.
			actual, err := network.Calculate(td.Input)
			if err != nil {
				fmt.Println("Error calculating network output:", err)
				os.Exit(-1)
			}
			// Calculate the distance away from the expected training result.
			dist, err := distance.Euclidean(actual, td.Expected)
			if err != nil {
				fmt.Println("Error calculating distance:", err)
				os.Exit(-1)
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
				if trainable, ok := nd.(rbf.TrainableNode); ok {
					memory := trainable.GetMemory()
					trainable.SetMemory(random.Float64Vector(-50, 50, len(memory)))
				}
			}
		}
	}
	fmt.Println("Output error:", finalError)
	fmt.Println(network)

	for _, v := range ideal {
		actual, err := network.Calculate(v.Input)
		if err != nil {
			fmt.Printf("failed to calculate network with input %v: %v\n", v.Input, err)
			os.Exit(-1)
		}
		fmt.Printf("input: %v, expected: %v, actual: %v\n", v.Input, v.Expected, actual)
	}
}
