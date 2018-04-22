package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/a-h/ml/distance"

	"github.com/a-h/ml/training"

	"github.com/a-h/ml/rbf"
)

func main() {
	ideal := []training.Data{
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
		rbf.NewNode(2, 1),
		rbf.NewNode(2, 1),
		rbf.NewBias(1),
	)
	if err != nil {
		fmt.Println("Error creating network:", err)
		os.Exit(-1)
	}

	algorithm := training.NewRandomGreedy(network.GetMemory())
	receiveStopper, c := training.StopWhenChannelReceives()
	relay := make(chan os.Signal, 1)
	go func() {
		fmt.Println("Press Ctrl-C to shut down.")
		<-relay
		fmt.Println("Shutting down.")
		c <- true
	}()
	signal.Notify(relay, os.Interrupt)
	start := time.Now()
	iterations, err := training.Complete(network, ideal, algorithm, distance.Euclidean, receiveStopper, training.StopWhenErrorIsLessThan(0.1))
	if err != nil {
		fmt.Println("Training error:", err)
		os.Exit(-1)
	}

	network.SetMemory(algorithm.BestMemory())
	fmt.Println("Time:", time.Now().Sub(start))
	fmt.Println("Iterations:", iterations)
	fmt.Println("Output error:", algorithm.BestError())

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
